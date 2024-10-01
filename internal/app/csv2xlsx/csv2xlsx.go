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
package csv2xlsx

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/kenita8/xlcmd/internal/app/csv2xlsx/config"
	"github.com/kenita8/xlcmd/internal/pkg/excel"
	"github.com/kenita8/xlcmd/internal/pkg/file/csv"
	"github.com/kenita8/xlcmd/internal/pkg/file/tsv"
	"github.com/kenita8/xlcmd/internal/pkg/file/txt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Csv2Xlsx struct {
	log   *zap.Logger
	excel Excel
}

type Excel interface {
	Open(pathname string) error
	NewSheet(name string) error
	PasteTxtFile(txt txt.TxtFiler, sheet string, opt *excel.CellOption) error
	Save() error
	Close()
}

func NewCsv2Xlsx(lc fx.Lifecycle, config config.Config, excel Excel, log *zap.Logger) *Csv2Xlsx {
	c2x := &Csv2Xlsx{log: log, excel: excel}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := c2x.convertExcel(config)
			if err != nil {
				log.Error("failed to convert to excel file", zap.NamedError("err", err))
				return err
			}
			log.Info("completed successfully")
			return nil
		},
	})
	return c2x
}

func newInputFile(pathname string, encoding string) (txt.TxtFiler, error) {
	ext := strings.ToLower(filepath.Ext(pathname))
	var input txt.TxtFiler
	if ext == ".csv" {
		input = csv.NewCsvFile(pathname, encoding)
	} else if ext == ".tsv" {
		input = tsv.NewTsvFile(pathname, encoding)
	} else if ext == ".txt" {
		input = txt.NewTxtFile(pathname, encoding)
	} else {
		return nil, ErrInputFileExtension.Details("ext", ext)
	}
	return input, nil
}

func (c2x *Csv2Xlsx) convertExcel(config config.Config) error {
	targets, encoding, err := config.InputFiles()
	if err != nil {
		return err
	}
	output := config.XlsxFilename()
	opt := config.CellOption()

	err = c2x.excel.Open(output)
	if err != nil {
		return err
	}
	defer c2x.excel.Close()

	for _, target := range targets {
		input, err := newInputFile(target, encoding)
		if err != nil {
			return err
		}
		err = c2x.excel.PasteTxtFile(input, input.Basename(), opt)
		if err != nil {
			return err
		}
	}

	err = c2x.excel.Save()
	if err != nil {
		return err
	}
	return nil
}
