package models

import (
	"errors"
	"github.com/gorilla/websocket"
	"marketplace/utils"
	"net/http"
	"time"
)

const (
	TypeConnectionEstablished = 0
	TypeConnectionMessage = 1
	TypeConnectionError = 2
)

const (
	OperationConnect = 0
	OperationMessageSend = 1
	OperationMessageGet	= 2
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

/*
	invalid parameters have been passed
 */
var ErrorInvalidParametersHaveBeenPassed = errors.New("invalid parameters have been passed")

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:    0,
	WriteBufferSize:   0,
}

type WSClient struct {
	Id						string					`json:"-"`
	Hub						*Hub					`json:"-"`
	Conn					*websocket.Conn			`json:"-"`
	Send					chan []byte				`json:"-"`
}

/*
	this function gets a message from channel, consecutively sends the message to the user
	channel - is a specific channel fo the user

	Note: it is not broadcast channel, which is for all users
 */
func (c *WSClient) writePump() {
	ticker := time.NewTicker( pingPeriod )

	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()

	for {
		select {

			case message, ok := <- c.Send:
				_ = c.Conn.SetWriteDeadline( time.Now().Add( writeWait ) )
				if !ok {
					_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				w, err := c.Conn.NextWriter( websocket.TextMessage )
				if err != nil {
					return
				}

				_, _ = w.Write( message )
				n := len(c.Send)

				for i := 0; i < n; i++ {
					_, _ = w.Write( <- c.Send )
				}

				if err := w.Close(); err != nil {
					return
				}

			case <- ticker.C:
				_ = c.Conn.SetWriteDeadline( time.Now().Add( writeWait ) )
				if err := c.Conn.WriteMessage( websocket.PingMessage, nil); err != nil {
					return
				}
		}
	}
}

/*
	this function connects the websocket connection to the hub
		returns error if something goes wrong
 */
func connectWS(hub *Hub, w http.ResponseWriter, r *http.Request) (error) {

	var username = r.Header.Get(utils.Username)
	var rule = r.Header.Get(utils.Rule)

	if rule == "" || username == "" {
		return ErrorInvalidParametersHaveBeenPassed
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	/*
		Here we come up with a name for the connection, as each connection is unique,
			we will take username and rule
			{ "username_rule": conn }
	*/
	var websocket_client_key = username + "_" + rule
	var newClientConn = &WSClient{
		Id:   	websocket_client_key,
		Hub:  	hub,
		Conn: 	conn,
		Send: 	make(chan []byte, 256),
	}

	/*
		this line of code sends to hub.register that it can register a new client websocket connection
			in simple words, adds a connection to the dictionary
			in relation of an id to a connection ({id : connection})
	 */
	newClientConn.Hub.Register <- newClientConn
	go newClientConn.writePump()

	return nil
}