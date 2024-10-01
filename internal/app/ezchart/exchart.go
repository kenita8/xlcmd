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
package ezchart

import (
	"context"
	"regexp"

	"github.com/kenita8/xlcmd/internal/app/ezchart/config"
	"github.com/kenita8/xlcmd/internal/pkg/excel/excelize"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type EzChart struct {
	log   *zap.Logger
	excel Excel
}

type Excel interface {
	Open(filename string) error
	GetSheetList() []string
	MaxRow(sheet string) (int, error)
	MaxCol(sheet string) (int, error)
	RangeString(sheet string, topCol int, topRow int, bottomCol int, bottomRow int) (string, error)
	GetCellValue(sheet string, col int, row int) (string, error)
	CoordinatesToCellName(col int, row int, abs ...bool) (string, error)
	AddChart(chart *excelize.ExcelizeChartOption) error
	Save() error
	Close()
}

func NewEzChart(lf fx.Lifecycle, config config.Config, excel Excel, log *zap.Logger) *EzChart {
	ezChart := &EzChart{log: log, excel: excel}
	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := ezChart.generateEzChart(config)
			if err != nil {
				log.Error("failed to generate charts", zap.NamedError("err", err))
				return err
			}
			log.Info("completed successfully")
			return nil
		},
	})
	return ezChart
}

func (x *EzChart) GenerateAutoGraph(chartType excelize.ChartType, sheet string) error {
	graphRow := 2
	graphCol := 2
	graphRowGrowth := 14
	graphColGrowth := 8
	maxCol, err := x.excel.MaxCol(sheet)
	if err != nil {
		return err
	}
	maxRow, err := x.excel.MaxRow(sheet)
	if err != nil {
		return err
	}
	for col := 2; col <= maxCol; col++ {
		graphCell, err := x.excel.CoordinatesToCellName(graphCol, graphRow)
		if err != nil {
			return err
		}
		Title, err := x.excel.GetCellValue(sheet, col, 1)
		if err != nil {
			return err
		}
		valueRange, err := x.excel.RangeString(sheet, col, 2, col, maxRow)
		if err != nil {
			return err
		}
		chart := &excelize.ChartOption{
			Sheet: sheet,
			Cell:  graphCell,
			Chart: &excelize.Chart{
				Type:   excelize.Line,
				Series: []excelize.ChartSeries{{Values: valueRange}},
				Title:  []excelize.RichTextRun{{Text: Title}},
			},
			Combo: nil,
		}
		chartOpt, err := chart.ConvertExcelizeOption()
		if err != nil {
			return err
		}
		err = x.excel.AddChart(chartOpt)
		if err != nil {
			return err
		}
		if graphCol >= 18 {
			graphRow = graphRow + graphRowGrowth
			graphCol = 2
		} else {
			graphCol = graphCol + graphColGrowth
		}
	}
	return nil
}

func (x *EzChart) GenerateAutoGraphs(chartType excelize.ChartType, targetSheetName *regexp.Regexp) error {
	sheetNames := x.excel.GetSheetList()
	for _, sheetName := range sheetNames {
		if !targetSheetName.MatchString(sheetName) {
			continue
		}
		err := x.GenerateAutoGraph(chartType, sheetName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (x *EzChart) generateEzChart(config config.Config) error {
	sheetName, err := config.SheetName()
	if err != nil {
		return err
	}
	chartType := config.ChartType()
	xlsxFilename := config.XlsxFilename()

	err = x.excel.Open(xlsxFilename)
	if err != nil {
		return err
	}
	defer x.excel.Close()
	err = x.GenerateAutoGraphs(chartType, sheetName)
	if err != nil {
		return err
	}
	err = x.excel.Save()
	if err != nil {
		return err
	}
	return nil
}
