package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"marketplace/models"
	"marketplace/utils"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var nameOfFunc = "JWT_AUTH"

		/*
			NOTE: this is to escape the auth middleware
		 */
		//next.ServeHTTP(writer, request)
		//return

		/*
			Before sending a POST request, client agent checks whether it is possible
				to knock the door of the server
			For that it uses OPTIONS method
		*/
		if request.Method == "OPTIONS" {
			next.ServeHTTP(writer, request)
			return
		}

		var requestPath = request.URL.Path
		for _, url := range utils.NoNeedToAuth {
			if requestPath == url {
				next.ServeHTTP(writer, request)
				return
			} else if strings.HasPrefix(requestPath, "/upload") {
				next.ServeHTTP(writer, request)
				return
			}
		}

		var response_to_return = utils.Message(http.StatusNetworkAuthenticationRequired, "An Invalid Token has been provided")

		/*
			Getting session_id (this is authentication token)
		*/
		var tokenHeader = request.Header.Get("Authorization")
		//fmt.Println("---")
		if tokenHeader == "" {
			response_to_return = utils.Message(http.StatusForbidden, "Missing Auth Token")
			utils.Respond(writer, response_to_return, &utils.LogMessage{
				Success:      false,
				FunctionName: nameOfFunc,
				Message:      "token is nil",
				Req:          request,
			})
			return
		}

		/*
			session_id is composed of two strings such as "Bearer", it is followed by "token_string"
		*/
		var splitCompile = regexp.MustCompile(`\s+`)
		var splitted = splitCompile.Split(tokenHeader, -1)
		if len(splitted) != 2 {
			response_to_return = utils.Message(http.StatusForbidden, "Missing Auth Token")
			utils.Respond(writer, response_to_return, &utils.LogMessage{
				Success:      false,
				FunctionName: nameOfFunc,
				Message:      "Must contain only two words",
				Req:          request,
			})
			return
		}

		fmt.Println("Token: ", os.Getenv("token_password"))

		var tokenNeeded = splitted[1]
		var tokenStruct = &models.Token{}

		var token, _ = jwt.ParseWithClaims(tokenNeeded, tokenStruct, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//if err != nil {
		//	response_to_return = utils.Message(http.StatusForbidden, "Invalid Auth Token")
		//	utils.Respond(writer, response_to_return, &utils.LogMessage{
		//		Success:      			false,
		//		FunctionName: 			nameOfFunc,
		//		Message:      			"Must contain only two words",
		//		Req:         			request,
		//	})
		//	return
		//}

		if !token.Valid {
			//response_to_return = utils.Message(http.StatusForbidden, "Auth Token is not Valid")
			//utils.Respond(writer, response_to_return, &utils.LogMessage{
			//	Success:      			false,
			//	FunctionName: 			nameOfFunc,
			//	Message:      			"Must contain only two words",
			//	Req:         			request,
			//})
			//return
		}

		//var cntx = context.WithValue(request.Context(), "username", tokenStruct.UserName)
		//request = request.WithContext(cntx)

		request.Header.Add("role", tokenStruct.Audience)
		request.Header.Add("username", tokenStruct.Subject)

		next.ServeHTTP(writer, request)
	})
}
