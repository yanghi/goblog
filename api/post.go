package api

import (
	"fmt"
	"goblog/auth"
	gerr "goblog/error"
	"goblog/model"
	"goblog/rep"
	"goblog/service/post"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var service post.CreatePostService

	var u, _ = c.Get("user")

	fmt.Println("gggg", u)
	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)

	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, service.Run())
}

func DeletePostByAuthor(c *gin.Context) {
	var service post.DeletePostService

	var u, _ = c.Get("user")

	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)
	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, service.DeleteByAuthor())
}

func GetPost(c *gin.Context) {
	var service post.GetPostService

	sid, hasId := c.GetQuery("id")

	if !hasId {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}
	id, er := strconv.Atoi(sid)

	if er != nil {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	service.Id = id

	c.JSON(200, service.Get())
}

func GetPostList(c *gin.Context) {
	var service post.GetPostListService

	pq := model.PaginationQuery{}

	c.ShouldBindQuery(&pq)
	pg := model.Pagination{}
	service.PaginationParams = *pg.Query(pq)

	c.JSON(200, service.Get())
}

func ModifyPost(c *gin.Context) {
	var service post.ModifyPostService

	var u, _ = c.Get("user")

	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)
	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, service.Modify())
}