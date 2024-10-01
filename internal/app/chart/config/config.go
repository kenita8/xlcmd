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
	"io"
	"os"

	"github.com/kenita8/xlcmd/internal/app/chart/param"
	"github.com/kenita8/xlcmd/internal/pkg/excel/excelize"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config interface {
	ExelizeChartOption() ([]*excelize.ExcelizeChartOption, error)
	XlsxFilename() string
}

type config struct {
	param param.Param
	log   *zap.Logger
}

func NewConfig(param param.Param, log *zap.Logger) Config {
	return &config{param: param, log: log}
}

func (c *config) ExelizeChartOption() ([]*excelize.ExcelizeChartOption, error) {
	configFilename := c.param.ConfigFilename()
	f, err := os.Open(configFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var opts excelize.ChartOptions
	err = yaml.Unmarshal(data, &opts)
	if err != nil {
		return nil, err
	}
	return excelize.ConvertExelizeChartOption(opts.ChartOptions)
}

func (c *config) XlsxFilename() string {
	return c.param.XlsxFilename()
}
