package user

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"testing"
)

func TestUser(t *testing.T) {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()

	data := GetAllDataWithoutPagination("s4fu2")
	ExportToFile(data, "file-sabtu-x.xlsx")
}
