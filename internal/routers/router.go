package router

import (
	"net/http"
	"strconv"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/handlers/login"
	"github.com/avp365/hl-sn/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {

	res := []string{"v:1"}
	c.JSON(200, res)

}

func Login(c *gin.Context) {
	var loginForm entities.LoginForm

	if err := c.ShouldBind(&loginForm); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	ok, err := login.LoginHandler(&loginForm)

	if ok == "ok" {
		c.JSON(http.StatusOK, SignedResponse{
			Token:   tokenStr,
			Message: "logged in",
		})

		return
	}
	c.String(http.StatusNotFound, "Пользователь не найден")
}

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

func Run() {

	router := gin.Default()

	router.GET("/", Version)
	router.GET("/user/get/:userid", UserGetById)

	router.POST("/user/register", UserRegister)
	router.POST("/login", Login)

	router.Run()
}
