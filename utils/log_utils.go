package utils

import (
	"fmt"
	"net/http"
)

type LogMessageInterface interface {
	LogExtended() string
	LogMini() string
}

type LogMessage struct {
	Success 				bool
	FunctionName			string
	Message					string
	Req 					*http.Request
}

func (message *LogMessage) LogExtended() string {
	if message.Req == nil {
		return message.LogMini()
	}

	if message.Success {
		return fmt.Sprintf("[%v] Successful. Url: %v | Host: %v | Message: %v",
			message.FunctionName, message.Req.URL.Path, message.Req.RemoteAddr, message.Message)
	}

	return fmt.Sprintf("[%v] Error occurred while calling a function. Url: %v | Host: %v | Message: %v",
			message.FunctionName, message.Req.URL.Path, message.Req.RemoteAddr, message.Message)
}

func (message *LogMessage) LogMini() string {
	if message.Success {
		return fmt.Sprintf("[%v] Successful. Message: %v",
			message.FunctionName, message.Message)
	}
	return fmt.Sprintf("[%v] Error occurred while calling a function. Message: %v",
		message.FunctionName, message.Message)
}

func LogExtended(message *LogMessage) {
	GetLogger().LogMessageToFile(message.LogExtended())
}

func LogMini(message *LogMessage) {
	GetLogger().LogMessageToFile(message.LogMini())
}
