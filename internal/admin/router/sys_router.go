package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/middleware"
	"github.com/kingwel-xie/k2/core/ws"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"mime"
	_ "wallet-example/api/swagger/docs"
	"wallet-example/internal/admin/controller/sys"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	//sysStaticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}
	// 需要认证
	//sysCheckRoleRouterInit(g, authMiddleware)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	r.GET("/version", sys.Version)    //版本
	r.GET("/health", sys.HealthCheck) //健康检查

	r.GET("/metrics", middleware.GinWrapper(promhttp.Handler()))
}

// sysStaticFileRouter 系统文件路由
func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	r.Static("/site", "./site")
	r.Static("/static", "./static")
	if config.ApplicationConfig.Mode != "prod" {
		r.Static("/form-generator", "./static/form-generator")
	}
}

// sysSwaggerRouter admin swagger
func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// sysCheckRoleRouterInit 角色校验路由
func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	registerBaseRouter(v1, authMiddleware)
}

/**
registerBaseRouter 注册基本路由 菜单相关的
*/
func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	/*api := apis.SysMenu{}
	api2 := apis.SysDept{}*/
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		/*	v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
			v1auth.GET("/roleDeptTreeselect/:roleId", api2.GetDeptTreeRoleSelect)*/
		v1auth.POST("/logout", authMiddleware.LogoutHandler)
	}
}
