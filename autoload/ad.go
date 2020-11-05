// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	PoT "github.com/yuw-pot/pot"
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
	start := PoT.New()
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
		rPoT.Tag(): {
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "Sample",
				Mode: R.PoTMethodGeT,
				Path: "/sample",
			}:{
				mPoT.Cors(),
				ctrl.Sample,
			},
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "SampleComponents",
				Mode: R.PoTMethodGeT,
				Path: "/sample_components",
			}:{
				ctrl.SampleComponents,
			},
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "SeTSampleCacheComponent",
				Mode: R.PoTMethodGeT,
				Path: "/set_sample_cache_components",
			}:{
				ctrl.SeTSampleCacheComponent,
			},
		},
		rAdminPoT.Tag(): {
			&R.KeY{
				Service: "Demo",
				Controller: "Demo",
				Action: "AdminDemo",
				Mode: R.PoTMethodGeT,
				Path: "/admin_demo",
			}:{
				ctrlAdmin.AdminDemo,
			},
		},
	}

	ad.ePoT.ErrMsg = configs.ErrMsg
	return ad
}