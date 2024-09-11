package router

import (
	"net/http"
	"strconv"

	"github.com/avp365/hl-sn/srv/wss/internal/entities"
	"github.com/avp365/hl-sn/srv/wss/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var registerForm entities.RegisterForm

	if err := c.ShouldBind(&registerForm); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	userId, err := user.RegisterUserHandler(&registerForm)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
	})

}

func UserGetById(c *gin.Context) {
	useridFromParam := c.Param("userid")

	userid, err := strconv.Atoi(useridFromParam)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	user, err := user.UserGetByIdHandler(userid)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"first_name":  user.FirstName,
		"second_name": user.SecondName,
		"birthdate":   user.Birthdate,
		"biography":   user.Biography,
		"city":        user.City,
	})

}
func UserSearch(c *gin.Context) {
	var searchForm entities.SearchForm

	if err := c.ShouldBind(&searchForm); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	users, err := user.UserSearchHandler(searchForm)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}
	c.JSON(http.StatusOK, users)

}
