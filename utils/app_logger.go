package utils

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

var logStruct = &Logger{}
var file *os.File

/*
	This function initiates for you a log file
	Using GetLogger() function, you can always log message to this file
	NOTE: it is your responsibility to close a file at at the end of the function you are calling
          this function
*/
func init() {
	file, err := os.OpenFile("logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	logStruct.logger = log.New(file, "DB_ACCESS: ", log.LstdFlags)
}

/*
	This function takes a message and writes it to the file
*/
func (logger *Logger) LogMessageToFile(logMessage string) {
	logger.logger.Printf(logMessage)
}

/*
	Using this function, you can get logger anywhere (e.g. anywhere in the project).
	Theoretically, it is possible to call outside the project, but I will not recommend it
*/
func GetLogger() *Logger {
	return logStruct
}

/*
	This function is used to get an access to a file and close the file at the end of the function
*/
func GetLogFile() *os.File {
	return file
}
