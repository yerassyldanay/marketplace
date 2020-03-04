package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Message(status_to_send uint, message_to_send string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status_to_send,
		"message": message_to_send,
	}
}

func Respond(w http.ResponseWriter, message_to_send map[string]interface{}, message *LogMessage) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Add("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Add("Content-Type", "application/json")

	var nameOfFunc = "NEW_ENCODER_RESPONSE"

	if message.Req != nil {
		fmt.Println("Got request from ", message.Req.RemoteAddr, " | to path "+message.Req.URL.Path)
	}
	if err := json.NewEncoder(w).Encode(message_to_send); err != nil {
		var newMessage = LogMessage{
			Success:      false,
			FunctionName: nameOfFunc,
			Message:      err.Error(),
			Req:          message.Req,
		}
		LogExtended(&newMessage)
	}

	// log the state of the response
	LogExtended(message)
}
