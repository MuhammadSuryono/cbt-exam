package pdf

import (
	"fmt"
	"testing"
)

func TestPdfGenerator(t *testing.T) {
	r := NewRequestPdf("")
	templatePath := "template.html"
	output := "output.pdf"

	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(output)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
