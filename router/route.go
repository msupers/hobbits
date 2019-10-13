package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/msupers/hobbits/controller"
	"github.com/msupers/hobbits/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//写入gin日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	router := gin.Default()
	router.Use(middlewares...)

	//demo
	v1 := router.Group("/demo")
	//v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	v1.Use(middleware.RequestLog())
	{
		controller.DemoRegister(v1)
	}

	//jenkins
	jks := router.Group("/jenkins")
	jks.Use(middleware.RequestLog())
	{
		controller.JenkinsRegister(jks)
	}

	//api
	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiController := &controller.Api{}
	apiNormalGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	apiNormalGroup.POST("/login", apiController.Login)
	apiNormalGroup.GET("/loginout", apiController.LoginOut)

	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	apiAuthGroup.GET("/user/listpage", apiController.ListPage)
	apiAuthGroup.GET("/user/add", apiController.AddUser)
	apiAuthGroup.GET("/user/edit", apiController.EditUser)
	apiAuthGroup.GET("/user/remove", apiController.RemoveUser)
	apiAuthGroup.GET("/user/batchremove", apiController.RemoveUser)
	return router
}
