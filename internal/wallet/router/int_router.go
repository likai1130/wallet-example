package router

import (
	"github.com/gin-gonic/gin"
	sdk "github.com/kingwel-xie/k2/common"
	common "github.com/kingwel-xie/k2/common/middleware"
	log "github.com/kingwel-xie/k2/core/logger"
	"os"
)

func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册业务路由
	initRouter(r, authMiddleware)
}
