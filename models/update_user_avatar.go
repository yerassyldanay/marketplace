package models

import (
	"marketplace/utils"
	"net/http"
	"strings"
)

type UserAvatar struct {
	UserName    			string 				`json:"user_name"`
	ImageName				string				`json:"image_name"`
	ImageFormat 			string 				`json:"image_format"`
}

/*
	UPLOAD
 */
func (userAvatar *UserAvatar) Upload_or_update_user_avatar() (map[string]interface{}, error) {
	var main_query = "update users set image_name = $1, image_format = $2 where user_name = $3;"
	var err error
	var response = utils.Message(http.StatusForbidden, "Could not update a user avatar")

	if !strings.HasPrefix(userAvatar.ImageFormat, ".") {
		userAvatar.ImageFormat = "." + userAvatar.ImageFormat
	}

	/*
		delete files from directory
	 */
	if tempResponse, err := userAvatar.Delete_user_avatar_from_directory(); err != nil {
		return tempResponse, err
	}
	/*
		storing new values on database
	 */
	if err = GetDB().Exec(main_query, userAvatar.ImageName, userAvatar.ImageFormat, userAvatar.UserName).Error; err != nil {
		return response, err
	}

	return utils.Message(http.StatusOK, "Ok"), nil
}

/*
	DELETE
 */
func (userAvatar *UserAvatar) Delete_user_avatar_from_directory() (map[string]interface{}, error) {
	var tempImageName = &struct {
		ImageName			string			`json:"image_name"`
		ImageFormat			string			`json:"image_format"`
	}{}

	/*
		getting the filename if it exists
	*/
	if err := GetDB().Table("users").Where("user_name = ?", userAvatar.UserName).First(tempImageName).Error; err != nil {
		return utils.Message(http.StatusForbidden, "Invalid. Could not update a user avatar"), err
	}

	/*
		deleting files from directory
	*/
	if err := Remove_files_by_pattern("./" + utils.Folder_where_avatars_are_stored, tempImageName.ImageName + "*", tempImageName.ImageFormat); err != nil {
		return utils.Message(http.StatusForbidden, "Invalid. Could not update a user avatar"), err
	}

	return utils.Message(http.StatusOK, "Successfully deleted"), nil
}