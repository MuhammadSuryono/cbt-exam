package question

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"time"
)

type ListQuestion struct {
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`                     //
	CreatedBy          int       `gorm:"column:created_by" json:"created_by"`                     //
	Description        string    `gorm:"column:description" json:"description"`                   //
	ID                 int64     `gorm:"column:id;primary_key" json:"id"`                         //
	IsPublish          int       `gorm:"column:is_publish" json:"is_publish"`                     //
	QuestionCategoryID int64     `gorm:"column:question_category_id" json:"question_category_id"` //
	SectionID          int64     `gorm:"column:section_id" json:"section_id"`                     //
	TitleQuestion      string    `gorm:"column:title_question" json:"title_question"`             //
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`                     //
	UpdatedBy          int       `gorm:"column:updated_by" json:"updated_by"`                     //
}

// TableName sets the insert table name for this struct type
func (l *ListQuestion) TableName() string {
	return "list_question"
}

func GetListQuestionCategory(categoryId int64) (listQuestion []ListQuestion) {
	_ = db.Connection.Table("list_question").
		Where("question_category_id = ?", categoryId).
		Find(&listQuestion)
	return
}
