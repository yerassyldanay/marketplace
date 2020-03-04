package models

import (
	"marketplace/utils"
	"net/http"
	"os"
	"strings"
	"time"
)

type CourseDocuments struct {
	ID 						int 				`json:"id" gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	UserName				string 				`json:"user_name" gorm:""`
	CourseID 				int 				`json:"course_id" gorm:"unique_index:course_documents_index"`
	SectionID				int					`json:"section_id" gorm:""`
	DocumentName			string 				`json:"document_name" gorm:"unique_index:course_documents_index"`
	DocumentFormat			string				`json:"document_format" gorm:"size:30"`
	UpdatedAt 				time.Time			`json:"updated_at" gorm:"default:current_timestamp"`
	DisplayName 			string				`json:"display_name" gorm:"unique_index:course_documents_index"`
}

func (*CourseDocuments) TableName() string {
	return utils.DbNameCourseDocuments
}

func (courseDocs *CourseDocuments) Upload_or_update_document_of_course_section() (map[string]interface{}, error) {
	var response = utils.Message(http.StatusBadRequest, "Could not upload a document")
	if !strings.HasPrefix(courseDocs.DocumentFormat, ".") {
		courseDocs.DocumentFormat = "." + courseDocs.DocumentFormat
	}

	/*
		firstly, check that course section exists
	 */
	var tempCourseSection = &struct {
		CourseID			int 			`json:"course_id"`
		SectionID			int 			`json:"section_id"`
	}{}

	var main_query = "select c.id as course_id, cs.id as section_id from users u inner join course c " +
		" on u.id = c.instructor_id " +
		" inner join course_section cs on c.id = cs.course_id " +
		" where u.user_name = $1 and c.id = $2 and cs.id = $3 limit 1;"
	
	if err := GetDB().Raw(main_query, courseDocs.UserName, courseDocs.CourseID, courseDocs.SectionID).Scan(tempCourseSection).Error;
		err != nil {
			return response, err
	}

	/*
		if it is verified that the course section indeed exists
			then we will create such course document
	 */
	if err := GetDB().Create(courseDocs).Error; err != nil {
		return response, err
	}

	return utils.Message(http.StatusOK, "Successfully added a document to a section"), nil
}

func (courseDocs *CourseDocuments) Get_documents_by_course_id() (map[string]interface{}, error) {
	var rows, err = GetDB().Table(utils.DbNameCourseDocuments).Where("course_id = $1", courseDocs.CourseID).Rows()
	if err != nil {
			return utils.Message(http.StatusBadRequest, "Could not upload a document"), err
	}
	defer rows.Close()

	var tempList = []CourseDocuments{}
	for rows.Next() {
		var temp = &CourseDocuments{}
		if err := GetDB().ScanRows(rows, temp); err != nil {
			continue
		}
		temp.UserName = ""
		if temp.DocumentName != "" {
			temp.DocumentName = "/" + utils.Folder_where_documents_are_stored + "/" + temp.DocumentName + temp.DocumentFormat
		} else {
			temp.DocumentName = ""
		}

		tempList = append(tempList, *temp)
	}

	var response = utils.Message(http.StatusOK, "Ok")
	response["items"] = tempList

	return response, nil
}

func (courseDocs *CourseDocuments) Delete_document_of_the_course_section_from_directory() (map[string]interface{}, error) {
	var temp = &struct {
		DocumentName			string			`json:"document_name"`
		DocumentFormat			string			`json:"document_format"`
	}{}

	var main_query string
	/*
		getting the file name, which further will be used to delete files
	 */
	main_query = " select document_name, document_format from course_documents where id = $1 and user_name = $2;"

	if err := GetDB().Raw(main_query, courseDocs.ID, courseDocs.UserName).Scan(temp).Error; err != nil {
		return utils.Message(http.StatusNotFound, "Could not delete a file"), err
	}

	/*
		firstly, delete file from the database
	 */
	//main_query = " update course_documents set document_name = '', document_format = '' where id = $1 and user_name = $2 ; "
	//if err := GetDB().Exec(main_query, courseDocs.ID, courseDocs.UserName).Error; err != nil {
	//	return utils.Message(http.StatusNotFound, "Could not delete a file"), err
	//}

	/*
		delete file from directory
	 */
	var path_to_file = utils.Folder_where_documents_are_stored + "/" + temp.DocumentName + temp.DocumentFormat
	if err := os.Remove(path_to_file); err != nil {
		return utils.Message(http.StatusNotFound, "Could not delete a file"), err
	}

	return utils.Message(http.StatusOK, "Ok"), nil
}

/*
	DELETE the document from the database and directory
 */
func (courseDocs *CourseDocuments) Delete_document_of_course_section() (map[string]interface{}, error) {
	var response = utils.Message(http.StatusBadRequest, "Could not delete a file")
	if _, err := courseDocs.Delete_document_of_the_course_section_from_directory(); err != nil {
		return response, err
	}

	var main_query = " delete from course_documents where id = $1 and user_name = $2 ; "
	if err := GetDB().Exec(main_query, courseDocs.ID, courseDocs.UserName).Error; err != nil {
		return response, err
	}

	return utils.Message(http.StatusOK, "Successfully deleted"), nil
}



