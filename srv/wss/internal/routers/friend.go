package router

import (
	"net/http"
	"strconv"

	"github.com/avp365/hl-sn/srv/wss/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

func FriendSet(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	friendId, err := strconv.Atoi(c.Param("friendid"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	id, err := user.FriendSet(userId, friendId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, id)

}
func FriendDelete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	friendId, err := strconv.Atoi(c.Param("friendid"))

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	id, err := user.FriendDelete(userId, friendId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	c.JSON(http.StatusOK, id)

}
