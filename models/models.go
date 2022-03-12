package models

import "time"

type CommonResponse struct {
	Code      int         `json:"code"`
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type GroupQuestion struct {
	Id             int64     `json:"id"`
	ListQuestionId int64     `json:"list_question_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	FilenameGroup  string    `json:"filename_group"`
	QuestionGroup  string    `json:"question_group"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ExamQuestion struct {
	Id              int64     `json:"id"`
	GroupQuestionId int64     `json:"group_question_id"`
	ListQuestionId  int64     `json:"list_question_id"`
	Question        string    `json:"question"`
	Option1         string    `json:"option_1" gorm:"column:option_1"`
	Option2         string    `json:"option_2" gorm:"column:option_2"`
	Option3         string    `json:"option_3" gorm:"column:option_3"`
	Option4         string    `json:"option_4" gorm:"column:option_4"`
	Option5         string    `json:"option_5" gorm:"column:option_5"`
	Answer          string    `json:"answer"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Uri struct {
	ListQuestionId int64  `json:"list_question_id" uri:"list_question_id"`
	Filename       string `uri:"filename"`
	Id             int64  `json:"id" uri:"id"`
}

type QueryRequest struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
