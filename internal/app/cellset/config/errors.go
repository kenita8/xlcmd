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

import "github.com/kenita8/errors"

var (
	ErrInvalidRange       = errors.New(`specified range is invalid. examples "A1", "A1:C30"`)
	ErrRequireReplacement = errors.New(`replacement parameter is required`)
	ErrRequirePattern     = errors.New(`pattern parameter is required`)
	ErrRegexpCompile      = errors.New(`failed to compile the regular expression`)
)
