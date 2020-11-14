// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/src/components/cache"
)

type CacheService struct {
	s *service
	client *cache.RedisComponent
}

func NewCacheService() *CacheService {
	// add cache prefix
	cachePrefix := properties.PropertyPoT.GeT("PoT.Name", "")

	// initialized cache client
	client := cache.NewRedis("I")
	client.SeTPrefix(cast.ToString(cachePrefix))

	return &CacheService{
		s: New(),
		client: client,
	}
}

func (srv *CacheService) Publish(channel string, message interface{}) *data.H {
	_, _ = srv.client.Publish(channel, message)
	return &data.H {
		"Publish":"",
	}
}

func (srv *CacheService) SeTComponentCache() *data.H {
	content := &data.H {
		"SeTCache": "success set cache",
		"HSeTCache": "Success h set cache",
	}

	_, err := srv.client.SeT("PoT_cache_component", "test.success!!!")
	if err != nil {
		(*content)["SeTCache"] = err.Error()
	} else {
		(*content)["SeTCache"] = "success!!"
	}

	rHSeT := map[string]interface{}{
		"PoT_HSeT": "success h set cache!!",
	}

	_, err = srv.client.HSeT("PoT_H_cache_component", rHSeT)
	if err != nil {
		(*content)["HSeTCache"] = err.Error()
	} else {
		(*content)["HSeTCache"] = "success!!"
	}


	return content
}

func (srv *CacheService) GeTComponentCache() *data.H {
	okExisT, _ := srv.client.IsExisT("PoT_cache_component")
	okHExisT, _ := srv.client.IsHExisT("PoT_H_cache_component", "PoT_HSeT")

	contentGeT, _ := srv.client.GeT("PoT_cache_component")
	contentHGeT, _ := srv.client.HGeT("PoT_H_cache_component", "PoT_HSeT")

	contentKeYs, _ := srv.client.KeYs()

	return &data.H {
		"components.redis.client.IsExist": okExisT,
		"components.redis.client.IsHExist": okHExisT,
		"components.redis.client.GeT": contentGeT,
		"components.redis.client.HGeT": contentHGeT,
		"components.redis.client.KeYs": contentKeYs,
	}
}

