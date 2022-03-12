package type_exam

import "github.com/MuhammadSuryono/go-helper/db"

type TypeExam struct {
	ID   int64  `gorm:"column:id;primary_key" json:"id"` //
	Slug string `gorm:"column:slug" json:"slug"`         //
	Type string `gorm:"column:type" json:"type"`         //
}

// TableName sets the insert table name for this struct type
func (t *TypeExam) TableName() string {
	return "type_exam"
}

func GetAllTypeExam() (typeExam []TypeExam) {
	_ = db.Connection.Table("type_exam").Find(&typeExam)
	return
}
