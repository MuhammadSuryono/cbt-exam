package question

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"gtihub.com/MuhammadSuryono/cbt-uploader/file/excel"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/exam"
	"strconv"
)

func AddExamQuestion(dataExcel excel.TemplateQuestion, groupId int64, listQuestionId int64) {
	i, _ := strconv.Atoi(dataExcel.Answer)
	dataExam := models.ExamQuestion{
		GroupQuestionId: groupId,
		ListQuestionId:  listQuestionId,
		Question:        dataExcel.Question,
		Option1:         dataExcel.Option1,
		Option2:         dataExcel.Option2,
		Option3:         dataExcel.Option3,
		Option4:         dataExcel.Option4,
		Option5:         dataExcel.Option5,
		Answer:          i,
	}

	_ = db.Connection.Table("exam_question").Create(&dataExam)
}

func GetTotalCorrect(registerNumber string, typeId int64) (totalCorrect int) {
	category := GetCategoryByTypeExam(typeId)
	listQuestions := GetListQuestionCategory(category.ID)

	for _, question := range listQuestions {
		total := CountTotalCorrect(question.ID, registerNumber)
		totalCorrect += total
	}
	return
}

func GetExamQuestionByListQuestion(listQuestionId int64) (examQuestion []models.ExamQuestion) {
	db.Connection.Table("exam_question").
		Where("list_question_id = ?", listQuestionId).
		Find(&examQuestion)

	return

}

func CountTotalCorrect(lisQuestionId int64, registerNumber string) int {
	var results []exam.ExamResult
	_ = db.Connection.Table("exam_result a").Select("a.*").
		Joins("JOIN exam_question b ON a.exam_question_id = b.id AND a.value = b.answer").
		Where("b.list_question_id = ? AND a.number_register = ?", lisQuestionId, registerNumber).
		Find(&results)

	return len(results)
}
