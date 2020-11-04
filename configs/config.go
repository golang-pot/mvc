// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package configs

import (
	"github.com/yuw-pot/pot/data"
)

const (
	CTR = "CTR"
	SRV = "SRV"
	MOD = "MOD"
	TPL = "TPL"

	SuccessOK	= "SuccessOK"
)

var (
	// define self Error Message
	//   - CTR: controller
	//   - SRV: service
	//   - MOD: model
	ErrMsg *data.ErrH = &data.ErrH{
		TPL: {
			SuccessOK: "Success",
		},
		CTR: {

		},
		SRV: {

		},
		MOD: {

		},
	}
)
