package models

import (
	"errors"
	"time"
)

var localErrorInvalidMessage = errors.New("invalid message")

type MessageGeneralType struct {
	Status					int						`json:"status"`
	Message					string					`json:"message"`
	Time 					time.Time				`json:"time"`
}

type MessageToSend struct {
	Message				*MessageGeneralType				`json:"message" gorm:""`
	ReceiverId			string							`json:"receiver"`
}

func(m *MessageToSend) Send_message_to_indicated_address( hub *Hub ) (error) {
	if m.Message == nil {
		return localErrorInvalidMessage
	}

	hub.Broadcast <- m
	return nil
}

/*
	this method of the object converts the message to byte
		which allows us in one place handle conversion
 */
func (m *MessageToSend) ConvertToByte() ([]byte, error) {
	return []byte(""), nil
}
