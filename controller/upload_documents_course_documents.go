package controller

import (
	"marketplace/models"
	"marketplace/utils"
	"net/http"
	"strconv"
	"time"
)

var Upload_documents_to_course_section = func(w http.ResponseWriter, r *http.Request) {

	var nameOfFunc = "UPLOAD_COURSE_DOCUMENTS_CONTR"
	var response = utils.Message(http.StatusBadRequest, "Could not upload a file")

	/*
		Get parameters of a document
	*/
	var courseId = r.URL.Query().Get("course_id")
	var sectionId = r.URL.Query().Get("section_id")
	var file_display_name = r.URL.Query().Get("file_display_name")

	/*
		checking whether valid parameters have been passed
	*/
	if courseId == "" || sectionId == "" || file_display_name == "" {
		response = utils.Message(http.StatusBadRequest, "Some parameters are not provided")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"Could not find all parameters",
			Req:          		r,
		})
		return
	}

	/*
		convert parameters to int
		Note: it is needed because documentCourse object has already been created
		and it is mapped on the database
		In case parameters are left as string, have to change the struct -> this changes column datatype on db
	*/
	var name string
	course_id, err_course := strconv.Atoi(courseId)
	section_id, err_section := strconv.Atoi(sectionId)

	if err_section != nil || err_course != nil {
		response = utils.Message(http.StatusBadRequest, "Some parameters are not provided")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"Could not find all parameters",
			Req:          		r,
		})
		return
	}

	/*
		get parameters of token
	*/
	var username = r.Header.Get("username")

	/*
		---------------------------------------- <<<>>> -------------------------------------------------
	 */

	/*
		Parsing a file to get meta-data on a file (name, format, etc.), file in bytes, generated string (filename), err
	*/
	var handler, file_in_bytes, filename, err = Parse_file(r)
	if err != nil {
		response = utils.Message(http.StatusBadRequest, "File. Could not parse a file being uploaded")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"could not find all parameters",
			Req:          		r,
		})
		return
	}

	file_format, err := Parse_the_format_of_the_file(handler)
	if err != nil {
		response = utils.Message(http.StatusBadRequest, "Some parameters are not provided")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"file format is not provided",
			Req:          		r,
		})
		return
	}

	var state = true
	var message = "Successfully uploaded"
	response = utils.Message(http.StatusOK, "Could not upload a document")

	/*
		-------------------------------------------<<<>>>---------------------------------------------
	*/

	/*
		creating courseDocuments object and assigning values
			* here we will call model to access database
	 */
	var courseDoc = models.CourseDocuments {
		UserName:			username,
		CourseID:       	course_id,
		SectionID:      	section_id,
		DocumentName:   	filename,
		DocumentFormat: 	file_format,
		UpdatedAt:      	time.Now(),
		DisplayName: 		name,
	}

	response, err = courseDoc.Upload_or_update_document_of_course_section()

	if err != nil {
		state = false
		message = err.Error()
	}

	/*
		Storing a document on the directory
	*/
	err = Store_document(utils.Folder_where_documents_are_stored, filename, file_format, file_in_bytes)
	if err != nil {
		state = false
		message = err.Error()
	}

	utils.Respond(w, response, &utils.LogMessage{
		Success:      		state,
		FunctionName: 		nameOfFunc,
		Message:      		message,
		Req:          		r,
	})
}

var Get_documents_of_course_section = func(w http.ResponseWriter, r *http.Request) {

	var nameOfFunc = "GET_ALL_DOCUMENTS_CONTR"

	/*
		get parameters of token
	*/
	//var role = r.Header.Get("role")
	//var username = r.Header.Get("username")

	var courseId = r.URL.Query().Get("course_id")
	var response map[string]interface{}

	if courseId == "" {
		response = utils.Message(http.StatusBadRequest, "Could not get all documents")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"course id is invalid",
			Req:          		r,
		})
		return
	}

	course_id, err := strconv.Atoi(courseId)
	if err != nil {
		response = utils.Message(http.StatusBadRequest, "Could not get all documents")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"course id is invalid",
			Req:          		r,
		})
		return
	}

	/*
		creating a course object
	 */
	var courseDocuments = models.CourseDocuments{
		CourseID:      		course_id,
	}

	response, err = courseDocuments.Get_documents_by_course_id()
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

var Delete_document_of_a_course_section = func(w http.ResponseWriter, r *http.Request) {

	var response map[string]interface{}
	var nameOfFunc = "GET_ALL_DOCUMENTS_CONTR"

	/*
		get parameters of token
	*/
	var username = r.Header.Get("username")

	/*
		Get parameters of a document
	*/
	var documentID = r.URL.Query().Get("document_id")

	if documentID == "" {
		response = utils.Message(http.StatusBadRequest, "Could not delete a document")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"document id is invalid",
			Req:          		r,
		})
		return
	}

	document_id, err := strconv.Atoi(documentID)

	if err != nil {
		response = utils.Message(http.StatusBadRequest, "Could not delete a document")
		utils.Respond(w, response, &utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		"document id is invalid",
			Req:          		r,
		})
		return
	}

	/*
		creating a course object
	*/
	var courseDocuments = models.CourseDocuments{
		UserName:			username,
		ID:      			document_id,
		DocumentName: 		"",
		DocumentFormat: 	"",
	}

	response, err = courseDocuments.Delete_document_of_course_section()
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
