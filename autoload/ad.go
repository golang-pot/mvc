// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	T "github.com/yuw-pot/pot"
	E "github.com/yuw-pot/pot/modules/err"
	R "github.com/yuw-pot/pot/routes"
	"mvc/configs"
	"mvc/router"
)

var (
	_ R.Routes = new(router.PoT)
	_ R.Routes = new(router.AdminPoT)

	rPoT *R.PoT = &R.PoT{
		Src: nil,
		Arr: nil,
	}
)

type autoload struct {
	rPoT *R.PoT
	ePoT *E.PoT
}

func init() {
	// Initialized Autoload
	ad := ad().initialized()

	// Initialized PoT
	//   - add PoT Route
	//   - add PoT Error
	start := T.New()
	start.PoTRoute = ad.rPoT
	start.PoTError = ad.ePoT

	// Run PoT
	start.PoT().Run()
}

func ad() *autoload {
	return &autoload {
		rPoT: &R.PoT {
			Src: nil,
			Arr: nil,
		},
		ePoT: &E.PoT {
			ErrMsg: nil,
		},
	}
}

func (ad *autoload) initialized() *autoload {
	// Initialized Router
	rPoT := new(router.PoT)
	rAdminPoT := new(router.AdminPoT)

	//   - add Router Source
	//   - add Router Array
	ad.rPoT.Src = &R.RouteSrc { rPoT, rAdminPoT }
	ad.rPoT.Arr = &R.RouteArr {
		rAdminPoT.Tag(): {
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "X",
				Mode: R.PoTMethodGeT,
				Path: "/demo_x",
			}:{
				c.PoT(ctrl.X),
			},
		},
		rPoT.Tag(): {
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "T",
				Mode: R.PoTMethodGeT,
				Path: "/demo",
			}:{
				R.Cors(), c.PoT(ctrl.T),
			},
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "X",
				Mode: R.PoTMethodGeT,
				Path: "/demo_x",
			}:{
				R.Cors(),
				c.PoT(ctrl.X),
			},
		},
	}

	ad.ePoT.ErrMsg = configs.ErrMsg
	return ad
}