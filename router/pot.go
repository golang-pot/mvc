// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/routes"
)

func (route *PoT) Tag() string {
	return TagPoT
}

func (route *PoT) PuT(r *gin.Engine, toFunc map[*routes.KeY][]gin.HandlerFunc) {
	routes.To(r.Group(""), toFunc)
}

