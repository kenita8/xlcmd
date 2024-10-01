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
	ChartType() string
	SheetName() string
	XlsxFilename() string
}

type param struct {
	log          *zap.Logger
	chartType    string
	sheetName    string
	xlsxFilename string
}

func NewParam(log *zap.Logger) Param {
	return &param{log: log}
}

func (p *param) Parse() {
	chartType := flag.String("type", `Line`, "Set chart type to create.")
	sheetName := flag.String("sheet", `.+\.(csv|tsv)$`, "Set the sheet name for the graph. Regex allowed.")
	xlsxFilename := flag.String("xlsx", "output.xlsx", "Set output Excel file name.")
	flag.Parse()
	p.chartType = *chartType
	p.sheetName = *sheetName
	p.xlsxFilename = *xlsxFilename
}

func (p *param) ChartType() string {
	return p.chartType
}

func (p *param) SheetName() string {
	return p.sheetName
}

func (p *param) XlsxFilename() string {
	return p.xlsxFilename
}
