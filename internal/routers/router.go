package router

import (
	"net/http"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {

	res := []string{"v:1"}
	c.JSON(200, res)

}

func Login(c *gin.Context) {

}

func UserRegister(c *gin.Context) {
	var form entities.Form

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	user.RegisterUserHandler(&form)

	c.String(200, "Hello %s", form.FirstName)
}

func UserGetById(c *gin.Context) {

}

func Run() {

	router := gin.Default()

	router.GET("/", Version)
	router.GET("/user/get/{id}", UserGetById)

	router.POST("/user/register", UserRegister)
	router.POST("/login", Login)

	router.Run()
}
