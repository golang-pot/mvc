// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

type DemoService struct {

}

func NewDemoService() *DemoService {
	return &DemoService{

	}
}

func (srv *DemoService) GeT() string {
	return "test service success!!"
}
