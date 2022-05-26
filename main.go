package main

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"gtihub.com/MuhammadSuryono/cbt-uploader/controllers"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/exam"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/user"
)

func main() {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()
	_ = db.Connection.AutoMigrate(&exam.ExamResult{})
	_ = db.Connection.AutoMigrate(&user.UserParticipantWithTypeExam{})
	_ = db.Connection.AutoMigrate(&user.TypeExamResult{})

	server := controllers.RunServer()

	controller := controllers.NewExam()
	api := server.Group("api/v1/exam")
	{
		api.POST("/push", controller.ExamPush)
		api.GET("/result/:session_key/download", controller.ExportToExcel)
		api.GET("/submit", controller.CountResultExam)
	}

	server.Run(":8081")
}
