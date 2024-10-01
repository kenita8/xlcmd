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
package param

import (
	"flag"

	"go.uber.org/zap"
)

type Param interface {
	Parse()
	XlsxFilename() string
	SheetName() string
	Range() string
	Format() string
}

type param struct {
	log          *zap.Logger
	xlsxFilename string
	sheetName    string
	rangeStr     string
	format       string
}

func NewParam(log *zap.Logger) Param {
	return &param{log: log}
}

func (p *param) Parse() {
	xlsxFilename := flag.String("xlsx", "output.xlsx", "Set output Excel file name.")
	sheetName := flag.String("sheet", "Sheet1", "Set the sheet name for the graph.")
	rangeStr := flag.String("range", "A1", "Set the range of cells to process.")
	format := flag.String("format", "csv", "Set the output format(csv, tsv, list).")

	flag.Parse()

	p.xlsxFilename = *xlsxFilename
	p.sheetName = *sheetName
	p.rangeStr = *rangeStr
	p.format = *format
}

func (p *param) XlsxFilename() string {
	return p.xlsxFilename
}

func (p *param) SheetName() string {
	return p.sheetName
}

func (p *param) Range() string {
	return p.rangeStr
}

func (p *param) Format() string {
	return p.format
}
