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
	"strings"

	"github.com/kenita8/xlcmd/internal/app/cellget/param"
	"go.uber.org/zap"
)

type Config interface {
	XlsxFilename() string
	SheetName() string
	Range() (string, string, error)
	Format() (string, error)
}

type config struct {
	param param.Param
	log   *zap.Logger
}

func NewConfig(param param.Param, log *zap.Logger) Config {
	return &config{param: param, log: log}
}

func (c *config) Range() (string, string, error) {
	rangeStr := c.param.Range()
	parts := strings.Split(rangeStr, ":")
	num := len(parts)
	var topLeft string
	var bottomRight string
	if num == 1 {
		topLeft = parts[0]
		bottomRight = parts[0]
	} else if num == 2 {
		topLeft = parts[0]
		bottomRight = parts[1]
	} else {
		return "", "", ErrInvalidRange.Details("range", rangeStr)
	}
	return topLeft, bottomRight, nil
}

func (c *config) SheetName() string {
	return c.param.SheetName()
}

func (c *config) XlsxFilename() string {
	return c.param.XlsxFilename()
}

func (c *config) Format() (string, error) {
	format := c.param.Format()
	if format != "csv" && format != "tsv" && format != "list" {
		return "", ErrInvalidFormat
	}
	return format, nil
}
