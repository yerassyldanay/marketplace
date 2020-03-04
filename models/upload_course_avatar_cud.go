package models

import (
	"marketplace/utils"
	"net/http"
	"os"
)

/*
	this is how info on course logo will be passed
	this function stores a logo of the course after validating that the current user has rights
		to make changes
*/
type CourseLogo struct {
	//UserID					int				`json:"user_id"`
	UserName				string 			`json:"user_name"`
	CourseID 				int				`json:"course_id"`
	ImageName 				string 			`json:"image_name"`
	ImageFormat				string			`json:"image_format"`
}

/*
	this uploads and updates course avatar
 */
func (logo *CourseLogo) Upload_or_update_course_avatar() (map[string]interface{}, error) {
	/*
		Firstly DELETE image files from the directory then create new one
	*/
	if _, err := logo.Delete_Logo_From_Directory(); err != nil {
		return utils.Message(http.StatusBadRequest, "Could not perform internal operations"), err
	}

	var main_query = "update course set image_name = $1, image_format = $2 " +
		" where id = $3 and instructor_id = (select id from users where user_name = $4 limit 1); "

	/*
		Now, let's update the logo path
	*/
	if err := GetDB().Exec(main_query, logo.ImageName, logo.ImageFormat, logo.CourseID, logo.UserName).Error;
		err != nil {
		return utils.Message(http.StatusForbidden, "Could not update course logo"), err
	}

	return utils.Message(http.StatusOK, "Successfully uploaded / updated / deleted"), nil
}

/*
	200 - has been successfully deleted
	404 - the image file is not found
	403 - a user has no permission to make changes
 */
func (logo *CourseLogo) Delete_Logo_From_Directory() (map[string]interface{}, error) {
	var main_query = "select c.image_name as image_name, c.image_format as image_format " +
		" from users u join course c on c.instructor_id = (select id from users where user_name = $1 limit 1) " +
		" where u.user_name = $1 and c.id = $2 ;"
	
	var temp = &struct {
		ImageName 		string 			`json:"image_name"`
		ImageFormat		string 			`json:"image_format"`
	}{}

	if tempDb := GetDB().Raw(main_query, logo.UserName, logo.CourseID).Scan(temp); tempDb.RecordNotFound() {
		return utils.Message(http.StatusOK, "Has been successfully deleted"), nil
	} else if tempDb.Error != nil {
		return utils.Message(http.StatusOK, "Could not delete"), tempDb.Error
	}

	if temp.ImageName != "" {
		for key, _ := range utils.Resize_Values {
			if key == "default" {
				_ = os.Remove(utils.Folder_where_avatars_are_stored + "/" + temp.ImageName + temp.ImageFormat)
				continue
			}
			_ = os.Remove(utils.Folder_where_avatars_are_stored + "/" + temp.ImageName + "_" + key + temp.ImageFormat)
		}
	}

	return utils.Message(http.StatusOK, "Successfully deleted"), nil
}
