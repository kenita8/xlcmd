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
package main

import (
	"context"
	"os"

	"github.com/kenita8/xlcmd/internal/app/csv2xlsx"
	"github.com/kenita8/xlcmd/internal/app/csv2xlsx/config"
	"github.com/kenita8/xlcmd/internal/app/csv2xlsx/param"
	"github.com/kenita8/xlcmd/internal/pkg/excel"
	"github.com/kenita8/xlcmd/internal/pkg/log"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var (
	Version string = ""
)

func main() {
	app := fx.New(
		fx.WithLogger(func(*zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: zap.NewNop()}
		}),
		fx.Provide(
			fx.Annotate(param.NewParam, fx.As(new(param.Param))),
			fx.Annotate(config.NewConfig, fx.As(new(config.Config))),
			fx.Annotate(excel.NewExcel, fx.As(new(csv2xlsx.Excel))),
			csv2xlsx.NewCsv2Xlsx,
			log.NewLog,
		),
		fx.Invoke(func(param param.Param) {
			param.Parse()
		}),
		fx.Invoke(func(log *zap.Logger) {
			log.Info("starting process", zap.String("version", Version))
		}),
		fx.Invoke(func(*csv2xlsx.Csv2Xlsx) {}),
	)
	err := app.Start(context.Background())
	if err != nil {
		os.Exit(1)
	}
	app.Stop(context.Background())
}
