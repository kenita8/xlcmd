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
package csv

import (
	rawCsv "encoding/csv"

	"github.com/kenita8/xlcmd/internal/pkg/file"
	"github.com/kenita8/xlcmd/internal/pkg/file/txt"
)

var (
	NewReader = rawCsv.NewReader
)

type CsvFile struct {
	txt.TxtFile
	csvReader *rawCsv.Reader
	Comma     rune
}

func NewCsvFile(pathname string, encoding string) *CsvFile {
	csv := &CsvFile{
		TxtFile: txt.TxtFile{
			Pathname: pathname,
			EncName:  encoding,
			Filer:    &file.File{},
		},
		Comma: ',',
	}
	csv.TxtFile.TxtFiler = csv
	return csv
}

func (c *CsvFile) OpenReadModeInternal() error {
	c.csvReader = NewReader(c.Fp)
	c.csvReader.Comma = c.Comma
	return nil
}

func (c *CsvFile) ReadOneLine() ([]string, error) {
	return c.csvReader.Read()
}
