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
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"testing"

	mock_txt "github.com/kenita8/xlcmd/internal/pkg/file/txt/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBasename(t *testing.T) {
	testcases := []struct {
		pathname        string
		expectBasename  string
		expectExtension string
	}{
		{
			pathname:        "/aaa/bbb/ccc.txt",
			expectBasename:  "ccc.txt",
			expectExtension: ".txt",
		},
		{
			pathname:        "ccc.csv",
			expectBasename:  "ccc.csv",
			expectExtension: ".csv",
		},
		{
			pathname:        `ccc`,
			expectBasename:  "ccc",
			expectExtension: "",
		},
		{
			pathname:        "Ccc.Txt",
			expectBasename:  "Ccc.Txt",
			expectExtension: ".txt",
		},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tf := NewTxtFile(tc.pathname, "UTF-8")
			bn := tf.Basename()
			ex := tf.Extension()
			fn := tf.Filename()
			assert.Equal(t, tc.expectBasename, bn)
			assert.Equal(t, tc.expectExtension, ex)
			assert.Equal(t, tc.pathname, fn)
		})
	}
}

func TestOpenReadMode(t *testing.T) {
	err := fmt.Errorf("ErrorOccurred")

	testcases := []struct {
		OpenErr     error
		EncodingErr error
		expectErr   error
	}{
		{err, nil, err},
		{nil, err, err},
		{nil, nil, nil},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			file := "file1"
			encoding := "UTF-8"
			var mode fs.FileMode = 0
			fp := &os.File{}
			mCtrl := gomock.NewController(t)
			defer mCtrl.Finish()
			mOsFiler := mock_txt.NewMockFiler(mCtrl)
			mOsFiler.EXPECT().OpenFile(file, os.O_RDONLY, mode).Return(fp, tc.OpenErr)
			if tc.OpenErr == nil {
				mOsFiler.EXPECT().Encoding(encoding).Return(nil, tc.EncodingErr)
			}
			if tc.OpenErr == nil && tc.EncodingErr == nil {
				mOsFiler.EXPECT().NewReader(fp, nil).Return(nil)
			}
			tx := NewTxtFile(file, encoding)
			tx.Filer = mOsFiler
			actualErr := tx.OpenReadMode()
			if tc.expectErr != nil {
				assert.EqualError(t, actualErr, tc.expectErr.Error())
			} else {
				assert.Nil(t, actualErr)
			}
		})
	}
}

func TestReadOneLine(t *testing.T) {
	testcases := []struct {
		str        string
		expectData []string
		expectErr  error
	}{
		{"abcdefg\nhijklmn", []string{"abcdefg"}, nil},
		{"", []string{}, io.EOF},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			mCtrl := gomock.NewController(t)
			defer mCtrl.Finish()
			mReadWriteCloser := mock_txt.NewMockReadWriteCloser(mCtrl)
			mReadWriteCloser.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (n int, err error) {
				copy(p, []byte(tc.str))
				return len(tc.str), tc.expectErr
			}).AnyTimes()
			file := "file1"
			encoding := "UTF-8"
			tx := NewTxtFile(file, encoding)
			tx.Fp = mReadWriteCloser
			tx.OpenReadModeInternal()
			actualData, actualErr := tx.ReadOneLine()
			if tc.expectErr != nil {
				assert.EqualError(t, actualErr, tc.expectErr.Error())
			} else {
				assert.Equal(t, actualData, tc.expectData)
			}
		})
	}
}

func TestOpenWriter(t *testing.T) {

}

func TestWriteOneLine(t *testing.T) {

}

func TestClose(t *testing.T) {

}
