package exam

import (
	"fmt"
	"github.com/MuhammadSuryono/go-helper/db"
	"math"
	"time"
)

type ExamResult struct {
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`             //
	ExamQuestionID int64     `gorm:"column:exam_question_id" json:"exam_question_id"` //
	ID             int64     `gorm:"column:id;primary_key" json:"id"`                 //
	NumberRegister string    `gorm:"column:number_register" json:"number_register"`   //
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`             //
	Value          int       `gorm:"column:value" json:"value"`                       //
}

// TableName sets the insert table name for this struct type
func (e *ExamResult) TableName() string {
	return "exam_result"
}

func CheckAnswer(registerNumber string, examQuestionId int64, answerCorrect string) bool {
	var result ExamResult
	db.Connection.Table(result.TableName()).Where("number_register = ? AND exam_question_id = ?", registerNumber, examQuestionId).First(&result)
	if result.ID != 0 {
		fmt.Println("Check Answer", convAnswer(result.Value), answerCorrect)
		return convAnswer(result.Value) == answerCorrect
	}
	return false
}

func CheckAnswerNew(registerNumber string, examQuestionId int64, answerCorrect string) bool {
	var result ExamResult
	db.Connection.Table(result.TableName()).Where("number_register = ? AND exam_question_id = ?", registerNumber, examQuestionId).First(&result)
	if result.ID != 0 {
		fmt.Println("Check Answer", convAnswer(result.Value), answerCorrect)
		return convAnswer(result.Value) == answerCorrect
	}
	return false
}

func convAnswer(param int) string {
	mapAnswer := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	return mapAnswer[param]
}

func CountTotalScore(totalCount int, totalType int) int {
	return int(math.Ceil(float64((totalCount / totalType) * 10)))
}
