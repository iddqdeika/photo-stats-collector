package table_writer

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"photo-stats-collector/definitions"
)

const ()

func New(filename string, outputSheetName string) definitions.TableWriter {
	return &tableWriter{
		filename:        filename,
		outputSheetName: outputSheetName,
	}
}

type tableWriter struct {
	filename        string
	outputSheetName string
}

func (t *tableWriter) WriteTable(data [][]string) error {
	f, err := excelize.OpenFile(t.filename)
	if err != nil {
		return fmt.Errorf("cant open file %v, error: %v", t.filename, err)
	}
	f.DeleteSheet(t.outputSheetName)
	f.NewSheet(t.outputSheetName)
	for i, row := range data {
		for j, val := range row {
			rowN := i + 1
			columnN := j + 1
			cell, err := excelize.CoordinatesToCellName(columnN, rowN)
			if err != nil {
				return fmt.Errorf("cant create cell name for coords c %v r %v, error: %v", rowN, columnN, err)
			}
			err = f.SetCellValue(t.outputSheetName, cell, val)
			if err != nil {
				return fmt.Errorf("cant set cell value for %v, error: %v", cell, err)
			}
		}
	}
	err = f.Save()
	if err != nil {
		return fmt.Errorf("cant save excel file with name: %v, error: %v", t.filename, err)
	}
	return nil
}
