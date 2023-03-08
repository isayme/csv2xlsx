package csv2xlsx

import (
	"encoding/csv"
	"io"
	"os"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

const (
	DEFAULT_SHEET_NAME = "Sheet1"
)

func Convert(options *Options) error {
	options.Ajust()

	file, err := os.Open(options.InputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	xlsxFile := excelize.NewFile()
	defer xlsxFile.Close()

	streamWriter, err := xlsxFile.NewStreamWriter(DEFAULT_SHEET_NAME)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(file)
	comma, _ := utf8.DecodeRuneInString(options.Comma)
	csvReader.Comma = comma
	csvReader.FieldsPerRecord = -1

	rowID := 0
	for {
		rowID++

		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		row := make([]interface{}, len(record))
		for i, v := range record {
			row[i] = v
		}

		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		streamWriter.SetRow(cell, row)
	}

	if err := streamWriter.Flush(); err != nil {
		return err
	}

	excelizeOptions := excelize.Options{
		Password: options.OutputFilePassword,
	}

	err = xlsxFile.SaveAs(options.OutputFilePath, excelizeOptions)
	if err != nil {
		return err
	}

	return nil
}
