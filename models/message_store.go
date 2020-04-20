package models

import (
	"errors"
	"time"
)

type MessageToStore struct {
	SenderName				string 			`json:"sender_name,omitempty"`
	ReceiverName			string			`json:"receiver_id,omitempty"`
	Message					string			`json:"message,omitempty"`
	Time					time.Time		`json:"time,omitempty"`
}

/*
	invalid parameters have been passed
 */
var NotValidMessage = errors.New("not valid credentials")

/*
	check for validity
 */
func (m *MessageToStore) IsValid() (bool) {
	if m.ReceiverName == "" || m.SenderName == "" || m.Message == "" {
		return false
	}
	return true
}

/*
	this function stores the message on db
	a sender sends a message -> the server accepts -> (we are here) stores on db -> notifies a receiver
 */
func (m *MessageToStore) StoreOnDb() (error) {

	if ok := m.IsValid(); !ok {
		return NotValidMessage
	}

	if err := GetDB().Create(m).Error; err != nil {
		return err
	}
	
	return nil

}
