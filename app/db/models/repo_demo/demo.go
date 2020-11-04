// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo_demo

import (
	"mvc/app/db/repositories/repo_demo"
)

const (
	repo string = "repo_demo"
)

type (
	Demo repo_demo.Demo
	DemoToMany repo_demo.DemoToMany
	Dt repo_demo.Dt

	DM struct {
		D Demo `xorm:"extends"`
		M DemoToMany `xorm:"extends"`
	}

	DMPoT struct {
		X interface{}
		Y interface{}
	}

	DMForMaTPoT struct {
		X interface{}
		Y []interface{}
	}
)


