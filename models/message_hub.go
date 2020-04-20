package models

import "fmt"

/*
	this struct is to handle websocket registration and reverse,
		and broadcast of messages available in one place
 */
type Hub struct {
	Clients				map[string]*WSClient
	Broadcast 			chan *MessageToSend
	Register			chan *WSClient
	Unregister			chan *WSClient
}

/*
	this global value will be assigned to the Hub object
		hub will be used using this global object
 */
var iHub *Hub

/*
	Create a new hub
 */
func NewHub() (*Hub) {
	return &Hub{
		Clients: 		make(map[string]*WSClient),
		Broadcast:  	make(chan *MessageToSend, 100),
		Register:   	make(chan *WSClient, 100),
		Unregister: 	make(chan *WSClient, 100),
	}
}

/*
	this function initiates a hub at the beginning
 */
func init() {
	iHub = NewHub()
}

/*
	this function will be used to access the hub object
 */
func GetHub() (*Hub) {
	return iHub
}

/*
	this function is to run the hub, and handle sending messages
		and registering / unregistering client websocket connections

	can:
		* register
		* unregister
		* send a message
 */
func (hub *Hub) Run() () {
	for {
		select {
			case conn := <- hub.Register:
				if oldConn, ok := hub.Clients[conn.Id]; ok {
					_ = oldConn.Conn.Close()
				}
				hub.Clients[conn.Id] = conn
				fmt.Println("registering a client " + conn.Id)

			case conn := <- hub.Unregister:
				if oldConn, ok := hub.Clients[conn.Id]; ok {
					_ = oldConn.Conn.Close()
				}

			case messageToSend := <- hub.Broadcast:
				/*
					first look for the client connection from the dictionary
					if there is such connection send to the user
				 */
				if conn, ok := hub.Clients[messageToSend.ReceiverId]; ok {
					/*
						convert the message to byte and send byte to the receiver connection
					 */
					if messageByte, err := messageToSend.ConvertToByte(); err != nil {
						select {
							case conn.Send <- messageByte:
							default:
								/*
									this means, an error has occurred while sending a message
										which indicates that the client connection is no more valid
										then remove it from the dictionary
								 */
								close(conn.Send)
								delete(hub.Clients, conn.Id)
						}
					}
				}
		}
	}
}


