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
package excel

import "github.com/kenita8/errors"

var (
	ErrNotOpened       = errors.New("XLSX file has not been opened yet")
	ErrOpenXlsxFile    = errors.New("unable to open XLSX file")
	ErrSaveAsFile      = errors.New("unable to save output file")
	ErrReadInputFile   = errors.New("unable to read from input file")
	ErrNewSheet        = errors.New("failed to create new sheet")
	ErrConvertCellName = errors.New("unable to convert cell name")
	ErrGetCellValue    = errors.New("unable to get cell value")
	ErrAddChart        = errors.New("failed to add chart")
	ErrSetCellValue    = errors.New("unable to write data to cell")
)
