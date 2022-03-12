package controllers

import "github.com/gin-gonic/gin"

func (e *Excel) ExamPush(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pushed",
	})
}
