package main

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"gtihub.com/MuhammadSuryono/cbt-uploader/controllers"
)

func main() {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()

	server := controllers.RunServer()

	controller := controllers.NewExam()
	api := server.Group("api/v1/exam")
	{
		api.POST("/push", controller.ExamPush)
		api.GET("/result/:session_key/download", controller.ExportToExcel)
	}

	server.Run(":8081")
}
