package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/user"
)

func (e *Excel) ExportToExcel(c *gin.Context) {
	var uri models.Uri
	_ = c.BindUri(&uri)

	dataUser := user.GetAllDataWithoutPagination(uri.SessionKey)
	user.ExportToFile(dataUser, fmt.Sprintf("%s.xlsx", uri.SessionKey))

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", uri.SessionKey))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Length", fmt.Sprintf("%d", len(dataUser)))
	c.File(fmt.Sprintf("%s.xlsx", uri.SessionKey))
}
