package controller

import (
	"marketplace/models"
	"marketplace/utils"
	"net/http"
)

/*
	This function get parameters as:
		id
		referer (e.g. course, etc)
	Updates the values on the database, while removing old images
*/
var Upload_or_update_user_avatar = func(w http.ResponseWriter, r *http.Request) {

	var nameOfFunc = "UPLOAD_USER_AVATAR"
	var response  = utils.Message(http.StatusBadRequest, "Could not upload a file")

	handler, file_in_bytes, filename, err := Parse_file(r)
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		r,
		})
		return
	}

	file_format, err := Parse_the_format_of_the_file(handler)
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		r,
		})
		return
	}

	if err := Resize_and_store_image(utils.Resize_Values, file_in_bytes, utils.Folder_where_avatars_are_stored, filename, file_format); err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		r,
		})
		return
	}

	/*
		Getting token parameters
	*/
	var username = r.Header.Get("username")
	//fmt.Print(username)

	/*
		the last step is to store it on database
	*/
	var user_avatar = models.UserAvatar{
		UserName:    		username,
		ImageName:   		filename,
		ImageFormat: 		file_format,
	}
	response, err = user_avatar.Upload_or_update_user_avatar()

	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		r,
		})
		return
	}

	utils.Respond(w, response, &utils.LogMessage{
		Success:      		true,
		FunctionName: 		nameOfFunc,
		Message:      		"Successfully uploaded",
		Req:          		r,
	})
}

var Delete_user_avatar = func(w http.ResponseWriter, r *http.Request) {

	var nameOfFunc = "DELETE_USER_AVATAR"

	/*
		get parameters from session token
	 */
	var username = r.Header.Get("username")
	var userAvatar = models.UserAvatar {
		UserName:    	username,
		ImageName:		"",
		ImageFormat:	"",
	}

	response, err := userAvatar.Upload_or_update_user_avatar()
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		r,
		})
		return
	}

	utils.Respond(w, response, &utils.LogMessage{
		Success:      		true,
		FunctionName: 		nameOfFunc,
		Message:      		"Successfully deleted",
		Req:          		r,
	})
}
