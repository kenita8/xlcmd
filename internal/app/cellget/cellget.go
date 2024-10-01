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
package cellget

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenita8/xlcmd/internal/app/cellget/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Cellget struct {
	log   *zap.Logger
	excel Excel
}

type Excel interface {
	Open(filename string) error
	GetCellValue(sheet string, col int, row int) (string, error)
	CoordinatesToCellName(col int, row int, abs ...bool) (string, error)
	CellNameToCoordinates(cell string) (int, int, error)
	Save() error
	Close()
}

func NewCellget(lc fx.Lifecycle, config config.Config, excel Excel, log *zap.Logger) *Cellget {
	chart := &Cellget{log: log, excel: excel}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := chart.cellGet(config)
			if err != nil {
				log.Error("failed to set cell", zap.NamedError("err", err))
				return err
			}
			log.Info("completed successfully")
			return nil
		},
	})
	return chart
}

func (c *Cellget) outputCsv(sheet string, top, bottom, left, right int, format string) error {
	f := "\t"
	if format == "csv" {
		f = ","
	}
	for i := top; i <= bottom; i++ {
		var values []string
		for j := left; j <= right; j++ {
			str, err := c.excel.GetCellValue(sheet, j, i)
			if err != nil {
				return err
			}
			values = append(values, str)
		}
		fmt.Printf("%s\n", strings.Join(values, f))
	}
	return nil
}

func (c *Cellget) outputList(sheet string, top, bottom, left, right int) error {
	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			str, err := c.excel.GetCellValue(sheet, j, i)
			if err != nil {
				return err
			}
			cell, err := c.excel.CoordinatesToCellName(j, i)
			if err != nil {
				return err
			}
			fmt.Printf("%s: %s\n", cell, str)
		}
	}
	return nil
}

func (c *Cellget) cellGet(config config.Config) error {
	topLeft, bottomRight, err := config.Range()
	if err != nil {
		return err
	}
	sheet := config.SheetName()
	output := config.XlsxFilename()
	format, err := config.Format()
	if err != nil {
		return err
	}

	left, top, err := c.excel.CellNameToCoordinates(topLeft)
	if err != nil {
		return err
	}
	right, bottom, err := c.excel.CellNameToCoordinates(bottomRight)
	if err != nil {
		return err
	}
	err = c.excel.Open(output)
	if err != nil {
		return err
	}
	defer c.excel.Close()

	if format == "csv" || format == "tsv" {
		c.outputCsv(sheet, top, bottom, left, right, format)
	} else {
		c.outputList(sheet, top, bottom, left, right)
	}

	return nil
}
