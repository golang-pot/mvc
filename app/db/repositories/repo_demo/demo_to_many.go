// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo_demo

import (
	"time"
)

type DemoToMany struct {
	Id         int       `xorm:"not null pk autoincr INT(11)" json:"id,omitempty"`
	DemoId     int       `xorm:"INT(11)" json:"demo_id,omitempty"`
	Name       string    `xorm:"VARCHAR(255)" json:"name,omitempty"`
	CreateTime *time.Time `xorm:"TIMESTAMP" json:"create_time,omitempty"`
	UpdateTime *time.Time `xorm:"TIMESTAMP" json:"update_time,omitempty"`
}
