// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import "github.com/yuw-pot/pot/libs"

type (
	M struct {
		lib *libs.PoT
	}
)

func New() *M {
	return &M {
		lib: new(libs.PoT),
	}
}

