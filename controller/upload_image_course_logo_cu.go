package controller

import (
	"marketplace/models"
	"marketplace/utils"
	"net/http"
	"strconv"
)

var Upload_course_logo = func(w http.ResponseWriter, r *http.Request) {
	var nameOfFunc = "UPLOAD_COURSE_LOGO"
	var response = utils.Message(http.StatusBadRequest, "Could not upload a file")

	/*
		downloaded a file from request
	 */
	handler, file_in_bytes, filename, err := Parse_file(r)
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"1. " + err.Error(),
			Req:          		r,
		})
		return
	}

	/*
		parsed a format of the file (e.g. image/png -> png)
	 */
	file_format, err := Parse_the_format_of_the_file(handler)
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"2. " + err.Error(),
			Req:          		r,
		})
		return
	}

	/*
		stored a file at different resolutions
	 */
	if err = Resize_and_store_image(utils.Resize_Values, file_in_bytes, utils.Folder_where_avatars_are_stored, filename, file_format); err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"3. " + err.Error(),
			Req:          		r,
		})
		return
	}

	/*
		Getting token parameters
	*/
	var username = r.Header.Get("username")
	course_ids := r.URL.Query()["course_id"]

	/*
		parse query parameters
	*/
	if len(course_ids) < 1 {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"referer / course id is not indicated",
			Req:          		r,
		})
		return
	}

	course_id, err := strconv.Atoi(course_ids[0])
	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"4. " + err.Error(),
			Req:          		r,
		})
		return
	}

	/*
		the last step is to store it on database
	*/
	var logo = models.CourseLogo{
		UserName:    		username,
		CourseID:    		course_id,
		ImageName:   		filename,
		ImageFormat: 		file_format,
	}

	response, err = logo.Upload_or_update_course_avatar()

	if err != nil {
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"5. " + err.Error(),
			Req:          		r,
		})
		return
	}

	utils.Respond(w, response, &utils.LogMessage{
		Success:      		true,
		FunctionName: 		nameOfFunc,
		Message:      		"uploaded",
		Req:          		r,
	})
}
