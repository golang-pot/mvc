// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package configs

import (
	"github.com/yuw-pot/pot/data"
)

const (
	CTR string = "CTR"
	SRV string = "SRV"
	MOD string = "MOD"
)

var (
	// define self Error Message
	//   - CTR: controller
	//   - SRV: service
	//   - MOD: model
	ErrMsg *data.ErrH = &data.ErrH{
		CTR: {
			"UnKnown": "未知错误",
			"NoPriority": "无权限",
		},
		SRV: {

		},
		MOD: {

		},
	}
)
