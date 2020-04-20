package utils

const (
	Folder_where_avatars_are_stored   = "upload_images"
	Folder_where_documents_are_stored = "upload_documents"

	DbNameCourseDocuments = "course_documents"

	Instructor = "Instructor"
	Student    = "Student"

	Course   = "course"
	HomePage = "home"

	Username = "username"
	Rule = "rule"

	Number_of_characters_in_an_image_name = 50
)

var Resize_Values = map[string][]int{
	"default":   {111, 111},
	"150_x_150": {150, 150},
	//"200_x_200": {200, 200},
	//"200_x_150": {200, 150},
	//"250_x_200": {250, 200},
}

var NoNeedToAuth = []string{
	"/api/check",
}
