// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import "github.com/yuw-pot/pot/modules/casbin"

type CasbinService struct {
	casbinEnforcer *casbin.PoT
}

func NewCasbinService() *CasbinService {
	casbinEnforcer := casbin.New()
	casbinEnforcer.Adapter = casbin.RBACAdapterRedis
	casbinEnforcer.AdapterInfo = &casbin.AdapterRedis{Tag:"I"}

	return &CasbinService{
		casbinEnforcer: casbinEnforcer,
	}
}

func (srv *CasbinService) AddPolicy()  {

}




