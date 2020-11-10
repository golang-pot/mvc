// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"mvc/app/services/demo/service"
	"mvc/configs"
)

type PublishController struct {
	srvCaches *service.CacheService
	v *utils.PoT
}

func NewPublishController() *PublishController {
	return &PublishController{
		srvCaches: service.NewCacheService(),
		v: utils.New(),
	}
}

func (c *PublishController) Publish(ctx *gin.Context) {
	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()

	TpL.Response = c.v.MergeH(
		TpL.Response,
		c.srvCaches.Publish("services", "test pool success !!"),
	)

	ctx.JSON(data.PoTStatusOK, TpL)
	return
}
