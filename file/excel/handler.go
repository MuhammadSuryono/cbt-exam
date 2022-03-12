package excel

type IExcel interface {
	ReadExcel() ([]TemplateQuestion, error)
}

type ExcelHandler struct {
	Filename string
}

func NewExcelHandler(filename string) IExcel {
	return &ExcelHandler{
		Filename: filename,
	}
}
