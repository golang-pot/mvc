// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/jwt"
	"github.com/yuw-pot/pot/modules/properties"
	"mvc/configs"
)

// UnRefresh Token
func (m *M) JwTCacheAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TpL := data.TpLInitialized()
		TpL.Status = data.PoTUnKnown

		// - GeT Token
		accessToken := ctx.Request.Header.Get("token")
		if accessToken == "" {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		// - GeT JwT
		inJwT := jwt.NewJCacheAuth(m.client.GeTClienT())
		inJwT.Method = data.TimeSecond
		inJwT.Expire = 60

		if ok := inJwT.IsAccessTokenExisT(accessToken); ok == false {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		toInfo := &configs.User{}
		inJwT.KeY = cast.ToString(properties.PropertyPoT.GeT("JwT.KeY", ""))
		_, err := inJwT.Parse(accessToken, &toInfo)
		if err != nil {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		ctx.Set("NewAccessToken", accessToken)
		ctx.Set("Info", toInfo)
		ctx.Next()
	}
}

// Refresh Token
func (m *M) JwTCacheAuthRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TpL := data.TpLInitialized()
		TpL.Status = data.PoTUnKnown

		// - GeT Token
		accessToken := ctx.Request.Header.Get("token")
		if accessToken == "" {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		// - GeT JwT
		inJwT := jwt.NewJCacheAuth(m.client.GeTClienT())
		inJwT.Method = data.TimeSecond
		inJwT.Expire = 60

		if ok := inJwT.IsAccessTokenExisT(accessToken); ok == false {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		toInfo := &configs.User{}
		inJwT.KeY = cast.ToString(properties.PropertyPoT.GeT("JwT.KeY", ""))
		newAccessToken, err := inJwT.Parse(accessToken, &toInfo, jwt.RefreshToken)
		if err != nil {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()
			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		ctx.Set("NewAccessToken", newAccessToken)
		ctx.Set("Info", toInfo)
		ctx.Next()
	}
}
