package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"marketplace/auth"
	"marketplace/controller"
	"marketplace/models"
	"marketplace/utils"
	"net/http"
	"regexp"
)

func main() {
	/*
		At the end of the main function, close the .log file
		(for more information, please, refer to utils.app_logger.go)
	*/
	defer utils.GetLogFile().Close()

	/*
		At the end of the function, close connection with database
	*/
	defer models.GetDB().Close()

	/*
		creating a new router
	*/
	var router = mux.NewRouter().StrictSlash(true)

	/*
		following lines of code enables us to access static files (images and documents in our case)
	*/
	var IMAGES_DIR = "/" + utils.Folder_where_avatars_are_stored + "/"
	router.PathPrefix(IMAGES_DIR).Handler(http.StripPrefix(IMAGES_DIR, http.FileServer(http.Dir("./"+utils.Folder_where_avatars_are_stored))))

	var DOCUMENTS_DIR = "/" + utils.Folder_where_documents_are_stored + "/"
	router.PathPrefix(DOCUMENTS_DIR).Handler(http.StripPrefix(DOCUMENTS_DIR, http.FileServer(http.Dir("./"+utils.Folder_where_documents_are_stored))))

	router.HandleFunc("/api/upload/image", controller.Delete_image).Methods("DELETE")
	router.HandleFunc("/api/upload/image", controller.Upload_image).Methods("POST")

	router.HandleFunc("/api/upload/document", controller.Get_document).Methods("GET")
	router.HandleFunc("/api/upload/document", controller.Upload_document).Methods("POST")
	router.HandleFunc("/api/upload/document", controller.Delete_document).Methods("DELETE")

	router.HandleFunc("/api/check", controller.Health_check_).Methods("GET")

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())
	}).Schemes("ws")
	/*
		SECURITY
		this part sets function to a session_id (session token) validation
	*/
	router.Use(auth.JwtAuthentication)

	/*
		When POST request is going to be made, client agent (browser) sends first OPTIONS requests
			to check whether CORS is enables or not
		For this reason, method that handles OPTIONS requests based on the url pattern is
			provided below
	*/
	router.Methods("OPTIONS").MatcherFunc(func(r *http.Request, match *mux.RouteMatch) bool {
		matchCase, err := regexp.MatchString("/*", r.URL.Path)
		if err != nil {
			return false
		}
		return matchCase
	}).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Add("Content-Type", "application/json")
	})

	var port = "6010"
	fmt.Println("Serving on port " + port)
	fmt.Println(http.ListenAndServe(":"+port, router))
}
