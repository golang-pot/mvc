// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
	"github.com/yuw-pot/pot/modules/properties"
)

type CryptoService struct {
	s *service
	cry *crypto.PoT
}

func NewCryptoService() *CryptoService {
	return &CryptoService{
		s: New(),
		cry: crypto.New(),
	}
}

func (srv *CryptoService) SampleCrypto() *data.H {
	srv.cry.Mode = data.ModeToken

	// - md5
	srv.cry.D = []interface{}{data.MD5, "password_md5"}
	md5, _ := srv.cry.Made()

	// - sha1
	srv.cry.D = []interface{}{data.SHA1, "password_sha1"}
	sha1, _ := srv.cry.Made()

	// - sha256
	srv.cry.D = []interface{}{data.SHA256, "password_sha256"}
	sha256, _ := srv.cry.Made()

	// - aes
	srv.cry.Mode = data.ModeAeS

	aesKeY := properties.PropertyPoT.GeT("AeS.KeY", nil)
	srv.cry.D = []interface{}{aesKeY}

	aes, _ := srv.cry.Made()
	aesEncrypt, err := aes.(*crypto.AeSPoT).EncrypT("test success")
	if err != nil {
		aesEncrypt = err.Error()
	}

	aesDecrypt, err := aes.(*crypto.AeSPoT).DecrypT(aesEncrypt)
	if err != nil {
		aesDecrypt = err.Error()
	}

	return &data.H {
		"components_crypto": map[string]interface{}{
			"md5": md5,
			"sha1": sha1,
			"sha256": sha256,

			"aes": map[string]interface{}{
				"Encrypt": aesEncrypt,
				"Decrypt": aesDecrypt,
			},
		},
	}
}
