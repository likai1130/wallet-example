// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import (
	"wallet-example/internal/pkg/errors"

	"github.com/novalagung/gubrak"
)

// ErrCode implements `github.com/marmotedu/errors`.Coder interface.
type ErrCode struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// client message en zn
	Messages map[string]string
}

func (e ErrCode) HTTPStatus() int {
	return e.HTTP
}

func (e ErrCode) Code() int {
	return e.C
}

func (e ErrCode) Message(lang string) string {
	return e.Messages[lang]
}

var _ errors.CoderExtend = &ErrCode{}

// nolint: unparam
func register(code int, httpStatus int, enMeg, cnMsg string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 409, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 409,500`")
	}

	message := map[string]string{
		"__default__": enMeg,
		"zh-cn":       cnMsg,
		"en":          cnMsg,
	}
	coder := &ErrCode{
		C:        code,
		HTTP:     httpStatus,
		Messages: message,
	}

	errors.MustRegisterExtend(coder)
}
