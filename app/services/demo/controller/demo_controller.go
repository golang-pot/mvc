// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/libs"
	"github.com/yuw-pot/pot/modules/crypto"
	"mvc/app/services/demo/service"
)

type DemoController struct {
	srv *service.DemoService
	cry *crypto.PoT
}

func NewDemoController() *DemoController {
	return &DemoController {
		srv: service.NewDemoService(),
		cry: crypto.New(),
	}
}

func (c *DemoController) X(lib *libs.PoT) {
	ctx := lib.Lib()
	//fmt.Println(ctx.GeT().Request.Header)
	//fmt.Println(ctx.GeT().DefaultQuery("t",""))
	ctx.JSON(data.PoTStatusOK, data.SrvPoT{
		Status:   data.PoTStatusOK,
		Msg:      "",
		Response: &data.H {
			"x":"success x",
		},
	})

	ctx.Abort()

	return
}

func (c *DemoController) T(lib *libs.PoT) {
	ctx := lib.Lib()

	c.cry.Mode = data.ModeToken
	c.cry.D = []interface{}{data.MD5, "token"}

	token, _ := c.cry.Made()
	fmt.Println(cast.ToString(token))

	c.cry.Mode = data.ModeRsA
	c.cry.D = []interface{}{"rsa"}

	rsa, _ := c.cry.Made()
	fmt.Println((rsa.(*crypto.RsAPoT)).T())

	//token, _ := c.jwt.KeY
	//fmt.Println(string(c.jwt.KeY))
	//
	//info := &C.JwTPoT{
	//	Info: map[string]interface{}{
	//		"test":"test",
	//	},
	//}

	//data, _ := C.JPoT.Made(info)
	//fmt.Println(data)

	//fmt.Println(C.JPoT.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJbmZvIjp7InRlc3QiOiJ0ZXN0In0sImV4cCI6MTYwMzI3NzExNn0._eSt-E_EAetXjd5fg5Dpibro_tuoFlUe3gNlIeUZbMk"))
	//
	//fmt.Println(d)

	//fmt.Println(E.Err(data.ErrPfx, "PoTModeErr", "test"))
	//fmt.Println(E.Err(configs.CTR, "CtrErr", E.Position()))

	ctx.JSON(data.PoTStatusOK, data.SrvPoT {
		Status:   data.PoTStatusOK,
		Msg:      "test",
		Response: &data.H {
			"T":"success T",
		},
	})
	ctx.Abort()

	return
}