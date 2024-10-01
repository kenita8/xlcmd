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
package txt

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kenita8/xlcmd/internal/pkg/file"

	"golang.org/x/text/encoding"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE
//go:generate mockgen -destination=./mock/mock_io.go -package=mock_txt io ReadWriteCloser

type Filer interface {
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Encoding(name string) (encoding.Encoding, error)
	NewReader(f io.Reader, enc encoding.Encoding) io.ReadCloser
	NewWriter(f io.Writer, enc encoding.Encoding) io.WriteCloser
	Close(f io.Closer) error
}

type TxtFiler interface {
	Basename() string
	Filename() string
	OpenReadMode() error
	OpenReadModeInternal() error
	ReadOneLine() ([]string, error)
	OpenWriteMode() error
	OpenWriteModeInternal() error
	WriteOneLine([]string) error
	Close()
}

type TxtFile struct {
	Filer    Filer
	TxtFiler TxtFiler
	Pathname string
	EncName  string
	Fp       io.ReadWriteCloser
	Rc       io.ReadCloser
	Wc       io.WriteCloser
	scanner  *bufio.Scanner
}

func NewTxtFile(pathname string, encoding string) *TxtFile {
	txt := &TxtFile{
		Pathname: pathname,
		EncName:  encoding,
		Filer:    &file.File{},
	}
	txt.TxtFiler = txt
	return txt
}

func (r *TxtFile) Extension() string {
	return strings.ToLower(filepath.Ext(r.Pathname))
}

func (r *TxtFile) Basename() string {
	return filepath.Base(r.Pathname)
}

func (r *TxtFile) Filename() string {
	return r.Pathname
}

func (r *TxtFile) OpenReadMode() error {
	fp, err := r.Filer.OpenFile(r.Pathname, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	enc, err := r.Filer.Encoding(r.EncName)
	if err != nil {
		return err
	}
	r.Fp = fp
	r.Rc = io.NopCloser(r.Filer.NewReader(fp, enc))
	return r.TxtFiler.OpenReadModeInternal()
}

func (r *TxtFile) OpenReadModeInternal() error {
	r.scanner = bufio.NewScanner(r.Fp)
	return nil
}

func (r *TxtFile) ReadOneLine() ([]string, error) {
	scanned := r.scanner.Scan()
	if scanned {
		return []string{r.scanner.Text()}, nil
	} else {
		return nil, io.EOF
	}
}

func (r *TxtFile) OpenWriteMode() error {
	//Not implemented
	return nil
}

func (r *TxtFile) OpenWriteModeInternal() error {
	//Not implemented
	return nil
}

func (r *TxtFile) WriteOneLine([]string) error {
	//Not implemented
	return nil
}

func (r *TxtFile) Close() {
	r.Fp.Close()
}
