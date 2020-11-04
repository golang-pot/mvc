// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
)

type AdminDemoController struct {

}

func NewAdminDemoController() *AdminDemoController {
	return &AdminDemoController{

	}
}

func (c *AdminDemoController) AdminDemo(ctx *gin.Context) {
	ctx.JSON(data.PoTStatusOK, data.TpLInitialized())
	return
}
