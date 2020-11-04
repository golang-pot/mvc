// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
)

func (m *M) MiddlewareDemo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("PoTMiddlewareDemo", &data.H{
			"MiddlewareDemo": "Demo PuT Success",
		})
		ctx.Next()
	}
}
