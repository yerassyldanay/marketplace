package controller

import (
	"marketplace/models"
	"marketplace/utils"
	"net/http"
	"strconv"
)

var Delete_logo_of_the_course = func(w http.ResponseWriter, r *http.Request) {
	var nameOfFunc = "DELETE_LOGO_OF_COURSE"
	var response = utils.Message(http.StatusBadRequest, "Could not delete a file")

	var username = r.Header.Get("username")
	course_ids := r.URL.Query()["course_id"]

	if len(course_ids) < 1 {
		utils.Respond(w, response,
			&utils.LogMessage{
				Success:      	false,
				FunctionName: 	nameOfFunc,
				Message:      	"Invalid parameters have been passed",
				Req:          	r,
			})
		return
	}

	course_id, err := strconv.Atoi(course_ids[0])
	if err != nil {
		utils.Respond(w, response,
			&utils.LogMessage{
				Success:      	false,
				FunctionName: 	nameOfFunc,
				Message:      	"Invalid parameters have been passed",
				Req:          	r,
			})
		return
	}

	var logo = models.CourseLogo{
		UserName:     	username,
		CourseID:    	course_id,
		ImageName:		"",
		ImageFormat:	"",
	}

	/*
		this function is used to set image name and format values on db to ""
	 */
	response, err = logo.Upload_or_update_course_avatar()

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
		Message:      		"Ok",
		Req:          		r,
	})
}
