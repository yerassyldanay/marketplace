## Marketplace (upload)

*This project has been created to handle uploading / updating / deleting image and documents of any type.*
##### This project uses:
- PostgreSQL ( > 9.x.x )
- docker ( >18.x.x )
- docker-compose (1.8.x)

##### Headers:
```cookie
Content-Type: multipart/form-data
session_id: Bearer this_is_a_token
```
**_Note:_** session_id must contain two strings 

##### Documentation & API Test Env
```
Open on browser: https://documenter.getpostman.com/view/8978378/SW7aXTQN?version=latest
Download Postman documentation: https://www.getpostman.com/collections/3012e19c67f2db1b9129
```

##### In main.go you can find:
```go
/*
    creating a new router
 */
var router = mux.NewRouter().StrictSlash(true)

router.HandleFunc("/api/upload/image", controller.Delete_image).Methods("DELETE")
router.HandleFunc("/api/upload/image", controller.Upload_image).Methods("POST")

router.HandleFunc("/api/upload/document", controller.Get_document).Methods("GET")
router.HandleFunc("/api/upload/document", controller.Upload_document).Methods("POST")
router.HandleFunc("/api/upload/document", controller.Delete_document).Methods("DELETE")

router.HandleFunc("/api/check", controller.Health_check_).Methods("GET")
```

##### CORS:
```go
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
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
    w.Header().Add("Content-Type", "application/json")
})
```

##### constants
```go
const (
	Folder_where_avatars_are_stored = "upload_images"
	Folder_where_documents_are_stored = "upload_documents"
)

const (
	Number_of_characters_in_an_image_name = 50
)

const (
	DbNameCourseDocuments = "course_documents"
)

const (
	Instructor = "instructor"
	Student = "student"
)

const (
	Course = "course"
	HomePage = "home"
)

var Resize_Values = map[string][]int{
	"default": 	 {111, 111},
	"150_x_150": {150, 150},
}

var NoNeedToAuth = []string {
	"/api/check",
}
```

##### to_do:
- [ ] can support only png and jpg
- [ ] in case of an error, upload part does not remove files from directory
- [ ] test the authorization part
- [ ] test dockerfile

