// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/libs"
	C "github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	"mvc/configs"
)

func (m *M) JwTAuth() *libs.PoT {
	return m.lib.SeT(func(lib *libs.PoT) {
		ctx := lib.Lib()

		token := ctx.Request.Header.Get("token")
		//fmt.Println(token)
		//fmt.Println(ctx.Request.Header)
		if token == "" {
			ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
				Status:   data.PoTUnKnown,
				Msg:      E.Err(configs.CTR, "NoPriority").Error(),
				Response: nil,
			})
			ctx.Abort()
			return
		}

		JwT, err := C.JPoT.Parse(token)
		if err != nil {
			if err == E.Err(data.ErrPfx, "TokenExpired") {
				ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
					Status:   data.PoTUnKnown,
					Msg:      E.Err(configs.CTR, "NoPriority").Error(),
					Response: nil,
				})
				ctx.Abort()
				return
			} else {
				ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
					Status:   data.PoTUnKnown,
					Msg:      err.Error(),
					Response: nil,
				})
				ctx.Abort()
				return
			}
		}

		JwTRefresh, _ := C.JPoT.Refresh(token)

		ctx.Set("JwT", JwTRefresh)
		ctx.Set("JwTInfo", JwT)
		ctx.Next()
	})
}
