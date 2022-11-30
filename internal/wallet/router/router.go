package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"wallet-example/internal/wallet/controller/v1/blockchain"
	"wallet-example/internal/wallet/store/mongo"
	mysql "wallet-example/internal/wallet/store/msyql"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// initRouter 路由示例
func initRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r, authMiddleware)

	return r
}

// noCheckRoleRouter 无需认证的路由示例
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
	mysqlStoreIns, _ := mysql.GetMySQLFactoryOr()
	mongoStoreIns, _ := mongo.GetMongoFactoryOr()

	{
		// chains RESTful resource
		chainv1 := v1.Group("/chains")
		{
			chainController := blockchain.NewBlockchainController(mysqlStoreIns, mongoStoreIns)
			chainv1.POST("", chainController.Create)
			chainv1.DELETE(":id", chainController.Delete)
			chainv1.PUT(":id", chainController.Update)
			chainv1.GET(":id", chainController.Get)
			chainv1.GET("", chainController.List)
		}
	}
}

// checkRoleRouter 需要认证的路由示例
func checkRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/v1")

	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}
