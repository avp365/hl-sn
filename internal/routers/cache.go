package router

import (
	"github.com/avp365/hl-sn/internal/pkg/configs"
	"github.com/go-redis/redis"
)

var cache = redis.NewClient(&redis.Options{
	Addr: configs.GetRedisUrl(),
})
