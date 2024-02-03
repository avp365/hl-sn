package router

import (
	"net/http"
	"strconv"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/handlers/login"
	"github.com/avp365/hl-sn/internal/handlers/user"
	"github.com/avp365/hl-sn/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	token, err := login.LoginHandler(&loginForm)

	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

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
func jwtMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		jwtToken, err := token.ExtractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Message": err.Error(),
			})
			return
		}

		token, err := token.ParseToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Message": err.Error(),
			})
			return
		}

		_, OK := token.Claims.(jwt.MapClaims)

		if !OK {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Message": "unable to parse claims",
			})
			return
		}
		c.Next()
	}
}
func Run() {

	router := gin.Default()

	router.GET("/", Version)
	router.GET("/user/get/:userid", jwtMiddleware(), UserGetById)
	router.POST("/user/register", UserRegister)
	router.POST("/login", Login)

	router.Run()
}
