// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/jwt"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/src/components/cache"
	"mvc/configs"
)

type AuthService struct {
	JwT *jwt.JCachePoT
}

func NewAuthService() *AuthService {
	return &AuthService{
		JwT: jwt.NewJCacheAuth(cache.NewRedis("I").GeTClienT()),
	}
}

func (srv *AuthService) SampleJwT() *data.H {
	srv.JwT.KeY = cast.ToString(properties.PropertyPoT.GeT("JwT.KeY", ""))
	srv.JwT.Method = data.TimeSecond
	srv.JwT.Expire = 60

	srv.JwT.Info = &configs.User{
		UserId:   2,
		UserName: "Kyle XY Test 2",
	}

	accessToken, _ := srv.JwT.Produce()

	return &data.H {
		"access_token": accessToken,
	}
}
