// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/yuw-pot/pot/src/components/cache"
)

type (
	M struct {
		client *cache.RedisComponent
	}
)

func New() *M {
	// initialized cache client
	client := cache.NewRedis("I")

	return &M {
		client: client,
	}
}

