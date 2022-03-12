package excel

type TemplateQuestion struct {
	Question         string `json:"question"`
	Option1          string `json:"option_1"`
	Option2          string `json:"option_2"`
	Option3          string `json:"option_3"`
	Option4          string `json:"option_4"`
	Option5          string `json:"option_5"`
	Answer           string `json:"answer"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
	FilenameGroup    string `json:"filename_group"`
	QuestionGroup    string `json:"question_group"`
}

const (
	Question         = 0
	Option1          = 1
	Option2          = 2
	Option3          = 3
	Option4          = 4
	Option5          = 5
	Answer           = 6
	GroupName        = 7
	GroupDescription = 8
	FilenameGroup    = 9
	QuestionGroup    = 10
)
