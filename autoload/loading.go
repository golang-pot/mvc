// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/yuw-pot/pot/src/controllers"
	"mvc/app/services/demo/controller"
)

var (
	// Initialized Controllers
	c *controllers.Controller = controllers.New()
	ctrl *controller.DemoController = controller.NewDemoController()
)
