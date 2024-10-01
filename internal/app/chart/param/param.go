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
	ConfigFilename() string
	XlsxFilename() string
}

type param struct {
	log            *zap.Logger
	configFilename string
	xlsxFilename   string
}

func NewParam(log *zap.Logger) Param {
	return &param{log: log}
}

func (p *param) Parse() {
	configFilename := flag.String("config", "chart.yml", "Set Excel chart configuration file.")
	xlsxFilename := flag.String("xlsx", "output.xlsx", "Set output Excel file name.")
	flag.Parse()
	p.configFilename = *configFilename
	p.xlsxFilename = *xlsxFilename
}

func (p *param) ConfigFilename() string {
	return p.configFilename
}

func (p *param) XlsxFilename() string {
	return p.xlsxFilename
}
