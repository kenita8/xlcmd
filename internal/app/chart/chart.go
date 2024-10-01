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
package chart

import (
	"context"

	"github.com/kenita8/xlcmd/internal/app/chart/config"
	"github.com/kenita8/xlcmd/internal/pkg/excel/excelize"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Chart struct {
	log   *zap.Logger
	excel Excel
}

type Excel interface {
	Open(filename string) error
	AddChart(chart *excelize.ExcelizeChartOption) error
	Save() error
	Close()
}

func NewChart(lc fx.Lifecycle, config config.Config, excel Excel, log *zap.Logger) *Chart {
	chart := &Chart{log: log, excel: excel}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := chart.generateChart(config)
			if err != nil {
				log.Error("failed to generate charts", zap.NamedError("err", err))
				return err
			}
			log.Info("completed successfully")
			return nil
		},
	})
	return chart
}

func (c *Chart) generateChart(config config.Config) error {
	opts, err := config.ExelizeChartOption()
	if err != nil {
		return err
	}
	output := config.XlsxFilename()

	err = c.excel.Open(output)
	if err != nil {
		return err
	}
	defer c.excel.Close()

	for _, opt := range opts {
		err = c.excel.AddChart(opt)
		if err != nil {
			return err
		}
	}

	err = c.excel.Save()
	if err != nil {
		return err
	}
	return nil
}
