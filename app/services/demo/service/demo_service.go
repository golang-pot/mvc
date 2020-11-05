// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/src/components/cache"
	"mvc/app/db/models/repo_demo"
)

type DemoService struct {
	mDemo *repo_demo.DemoModel
	s *service
	client *cache.RedisComponent
}

func NewDemoService() *DemoService {
	// add cache prefix
	cachePrefix := properties.PropertyPoT.GeT("PoT.Name", "")

	// initialized cache client
	client := cache.NewRedis("I")
	client.SeTPrefix(cast.ToString(cachePrefix))

	return &DemoService{
		mDemo: repo_demo.NewDemoModel(),
		s: New(),
		client: client,
	}
}

func (srv *DemoService) GeTPath() *data.H {
	content, _ := srv.s.parent.GeTPath("Demo", "Demo", "Sample")

	return &data.H {
		"func.path": content,
	}
}

func (srv *DemoService) SeTComponentCache() *data.H {
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

func (srv *DemoService) GeTComponentCache() *data.H {
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

func (srv *DemoService) FetchOne() *data.SrvTpL {
	TpL := data.SrvTpLInitialized()

	mPoT := &data.ModPoT{}

	d, err := srv.mDemo.FetchOne(mPoT)
	if err != nil {
		TpL.Status = data.PoTUnKnown
		TpL.Msg = err.Error()

		return TpL
	}

	TpL.Data = &data.H {
		"db.select.one": d,
	}

	return TpL
}

func (srv *DemoService) FetchOneById(id int, cols ... string) *data.SrvTpL {
	TpL := data.SrvTpLInitialized()

	mPoT := &data.ModPoT{}
	mPoT.Query = "demo.id=?"
	mPoT.QueryArgs = []interface{}{id}

	if len(cols) > 0 {
		mPoT.Columns = cols
	}

	d, err := srv.mDemo.FetchOneById(mPoT)
	if err != nil {
		TpL.Status = data.PoTUnKnown
		TpL.Msg = err.Error()

		return TpL
	}

	TpL.Data = &data.H {
		"db.select.by.id": d,
	}

	return TpL
}

func (srv *DemoService) FetchOneJoinById(id ... interface{}) *data.SrvTpL {
	TpL := data.SrvTpLInitialized()

	mPoT := &data.ModPoT{}

	mPoT.Query = "demo.id=?"
	mPoT.QueryArgs = id

	d, err := srv.mDemo.FetchOneJoinById(mPoT)
	if err != nil {
		TpL.Status = data.PoTUnKnown
		TpL.Msg = err.Error()

		return TpL
	}

	TpL.Data = &data.H {
		"db.select.join": d,
	}

	return TpL
}

func (srv *DemoService) FetchAll() *data.SrvTpL {
	TpL := data.SrvTpLInitialized()

	mPoT := &data.ModPoT{}
	d, err := srv.mDemo.FetchAll(mPoT)
	if err != nil {
		TpL.Status = data.PoTUnKnown
		TpL.Msg = err.Error()

		return TpL
	}

	TpL.Data = &data.H {
		"db.select.all": d,
	}

	return TpL
}

func (srv *DemoService) Total() *data.H {
	total, _ := srv.mDemo.Total()
	return &data.H {
		"db.table.total": total,
	}
}
