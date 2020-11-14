// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo_demo

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/adapter"
	"github.com/yuw-pot/pot/src/models"
)

type DemoModel struct {
	mod *models.Models
}

func NewDemoModel() *DemoModel {
	conn, err := adapter.Conns(adapter.Mysql, "I")
	if err != nil { panic(err) }

	return &DemoModel {
		mod: models.New(conn),
	}
}

func (m *DemoModel) FetchOne(mPoT *data.ModPoT) (*Demo, error) {
	mPoT.Types = data.ModONE
	mPoT.Columns = []string{"id", "name", "create_time"}

	d := &Demo{}
	_, err := m.mod.GeT(mPoT, d)

	return d, err
}

func (m *DemoModel) FetchOneById(mPoT *data.ModPoT) (*Demo, error) {
	mPoT.Types = data.ModONE

	d := &Demo{}
	_, err := m.mod.GeT(mPoT, d)

	return d, err
}

func (m *DemoModel) FetchOneJoinById(mPoT *data.ModPoT) ([]DM, error) {
	mPoT.Types = data.ModALL
	mPoT.Joins = []*data.ModJoinPoT {
		{
			JoinOperator: data.ModJoinRIGHT,
			TableName:    "demo_to_many",
			Condition:    "demo.id=demo_to_many.demo_id",
		},
	}

	mPoT.Table = "demo"
	mPoT.Field = []string{"demo.*","demo_to_many.*"}

	d := make([]DM, 0)
	_, err := m.mod.GeT(mPoT, &d)

	return d, err
}

func (m *DemoModel) FetchAll(mPoT *data.ModPoT) ([]Demo, error) {
	mPoT.Types = data.ModALL

	d := make([]Demo, 0)
	_, err := m.mod.GeT(mPoT, &d)

	return d, err
}

func (m *DemoModel) Total() (int64, error) {
	return m.mod.Total(&Demo{Id:1})
}
