package controller

import (
	"marketplace/utils"
	"net/http"
)

/*
	POST
	this function enables to clearly deal with role of a user
 */
var Upload_document = func(w http.ResponseWriter, r *http.Request) {

	var role = r.Header.Get("role")
	//fmt.Println()
	if role == utils.Instructor {
		Upload_documents_to_course_section(w, r)
		return
	}
}

/*
	GET
 */
var Get_document = func(w http.ResponseWriter, r *http.Request) {
	var role = r.Header.Get("role")

	//fmt.Println()

	if role == utils.Instructor {
		Get_documents_of_course_section(w, r)
		return
	}
}

/*
	DELETE
 */
var Delete_document = func(w http.ResponseWriter, r *http.Request) {
	var role = r.Header.Get("role")
	//fmt.Println()
	if role == utils.Instructor {
		Delete_document_of_a_course_section(w, r)
		return
	}
}
