package controllers

import (
	"github.com/gin-gonic/gin"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/user"
)

func (e *Excel) ListUserParticipant(c *gin.Context) {
	var queryRequest models.QueryRequest
	_ = c.BindQuery(&queryRequest)

	dataUser := user.GetAllParticipant(queryRequest.Page, queryRequest.Size)
	c.JSON(200, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Success Retrieve Data",
		Data:      dataUser,
	})
}
