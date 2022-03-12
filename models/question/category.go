package question

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"time"
)

type CategoryQuestion struct {
	Code       string    `gorm:"column:code" json:"code"`                 //
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`     //
	ID         int64     `gorm:"column:id;primary_key" json:"id"`         //
	Name       string    `gorm:"column:name" json:"name"`                 //
	Sequence   int       `gorm:"column:sequence" json:"sequence"`         //
	TypeExamID int64     `gorm:"column:type_exam_id" json:"type_exam_id"` //
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (c *CategoryQuestion) TableName() string {
	return "category_question"
}

func GetCategoryByTypeExam(typeExamId int64) (category CategoryQuestion) {
	_ = db.Connection.Table(category.TableName()).Where("type_exam_id = ?", typeExamId).First(&category)
	return
}
