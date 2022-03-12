package user

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"testing"
)

func TestUser(t *testing.T) {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()

	data := GetAllDataWithoutPagination()
	ExportToFile(data)
}
