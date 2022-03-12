package excel

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	excel := NewExcelHandler("Template Soal.xlsx")
	rows, _ := excel.ReadExcel()

	fmt.Println(rows)

	//for _, v := range rows {
	//	dataGroup := group.CreateGroupQuestion(60, v.GroupName, v.GroupDescription, v.FilenameGroup, v.QuestionGroup)
	//	question.AddExamQuestion(v, dataGroup.Id, 60)
	//}

}
