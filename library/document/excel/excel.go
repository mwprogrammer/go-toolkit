package excel

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/mwprogrammer/go-utilities/models"
	"github.com/xuri/excelize/v2"
)

func Open(path string, logger *slog.Logger) (*excelize.File, bool) {

	is_success := true
	excel_file, err := excelize.OpenFile(path)

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return nil, is_success
	}

	defer func() {
		if excel_err := excel_file.Close(); excel_err != nil {

			logger.Error(err.Error())
			is_success = false
			
		}
	}()

	return excel_file, is_success

}

func Create(sheets []models.ExcelSheet, excel_file string, logger *slog.Logger) (string, bool) {

	is_success := true
	file := excelize.NewFile()
	
	for index, sheet := range sheets {
				
		sheet_number := fmt.Sprintf("Sheet%d", index + 1)
		sheet_index, _ := file.NewSheet(sheet_number)

		if index == 0 {
		
			file.SetActiveSheet(sheet_index)

		}
		
		header_values := make([]interface{}, len(sheet.Headers))

		for index, value := range sheet.Headers {
			header_values[index] = value
		}

		file.SetSheetName(sheet_number, sheet.Name)
		file.SetSheetRow(sheet.Name, sheet.InitialCell, &header_values)
		
		
		for index, row := range sheet.Values {

			cell_values := make([]interface{}, len(row))

			for index, value := range row {
				cell_values[index] = value
			}

			initial_cell_letter := strings.Split(sheet.InitialCell, "")
			initial_cell := fmt.Sprintf("%s%d", initial_cell_letter[0], index + 2)

			file.SetSheetRow(sheet.Name, initial_cell, &cell_values)

		}
		
	}

	if err := file.SaveAs(excel_file); err != nil {

        logger.Error(err.Error())
		is_success = false

		return "", is_success
    }

	file_path := file.Path

	return file_path, is_success

}