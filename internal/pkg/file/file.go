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
package file

import (
	"io"
	"io/fs"
	"os"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

type File struct {
}

func (r *File) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *File) Encoding(name string) (encoding.Encoding, error) {
	e, err := ianaindex.IANA.Encoding(name)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, ErrNotFoundEncoding.Details("encoding", name)
	}
	return e, nil
}

func (r *File) NewReader(f io.Reader, enc encoding.Encoding) io.ReadCloser {
	return io.NopCloser(transform.NewReader(f, enc.NewDecoder()))
}

func (r *File) NewWriter(f io.Writer, enc encoding.Encoding) io.WriteCloser {
	return transform.NewWriter(f, enc.NewEncoder())
}

func (r *File) Close(f io.Closer) error {
	return f.Close()
}

func (r *File) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}
