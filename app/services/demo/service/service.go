// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import "github.com/yuw-pot/pot/src/services"

type service struct {
	parent *services.Services
}

func New() *service {
	return &service {
		parent: services.New(),
	}
}
