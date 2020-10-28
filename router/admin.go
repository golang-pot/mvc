// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import "github.com/yuw-pot/pot/routes"

func (route *AdminPoT) Tag() string {
	return TagAdminPoT
}

func (route *AdminPoT) PuT(r *routes.PoT, toFunc map[*routes.KeY][]interface{}) {
	routes.To(r.Eng.Group("/admin"), toFunc)
}
