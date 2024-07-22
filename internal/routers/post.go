package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/handlers/post"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var postData entities.Post

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	if err := c.ShouldBind(&postData); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	postData.UserID = userId
	postData.Date = time.Now()

	post, err := post.CreatePost(postData)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, post)

}
func PostDelete(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	postId, err := strconv.Atoi(c.Param("postid"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	isDelete, err := post.Delete(postId, userId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, isDelete)

}

func PostUpdate(c *gin.Context) {
	var postData entities.Post

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	if err := c.ShouldBind(&postData); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	postData.UserID = userId
	postData.Date = time.Now()

	users, err := post.PostUpdate(postData)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, users)

}

func PostGet(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	postId, err := strconv.Atoi(c.Param("postid"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	post, err := post.Get(postId, userId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, post)

}

func PostFeed(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	posts, err := post.Feed(userId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, posts)

}
