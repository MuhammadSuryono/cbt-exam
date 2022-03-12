package user

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

var headers = []string{
	"No",
	"Nama Sesi",
	"Nama Lengkap",
	"Tempat Lahir",
	"Tanggal Lahir",
	"Email",
	"Nomor Telepon",
	"Alamat",
	"Nomor Registrasi",
	"Tanggal Registrasi",
	"Total Benar",
	"Total Score",
}

func ExportToFile(data []UserParticipantResponseWithTypeExam) {
	xlsx := excelize.NewFile()
	setHeader(xlsx)
	setDataBody(xlsx, data)
	err := xlsx.SaveAs("file_senin.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func setHeader(xlsx *excelize.File) {
	sheetName := xlsx.GetSheetName(1)
	for i, header := range headers {
		xlsx.SetCellValue(sheetName, excelize.ToAlphaString(i)+"1", header)
	}
}

func setDataBody(xlsx *excelize.File, data []UserParticipantResponseWithTypeExam) {
	sheetName := xlsx.GetSheetName(1)
	totalColumnHeader := len(headers)
	for iRow, datum := range data {
		for iColumn, _ := range headers {
			strColumn := strconv.Itoa(iRow + 2)

			switch iColumn {
			case 0:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, iRow+1)
			case 1:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.SessionName)
			case 2:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.FullName)
			case 3:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.PlaceOfBirth)
			case 4:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.DateOfBirth)
			case 5:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.Email)
			case 6:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.PhoneNumber)
			case 7:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.Address)
			case 8:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.NumberOfRegister)
			case 9:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.DateRegister)
			case 10:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.TotalCorrection)
			case 11:
				xlsx.SetCellValue(sheetName, excelize.ToAlphaString(iColumn)+strColumn, datum.TotalScore)
			default:
				fmt.Println("Unknown column:", iColumn)
			}
		}

		for iType, result := range datum.TypeExamResult {
			nextColumn := totalColumnHeader + iType
			xlsx.SetCellValue(sheetName, excelize.ToAlphaString(nextColumn)+"1", result.Name)
		}

		for iType, _ := range datum.TypeExamResult {
			nextColumn := totalColumnHeader + iType
			strColumn := strconv.Itoa(iRow + 2)
			xlsx.SetCellValue(sheetName, excelize.ToAlphaString(nextColumn)+strColumn, datum.TypeExamResult[iType].TotalScore)
		}
	}
}
