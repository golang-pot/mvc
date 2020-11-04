// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo_demo

type DemoDt struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	DemoId int `xorm:"INT(11)"`
	DtId   int `xorm:"INT(11)"`
}
