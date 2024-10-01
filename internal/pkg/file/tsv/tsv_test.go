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
package tsv

import (
	"io"
	"strconv"
	"testing"

	mock_txt "github.com/kenita8/xlcmd/internal/pkg/file/txt/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestReadOneLine(t *testing.T) {
	testcases := []struct {
		str        string
		expectData []string
		expectErr  error
	}{
		{"abc\tdef\tghi\njkl\tmno\tpqr", []string{"abc", "def", "ghi"}, nil},
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
			tx := NewTsvFile(file, encoding)
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
