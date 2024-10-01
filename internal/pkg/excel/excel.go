// Copyright 2024 kenita8
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package excel

import (
	"fmt"
	"io"
	"io/fs"
	"strconv"

	"github.com/kenita8/xlcmd/internal/pkg/excel/excelize"
	"github.com/kenita8/xlcmd/internal/pkg/file"
	"github.com/kenita8/xlcmd/internal/pkg/file/txt"
	"go.uber.org/zap"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type Filer interface {
	Stat(name string) (fs.FileInfo, error)
}

type CellOption struct {
	DecimalPlaces int
}

var (
	filer     Filer              = &file.File{}
	excelizer excelize.Excelizer = &excelize.Excelize{}
)

type Excel struct {
	xlFile   excelize.ExcelizeFiler
	log      *zap.Logger
	pathname string
	new      bool
}

func NewExcel(log *zap.Logger) *Excel {
	return &Excel{
		log: log,
	}
}

func (e *Excel) newFile() {
	e.log.Info("new file", zap.String("path", e.pathname))
	e.xlFile = excelizer.NewFile()
	e.new = true
}

func (e *Excel) openFile() error {
	e.log.Info("open file", zap.String("path", e.pathname))
	xlFile, err := excelizer.OpenFile(e.pathname)
	if err != nil {
		return ErrOpenXlsxFile.Wrap(err)
	}
	e.xlFile = xlFile
	e.new = false
	return nil
}

func (e *Excel) Open(pathname string) error {
	e.pathname = pathname
	_, err := filer.Stat(pathname)
	if err != nil {
		e.newFile()
		err = nil
	} else {
		err = e.openFile()
	}
	return err
}

func (e *Excel) NewSheet(name string) error {
	if e.xlFile == nil {
		return ErrNotOpened
	}
	_, err := e.xlFile.NewSheet(name)
	if err != nil {
		return ErrNewSheet.Details("sheet", name).Wrap(err)
	}
	return nil
}

func (e *Excel) SetCellValue(value string, sheet string, col int, row int, opt *CellOption) error {
	if e.xlFile == nil {
		return ErrNotOpened
	}
	cell, err := excelizer.CoordinatesToCellName(col, row)
	if err != nil {
		return ErrConvertCellName.Details("col", col, "row", row).Wrap(err)
	}
	valuef, err := strconv.ParseFloat(value, 64)
	if err != nil {
		err = e.xlFile.SetCellValue(sheet, cell, value)
	} else {
		if opt != nil {
			err = e.xlFile.SetCellFloat(sheet, cell, valuef, opt.DecimalPlaces, 64)
		} else {
			err = e.xlFile.SetCellFloat(sheet, cell, valuef, -1, 64)
		}
	}
	if err != nil {
		return ErrSetCellValue.Details("sheet", sheet, "cell", cell, "data", value).Wrap(err)
	}
	e.log.Info("replace", zap.String("sheet", sheet), zap.String("cell", cell), zap.String("text", value))
	return nil
}

func (e *Excel) GetCellValue(sheet string, col int, row int) (string, error) {
	if e.xlFile == nil {
		return "", ErrNotOpened
	}
	cell, err := excelizer.CoordinatesToCellName(col, row)
	if err != nil {
		return "", ErrConvertCellName.Details("col", col, "row", row).Wrap(err)
	}

	value, err := e.xlFile.GetCellValue(sheet, cell)
	if err != nil {
		return "", ErrGetCellValue.Details("sheet", sheet, "col", col, "row", row).Wrap(err)
	}
	return value, nil
}

func (e *Excel) PasteTxtFile(txt txt.TxtFiler, sheet string, opt *CellOption) error {
	if e.xlFile == nil {
		return ErrNotOpened
	}
	_, err := e.xlFile.NewSheet(sheet)
	if err != nil {
		return err
	}
	err = txt.OpenReadMode()
	if err != nil {
		return err
	}
	defer txt.Close()
	row := 1
	for {
		values, err := txt.ReadOneLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ErrReadInputFile.Details("file", txt.Filename()).Wrap(err)
		}
		col := 1
		for _, value := range values {
			e.SetCellValue(value, sheet, col, row, opt)
			col++
		}
		row += 1
	}
	e.log.Info("add sheet", zap.String("sheet", sheet), zap.String("src", txt.Filename()))
	return nil
}

func (e *Excel) AddChart(chart *excelize.ExcelizeChartOption) error {
	if e.xlFile == nil {
		return ErrNotOpened
	}
	e.xlFile.NewSheet(chart.Sheet)
	err := e.xlFile.AddChart(chart.Sheet, chart.Cell, chart.Chart, chart.Combo...)
	if err != nil {
		return ErrAddChart.Wrap(err)
	}
	e.log.Info("add chart", zap.String("sheet", chart.Sheet), zap.String("pos", chart.Cell))
	return nil
}

func (e *Excel) MaxRow(sheet string) (int, error) {
	if e.xlFile == nil {
		return 0, ErrNotOpened
	}
	row := 0
	for {
		value, err := e.GetCellValue(sheet, 1, row+1)
		if err != nil {
			return 0, err
		}
		if len(value) <= 0 {
			break
		}
		row++
	}
	return row, nil
}

func (e *Excel) MaxCol(sheet string) (int, error) {
	if e.xlFile == nil {
		return 0, ErrNotOpened
	}
	col := 0
	for {
		value, err := e.GetCellValue(sheet, col+1, 1)
		if err != nil {
			return 0, err
		}
		if len(value) <= 0 {
			break
		}
		col++
	}
	return col, nil
}

func (e *Excel) CoordinatesToCellName(col int, row int, abs ...bool) (string, error) {
	return excelizer.CoordinatesToCellName(col, row, abs...)
}

func (e *Excel) CellNameToCoordinates(cell string) (int, int, error) {
	return excelizer.CellNameToCoordinates(cell)
}

func (e *Excel) RangeString(sheet string, topCol int, topRow int, bottomCol int, bottomRow int) (string, error) {
	top, err := excelizer.CoordinatesToCellName(topCol, topRow)
	if err != nil {
		return "", ErrConvertCellName.Details("col", topCol, "row", topRow).Wrap(err)
	}
	bottom, err := excelizer.CoordinatesToCellName(bottomCol, bottomRow)
	if err != nil {
		return "", ErrConvertCellName.Details("col", bottomCol, "row", bottomRow).Wrap(err)
	}
	return fmt.Sprintf("%s!%s:%s", sheet, top, bottom), nil
}

func (e *Excel) GetSheetList() []string {
	return e.xlFile.GetSheetList()
}

func (e *Excel) Save() error {
	if e.xlFile == nil {
		return ErrNotOpened
	}
	var err error
	if e.new {
		e.xlFile.DeleteSheet("Sheet1")
		err = e.xlFile.SaveAs(e.pathname)
	} else {
		err = e.xlFile.Save()
	}
	if err != nil {
		return ErrSaveAsFile.Details("path", e.pathname).Wrap(err)
	}
	e.log.Info("saved", zap.String("path", e.pathname))
	return nil
}

func (e *Excel) Close() {
	if e.xlFile == nil {
		return
	}
	e.log.Info("closed")
	e.xlFile.Close()
}
