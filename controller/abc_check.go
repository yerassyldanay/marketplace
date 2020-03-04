package controller

import (
	"fmt"
	"marketplace/utils"
	"net/http"
)

/*
	Welcome, home!
*/
var Health_check_ = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HEALTH_CHECK. Received request from " + r.RemoteAddr)
	response := utils.Message(http.StatusOK, "Welcome home, brother!")
	utils.Respond(w, response, &utils.LogMessage{
		Success:      true,
		FunctionName: "HEALTH_CHECK",
		Message:      "just checking",
		Req:          r,
	})
}
