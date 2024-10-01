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
package config

import (
	"regexp"

	"github.com/kenita8/xlcmd/internal/app/ezchart/param"
	"github.com/kenita8/xlcmd/internal/pkg/excel/excelize"
	"go.uber.org/zap"
)

type Config interface {
	ChartType() excelize.ChartType
	SheetName() (*regexp.Regexp, error)
	XlsxFilename() string
}

type config struct {
	param param.Param
	log   *zap.Logger
}

func NewConfig(param param.Param, log *zap.Logger) Config {
	return &config{param: param, log: log}
}

func (c *config) ChartType() excelize.ChartType {
	return excelize.ChartType(c.param.ChartType())
}

func (c *config) SheetName() (*regexp.Regexp, error) {
	sheetName := c.param.SheetName()
	re, err := regexp.Compile(sheetName)
	if err != nil {
		return nil, ErrCompileRegexp.Details("sheet", sheetName)
	}
	return re, nil
}

func (c *config) XlsxFilename() string {
	return c.param.XlsxFilename()
}
