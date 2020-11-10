// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
)

type SubscriberService struct {

}

func NewSubscriberService() *SubscriberService {
	return &SubscriberService{}
}

func (srv *SubscriberService) Provided(channel string, content interface{}) {
	fmt.Println("test subscribe success !!")
}


