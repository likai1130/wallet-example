// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	log "github.com/kingwel-xie/k2/core/logger"
	"net/http"
	"wallet-example/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

var Default = &response{}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	resp := Default.Clone()
	msgId := generateMsgIDFromContext(c)
	lang := getAcceptLanguage(c)
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoderExtend(err)
		c.JSON(coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.Message(lang),
			RequestId: msgId,
		})
		return
	}

	resp.SetTraceID(msgId)
	resp.SetData(data)
	resp.SetCode(0)
	if lang == DefaultLanguage {
		resp.SetMsg("Ok")
	} else {
		resp.SetMsg("成功")
	}
	c.JSON(http.StatusOK, resp)
}
