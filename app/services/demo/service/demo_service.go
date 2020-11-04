// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/src/services/components/cache"
	"mvc/app/db/models/repo_demo"
)

type DemoService struct {
	mDemo *repo_demo.DemoModel
	s *service
	r *cache.RedisComponent
}

func NewDemoService() *DemoService {
	return &DemoService{
		mDemo: repo_demo.NewDemoModel(),
		s: New(),
		r: cache.NewRedis("I"),
	}
}

func (srv *DemoService) GeTPath() *data.H {
	content, _ := srv.s.parent.GeTPath("Demo", "Demo", "Sample")

	return &data.H {
		"func.path": content,
	}
}

func (srv *DemoService) GeTComponentCache() *data.H {
	content, _ := srv.r.GeT("test")

	return &data.H {
		"components.redis.client": content,
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
