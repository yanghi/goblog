package main

import (
	"goblog/api"
	mysql "goblog/database/mysql"
	"goblog/router/middlewares"

	// serv "goblog/service"
	// "net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
}

func main() {
	mysql.New()
	r := gin.Default()
	// r := gin.New()

	var usePprof = false
	if usePprof {
		pprof.Register(r)
	}

	r.Use(middlewares.Cors())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/register", api.Register)
		v1.POST("/user/login", api.Login)
		v1.GET("/post/post", middlewares.TryGetUserByToken(), api.GetPost)
		v1.GET("/post/list", api.GetPostList)
		v1.PUT("/post/action/view", api.PostActionView)

		needAuth := v1.Group("/", middlewares.GetUserByToken(), middlewares.AuthRequired())

		needAuth.GET("user/post/list", api.GetUserPostList)

		needAuth.GET("user/info/base", api.GetUserBaseInfo)
		needAuth.PUT("user/info/info", api.ModifyUserInfo)

		needAuth.GET("tag/list", api.TagGetList)
		needAuth.POST("tag/tag", api.TagCreate)
		needAuth.PUT("tag/tag", api.TagModify)
		needAuth.DELETE("tag/tag", api.TagDelete)

		needAuth.POST("post/post", api.CreatePost)
		needAuth.DELETE("post/post", api.DeletePostByAuthor)
		needAuth.PUT("post/post", api.ModifyPost)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
