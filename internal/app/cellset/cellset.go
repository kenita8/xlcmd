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
package cellset

import (
	"context"

	"github.com/kenita8/xlcmd/internal/app/cellset/config"
	"github.com/kenita8/xlcmd/internal/pkg/excel"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Cellset struct {
	log   *zap.Logger
	excel Excel
}

type Excel interface {
	Open(filename string) error
	GetCellValue(sheet string, col int, row int) (string, error)
	SetCellValue(value string, sheet string, col int, row int, opt *excel.CellOption) error
	CoordinatesToCellName(col int, row int, abs ...bool) (string, error)
	CellNameToCoordinates(cell string) (int, int, error)
	Save() error
	Close()
}

func NewCellset(lc fx.Lifecycle, config config.Config, excel Excel, log *zap.Logger) *Cellset {
	chart := &Cellset{log: log, excel: excel}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := chart.cellSet(config)
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

func (c *Cellset) cellSet(config config.Config) error {
	re, err := config.ReplaceConfig()
	if err != nil {
		return err
	}
	topLeft, bottomRight, err := config.Range()
	if err != nil {
		return err
	}
	sheet := config.SheetName()
	output := config.XlsxFilename()

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

	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			str, err := c.excel.GetCellValue(sheet, j, i)
			if err != nil {
				return err
			}
			replaced, err := re.Replace(str)
			if err != nil {
				return err
			}
			err = c.excel.SetCellValue(replaced, sheet, j, i, nil)
			if err != nil {
				return err
			}
		}
	}

	err = c.excel.Save()
	if err != nil {
		return err
	}
	return nil
}
