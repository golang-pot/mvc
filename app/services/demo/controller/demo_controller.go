// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/modules/utils"
	"github.com/yuw-pot/pot/src/controllers"
	"mvc/app/services/demo/service"
	"mvc/configs"
)

type DemoController struct {
	parent *controllers.Controller
	srv *service.DemoService
	v *utils.PoT
	cry *crypto.PoT
}

func NewDemoController() *DemoController {
	return &DemoController {
		parent: controllers.New(),
		srv: service.NewDemoService(),
		v: utils.New(),
		cry: crypto.New(),
	}
}

func (c *DemoController) SeTSampleCacheComponent(ctx *gin.Context) {
	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()

	TpL.Response = c.v.MergeH(
		TpL.Response,
		c.srv.SeTComponentCache(),
	)

	ctx.JSON(data.PoTStatusOK, TpL)
	return
}

func (c *DemoController) Sample(ctx *gin.Context) {
	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()

	// get fetch one
	do := c.srv.FetchOne()
	doById := c.srv.FetchOneById(2, "id", "create_time")
	doJoinById := c.srv.FetchOneJoinById(1)

	// get fetch all
	da := c.srv.FetchAll()

	TpL.Response = c.v.MergeH(
		TpL.Response,
		do.Data,
		doById.Data,
		doJoinById.Data,
		da.Data,

		c.srv.Total(),
		// - get this url path
		c.srv.GeTPath(),
	)

	ctx.JSON(TpL.Status, TpL)
	return
}

func (c *DemoController) SampleComponents(ctx *gin.Context) {
	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()

	c.cry.Mode = data.ModeToken

	// - md5
	c.cry.D = []interface{}{data.MD5, "password_md5"}
	md5, _ := c.cry.Made()

	// - sha1
	c.cry.D = []interface{}{data.SHA1, "password_sha1"}
	sha1, _ := c.cry.Made()

	// - sha256
	c.cry.D = []interface{}{data.SHA1, "password_sha256"}
	sha256, _ := c.cry.Made()

	// - aes
	c.cry.Mode = data.ModeAeS

	aesKeY := properties.PropertyPoT.GeT("AeS.KeY", nil)
	c.cry.D = []interface{}{aesKeY}

	aes, _ := c.cry.Made()
	aesEncrypt, err := aes.(*crypto.AeSPoT).EncrypT("test success")
	if err != nil {
		aesEncrypt = err.Error()
	}

	aesDecrypt, err := aes.(*crypto.AeSPoT).DecrypT(aesEncrypt)
	if err != nil {
		aesDecrypt = err.Error()
	}

	// components.crypto
	sampleCrypto := &data.H {
		"components_crypto": map[string]interface{}{
			"md5": md5,
			"sha1": sha1,
			"sha256": sha256,

			"aes": map[string]interface{}{
				"Encrypt": aesEncrypt,
				"Decrypt": aesDecrypt,
			},
		},
	}

	// components.cache.redis.client
	sampleCaches := c.srv.GeTComponentCache()

	TpL.Response = c.v.MergeH(
		TpL.Response,
		sampleCrypto,
		sampleCaches,
	)

	ctx.JSON(data.PoTStatusOK, TpL)
	return
}