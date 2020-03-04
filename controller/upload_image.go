package controller

import (
	"marketplace/utils"
	"net/http"
)

var Upload_image = func(w http.ResponseWriter, r *http.Request) {
	var name_of_func = "UPLOAD_IMAGE"
	var response = utils.Message(http.StatusBadRequest, "Invalid Parameters have been provided")

	refs := r.URL.Query()["referer"]
	if len(refs) == 0 {
		utils.Respond(w, response,
			&utils.LogMessage{
				Success:      false,
				FunctionName: name_of_func,
				Message:      "referer / id is not indicated",
				Req:          r,
			})
		return
	}

	var role = r.Header.Get("role")
	//fmt.Println("-------")
	if role == utils.Instructor && refs[0] == utils.Course {
		Upload_course_logo(w, r)
		return
	} else if refs[0] == utils.HomePage {
		Upload_or_update_user_avatar(w, r)
		return
	}

	utils.Respond(w, response,
		&utils.LogMessage{
			Success:      false,
			FunctionName: name_of_func,
			Message:      "invalid parameters have been passed",
			Req:          r,
		})
}

var Get_image = func(w http.ResponseWriter, r *http.Request) {
	/*
		we do not need this function
	*/
}

var Delete_image = func(w http.ResponseWriter, r *http.Request) {
	var name_of_func = "DELETE_IMAGE"
	refs := r.URL.Query()["referer"]
	if len(refs) < 1 {
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Invalid Parameters have been provided"),
			&utils.LogMessage{
				Success:      false,
				FunctionName: name_of_func,
				Message:      "referer / id is not indicated",
				Req:          r,
			})
		return
	}

	var role = r.Header.Get("role")
	//fmt.Println()
	if role == utils.Instructor && refs[0] == utils.Course {
		Delete_logo_of_the_course(w, r)
		return
	} else if refs[0] == utils.HomePage {
		Delete_user_avatar(w, r)
		return
	}
}
