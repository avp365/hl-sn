package router

import (
	"net/http"

	"github.com/avp365/hl-sn/srv/sn/internal/entities"
	"github.com/avp365/hl-sn/srv/sn/internal/handlers/login"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginForm entities.LoginForm

	if err := c.ShouldBind(&loginForm); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	token, err := login.LoginHandler(&loginForm)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
