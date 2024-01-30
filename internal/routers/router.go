package router

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func UserRegister(c *gin.Context) {

}

func UserGetById(c *gin.Context) {

}

func Run() {

	router := gin.Default()

	router.GET("/login", Login)
	router.POST("/user/register", UserRegister)
	router.POST("/user/get/{id}", UserGetById)

	router.Run()
}
