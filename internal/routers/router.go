package router

import (
	"net/http"

	"github.com/avp365/hl-sn/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Version(c *gin.Context) {

	res := []string{"v:1"}
	c.JSON(200, res)

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

		data := token.Claims.(jwt.MapClaims)

		c.Params = append(c.Params, gin.Param{Key: "id", Value: data["id"].(string)})
		c.Next()
	}
}
func Run() {

	router := gin.Default()

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(router)

	router.GET("/", Version)
	router.GET("/user/get/:userid", jwtMiddleware(), UserGetById)
	router.GET("/user/search", jwtMiddleware(), UserSearch)
	router.POST("/user/register", UserRegister)

	router.POST("/friend/set/:friendid", jwtMiddleware(), FriendSet)
	router.POST("/friend/delete/:friendid", jwtMiddleware(), FriendDelete)

	router.POST("/post/create", jwtMiddleware(), PostCreate)
	router.POST("/post/delete/:postid", jwtMiddleware(), PostDelete)
	router.POST("/post/update", jwtMiddleware(), PostUpdate)
	router.POST("/post/get/:postid", jwtMiddleware(), PostGet)
	router.POST("/post/feed", jwtMiddleware(), PostFeed)

	router.POST("/login", Login)

	router.Run()
}
