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
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/kenita8/xlcmd/internal/app/csv2xlsx/param"
	"github.com/kenita8/xlcmd/internal/pkg/excel"
	"go.uber.org/zap"
)

type Config interface {
	InputFiles() ([]string, string, error)
	CellOption() *excel.CellOption
	XlsxFilename() string
}

type config struct {
	param param.Param
	log   *zap.Logger
}

func NewConfig(param param.Param, log *zap.Logger) Config {
	return &config{param: param, log: log}
}

func (c *config) InputFiles() ([]string, string, error) {
	encoding := c.param.Encoding()
	input := c.param.Input()
	depth := c.param.Depth()
	exts := strings.Split(strings.ToLower(c.param.Extension()), ",")

	input, err := filepath.Abs(input)
	if err != nil {
		return nil, "", ErrConvertAbsPath.Wrap(err)
	}
	input = filepath.Clean(input)
	stat, err := os.Stat(input)
	if err != nil {
		return nil, "", err
	}
	if !stat.IsDir() {
		return []string{input}, encoding, nil
	}

	rootDepth := strings.Count(input, string(os.PathSeparator))
	paths := []string{}
	err = filepath.WalkDir(input, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.Count(path, string(os.PathSeparator)) > rootDepth+depth {
				return fs.SkipDir
			}
		} else {
			ext := strings.ToLower(filepath.Ext(path))
			if len(ext) > 0 && slices.Contains(exts, ext[1:]) {
				paths = append(paths, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, "", ErrWalkInputDir.Wrap(err)
	}

	if len(paths) <= 0 {
		return nil, "", ErrNotFoundInputFile
	}

	return paths, encoding, nil
}

func (c *config) CellOption() *excel.CellOption {
	return &excel.CellOption{
		DecimalPlaces: c.param.DecimalPlaces(),
	}
}

func (c *config) XlsxFilename() string {
	return c.param.XlsxFilename()
}
