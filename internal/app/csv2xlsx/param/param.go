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
	Input() string
	XlsxFilename() string
	Extension() string
	Depth() int
	DecimalPlaces() int
	Encoding() string
}

type param struct {
	log           *zap.Logger
	input         string
	xlsxFilename  string
	ext           string
	depth         int
	decimalPlaces int
	encoding      string
}

func NewParam(log *zap.Logger) Param {
	return &param{log: log}
}

func (p *param) Parse() {
	input := flag.String("input", ".", "Set input files or directories to convert to Excel.")
	xlsxFilename := flag.String("xlsx", "output.xlsx", "Set output Excel file name.")
	ext := flag.String("ext", "csv,tsv", "Set file extensions to search within input directories. csv, tsv, txt.")
	depth := flag.Int("depth", 0, "Set maximum directory depth for input.")
	decimalPlaces := flag.Int("decimal-places", 2, "Set number of decimal places for numbers.")
	encoding := flag.String("encoding", "UTF-8", "Set input file encoding(IANA-registered name).")
	flag.Parse()
	p.input = *input
	p.xlsxFilename = *xlsxFilename
	p.ext = *ext
	p.depth = *depth
	p.decimalPlaces = *decimalPlaces
	p.encoding = *encoding
}

func (p *param) Input() string {
	return p.input
}

func (p *param) XlsxFilename() string {
	return p.xlsxFilename
}

func (p *param) Extension() string {
	return p.ext
}

func (p *param) Depth() int {
	return p.depth
}

func (p *param) DecimalPlaces() int {
	return p.decimalPlaces
}

func (p *param) Encoding() string {
	return p.encoding
}
