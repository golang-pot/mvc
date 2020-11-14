// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/casbin"
	"github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"github.com/yuw-pot/pot/src/controllers"
	"mvc/app/services/demo/service"
	"mvc/configs"
)

type DemoController struct {
	parent 	*controllers.Controller
	v 		*utils.PoT
	cry 	*crypto.PoT

	srv 		*service.DemoService
	srvAuth 	*service.AuthService
	srvCaches 	*service.CacheService
	srvCrypto 	*service.CryptoService
}

func NewDemoController() *DemoController {
	return &DemoController {
		parent: controllers.New(),
		v: 		utils.New(),
		cry: 	crypto.New(),

		srv: 		service.NewDemoService(),
		srvAuth: 	service.NewAuthService(),
		srvCaches: 	service.NewCacheService(),
		srvCrypto: 	service.NewCryptoService(),
	}
}

func (c *DemoController) GeTTesTModel(ctx *gin.Context) {
	//conn, err := adapter.Start(adapter.Mysql, "I")
	//if err != nil { panic(err) }
	//
	//fmt.Println(conn.DataSourceName())
	//
	//mod := models.New(conn)
	//mPoT := &data.ModPoT{
	//	Types:     "",
	//	Table:     "",
	//	Field:     nil,
	//	Joins:     nil,
	//	Limit:     0,
	//	Start:     nil,
	//	Query:     nil,
	//	QueryArgs: nil,
	//	Columns:   nil,
	//	OrderType: "",
	//	OrderArgs: nil,
	//}
	//
	//mPoT.Types = data.ModONE
	//mPoT.Columns = []string{"id", "name", "create_time"}
	//
	//d := &repo_demo.Demo{}
	//_, err = mod.GeT(mPoT, d)
	//
	//if err != nil { panic(err) }
	//fmt.Println(d)

	ctx.Abort()
	return
}

func (c *DemoController) SeTSampleCacheComponent(ctx *gin.Context) {
	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()

	TpL.Response = c.v.MergeH(
		TpL.Response,
		c.srvCaches.SeTComponentCache(),
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

	// GeT This Service, Controller, Action
	service, controller, action := c.v.GeTUintPtrFuncPC()

	casbinEnforcer := casbin.New()
	casbinEnforcer.AdapterInfo = &casbin.AdapterRedis{Tag:"I"}
	fmt.Println(casbinEnforcer.Enforcer())
	fmt.Println(casbinEnforcer.Add("admin", "d1", "read"))

	fmt.Println(casbinEnforcer.Check("admin","d1","read"))
	fmt.Println(casbinEnforcer.Check("admin","d1","w"))

	TpL.Response = c.v.MergeH(
		TpL.Response,
		c.srvCrypto.SampleCrypto(),

		// components.cache.redis.client
		//c.srvCaches.GeTComponentCache(),
		// JwT
		//c.srvAuth.SampleJwT(),

		// GeT UintPtr Service, Controller, Action
		&data.H {
			"uintptr_func_pc_service": service,
			"uintptr_func_pc_controller": controller,
			"uintptr_func_pc_action": action,
		},
	)

	ctx.JSON(data.PoTStatusOK, TpL)
	return
}

func (c *DemoController) SampleJwTParse(ctx *gin.Context) {
	info, ok := ctx.Get("Info")
	if ok == false {
		info = &configs.User{}
	}

	accessToken, ok := ctx.Get("NewAccessToken")
	if ok == false {
		accessToken = ""
	}

	TpL := data.TpLInitialized()
	TpL.Msg = E.Err(configs.TPL, configs.SuccessOK).Error()
	TpL.Response = &data.H {
		"Login": "accessToken Login",
		"Info": info,
		"AccessToken": accessToken,
	}

	ctx.JSON(data.PoTStatusOK, TpL)
	return
}