package controllers

import (
	"github.com/gin-gonic/gin"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/exam"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/user"
)

func (e *Excel) ExamPush(c *gin.Context) {
	var param models.ParameterPushAnswer
	_ = c.BindJSON(&param)

	exam.PushAnswer(param.RegisterNUmber, param.Id, param.Value)

	c.JSON(200, gin.H{
		"response": "pushed",
	})
}

func (e *Excel) CountResultExam(c *gin.Context) {
	var param models.CountResultExam
	_ = c.BindJSON(&param)

	go user.SubmitResultQuestion(param.RegisterNUmber)

	c.JSON(200, gin.H{
		"response": "submitting",
	})
}
