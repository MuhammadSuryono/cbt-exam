package controllers

import "github.com/gin-gonic/gin"

type IExcel interface {
	ExamPush(c *gin.Context)
	ExportToExcel(c *gin.Context)
	CountResultExam(c *gin.Context)
}

type Excel struct {
}

func NewExam() IExcel {
	return &Excel{}
}
