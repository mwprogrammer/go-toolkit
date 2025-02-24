package csv

import (
	"encoding/csv"
	"log/slog"
	"os"

	"github.com/xuri/excelize/v2"
)

func Read(path string, numberOfRecords int, logger *slog.Logger) ([][]string, bool) {

	is_success := true
	file, err := os.Open(path)

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return nil, is_success
	}

	defer file.Close()

	reader := csv.NewReader(file)
	
	if numberOfRecords > 0 {	
		reader.FieldsPerRecord = numberOfRecords
	}
		
	data, err := reader.ReadAll()

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return nil, is_success
	}
	
	return data, is_success

}

func CreateFromCollection(headers []string, data [][]string, path string, logger *slog.Logger) (string, bool) {

	is_success := true
	csv_file, err := os.Create(path)

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return "", is_success
	}

    defer csv_file.Close()

    writer := csv.NewWriter(csv_file)
    defer writer.Flush()

	writer.Write(headers)

	for _, row := range data {
		writer.Write(row)
	}

	return csv_file.Name(), is_success

}

func CreateFromExcelSheet(excel *excelize.File, sheet_name string, path string, logger *slog.Logger) (string, bool) {

	is_success := true
	sheet_rows, err := excel.GetRows(sheet_name)

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return "", is_success
	}

	csv_file, csv_err := os.Create(path)

	if csv_err != nil {

		logger.Error(csv_err.Error())
		is_success = false

		return "", is_success
	}

	defer func() {

		if csv_err := csv_file.Close(); csv_err != nil {

			logger.Error(csv_err.Error())
			is_success = false

		}
	}()

	csv_writer := csv.NewWriter(csv_file)

	var writer_err error = csv_writer.WriteAll(sheet_rows)

	if writer_err != nil {

		logger.Error(writer_err.Error())
		is_success = false

	}

	return csv_file.Name(), is_success
}