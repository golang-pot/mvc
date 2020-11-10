// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	middlewarePoT "github.com/yuw-pot/pot/src/middleware"
	middlewareMvc "mvc/app/middleware"
	II "mvc/app/services/admin/controller"
	I "mvc/app/services/demo/controller"
)

var (
	// Initialized Controllers
	ctrl *I.DemoController = I.NewDemoController()
	ctrlPublish *I.PublishController = I.NewPublishController()
	ctrlAdmin *II.AdminDemoController = II.NewAdminDemoController()

	// PoT. Middleware Struct
	mPoT *middlewarePoT.M = middlewarePoT.New()
	//Self Define Middleware Struct
	mMvc *middlewareMvc.M = middlewareMvc.New()
)
