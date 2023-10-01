// Copyright (c) 2021 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package internal

import (
	"time"

	"github.com/cristalhq/jwt/v3"

	"go.uber.org/cadence/internal/common/auth"
	"go.uber.org/cadence/internal/common/util"
)

type JWTAuthProvider struct {
	PrivateKey []byte
}

func NewAdminJwtAuthorizationProvider(privateKey []byte) auth.AuthorizationProvider {
	return &JWTAuthProvider{
		PrivateKey: privateKey,
	}
}

func (j *JWTAuthProvider) GetAuthToken() ([]byte, error) {
	claims := auth.JWTClaims{
		Admin: true,
		Iat:   time.Now().Unix(),
		TTL:   60 * 10,
	}
	key, err := util.LoadRSAPrivateKey(j.PrivateKey)
	if err != nil {
		return nil, err
	}
	signer, err := jwt.NewSignerRS(jwt.RS256, key)
	if err != nil {
		return nil, err
	}
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	if token == nil {
		return nil, err
	}

	return token.Raw(), nil
}
