package controllers

import (
	"github.com/gin-gonic/gin"
	"gtihub.com/MuhammadSuryono/cbt-uploader/file/excel"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/group"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/question"
	"path/filepath"
)

func (e *Excel) ImportExcelToDatabase(c *gin.Context) {
	var uri models.Uri
	_ = c.BindUri(&uri)

	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Code:      500,
			IsSuccess: false,
			Message:   "Failed Upload File",
		})
		return
	}

	excelHandler := excel.NewExcelHandler(filename)
	data, _ := excelHandler.ReadExcel()

	for _, v := range data {
		dataGroup := group.CreateGroupQuestion(uri.ListQuestionId, v.GroupName, v.GroupDescription, v.FilenameGroup, v.QuestionGroup)
		question.AddExamQuestion(v, dataGroup.Id, uri.ListQuestionId)
	}

	c.JSON(200, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Success Upload File",
	})

}
