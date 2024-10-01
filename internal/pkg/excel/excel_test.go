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

import (
	"fmt"
	"strconv"
	"testing"

	mock_excelize "github.com/kenita8/xlcmd/internal/pkg/excel/excelize/mock"
	mock_excel "github.com/kenita8/xlcmd/internal/pkg/excel/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestOpen(t *testing.T) {
	err := fmt.Errorf("ErrorOccurred")
	testcases := []struct {
		StatErr     error
		OpenFileErr error
		expectErr   error
		expectNew   bool
	}{
		{
			StatErr:   err,
			expectErr: nil,
			expectNew: true,
		},
		{
			StatErr:     nil,
			OpenFileErr: err,
			expectErr:   ErrOpenXlsxFile.Wrap(err),
			expectNew:   false,
		},
		{
			StatErr:     nil,
			OpenFileErr: nil,
			expectErr:   nil,
			expectNew:   false,
		},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			file := "file1"
			mCtrl := gomock.NewController(t)
			defer mCtrl.Finish()
			mFiler := mock_excel.NewMockFiler(mCtrl)
			mExcelizer := mock_excelize.NewMockExcelizer(mCtrl)
			mFiler.EXPECT().Stat(file).Return(nil, tc.StatErr).AnyTimes()
			mExcelizer.EXPECT().NewFile().Return(nil).AnyTimes()
			mExcelizer.EXPECT().OpenFile(file).Return(nil, tc.OpenFileErr).AnyTimes()
			e := NewExcel(zap.NewNop())
			filer = mFiler
			excelizer = mExcelizer
			actualErr := e.Open(file)
			if tc.expectErr != nil {
				assert.EqualError(t, actualErr, tc.expectErr.Error())
			} else {
				assert.Nil(t, actualErr)
				assert.Equal(t, e.new, tc.expectNew)
			}
		})
	}
}

func TestNewSheet(t *testing.T) {
	err := fmt.Errorf("ErrorOccurred")
	testcases := []struct {
		NewSheetErr error
		expectErr   error
	}{
		{
			NewSheetErr: nil,
			expectErr:   nil,
		},
		{
			NewSheetErr: err,
			expectErr:   ErrNewSheet.Details("sheet", "sheet1").Wrap(err),
		},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			sheet := "sheet1"
			mCtrl := gomock.NewController(t)
			defer mCtrl.Finish()
			mXlFiler := mock_excelize.NewMockExcelizeFiler(mCtrl)
			mXlFiler.EXPECT().NewSheet(sheet).Return(1, tc.NewSheetErr).AnyTimes()
			e := NewExcel(zap.NewNop())
			e.xlFile = mXlFiler
			actualErr := e.NewSheet(sheet)
			if tc.expectErr != nil {
				assert.EqualError(t, actualErr, tc.expectErr.Error())
			} else {
				assert.Nil(t, actualErr)
			}
		})
	}
}

func TestPasteTxtFile(t *testing.T) {

}
