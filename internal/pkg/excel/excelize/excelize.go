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
package excelize

import (
	"github.com/xuri/excelize/v2"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type ExcelizeFiler interface {
	NewSheet(sheet string) (int, error)
	GetSheetList() (list []string)
	SetActiveSheet(index int)
	SetCellValue(sheet, cell string, value interface{}) error
	SetCellFloat(sheet, cell string, value float64, precision, bitSize int) error
	DeleteSheet(sheet string) error
	GetCellValue(sheet, cell string, opts ...excelize.Options) (string, error)
	SaveAs(filename string, opts ...excelize.Options) error
	Save(opts ...excelize.Options) error
	SetCellStyle(sheet, topLeftCell, bottomRightCell string, styleID int) error
	AddChart(sheet, cell string, chart *excelize.Chart, combo ...*excelize.Chart) error
	Cols(sheet string) (*excelize.Cols, error)
	Rows(sheet string) (*excelize.Rows, error)
	Close() error
}

type Excelizer interface {
	NewFile() ExcelizeFiler
	CoordinatesToCellName(col int, row int, abs ...bool) (string, error)
	CellNameToCoordinates(cell string) (int, int, error)
	OpenFile(filename string, opts ...excelize.Options) (ExcelizeFiler, error)
}

type Excelize struct {
}

func (e *Excelize) NewFile() ExcelizeFiler {
	return excelize.NewFile()
}

func (e *Excelize) CoordinatesToCellName(col int, row int, abs ...bool) (string, error) {
	return excelize.CoordinatesToCellName(col, row, abs...)
}

func (e *Excelize) CellNameToCoordinates(cell string) (int, int, error) {
	return excelize.CellNameToCoordinates(cell)
}

func (e *Excelize) OpenFile(filename string, opts ...excelize.Options) (ExcelizeFiler, error) {
	return excelize.OpenFile(filename, opts...)
}
