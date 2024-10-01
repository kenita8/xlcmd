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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	fl := File{}
	_, err := fl.OpenFile("invalid file path", os.O_RDONLY, 0)
	assert.EqualError(t, err, "open invalid file path: no such file or directory")

	f, err := fl.OpenFile("file_test.go", os.O_RDONLY, 0)
	assert.Nil(t, err)
	enc, err := fl.Encoding("UTF-8")
	assert.Nil(t, err)
	fl.NewReader(f, enc)
	fl.Close(f)
	assert.Nil(t, err)

	_, err = fl.Encoding("xxx")
	assert.EqualError(t, err, "ianaindex: invalid encoding name")

	_, err = fl.Stat("invalid file path")
	assert.EqualError(t, err, "stat invalid file path: no such file or directory")

	_, err = fl.Stat("file_test.go")
	assert.Nil(t, err)
}
