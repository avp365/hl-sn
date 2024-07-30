package router

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/handlers/post"
	"github.com/avp365/hl-sn/internal/handlers/user"
	"github.com/avp365/hl-sn/internal/pkg/configs"
	"github.com/go-redis/redis"
)

var cache = redis.NewClient(&redis.Options{
	Addr: configs.GetRedisUrl(),
})

var queueName = "queue_post_update"

var queue = redis.NewClient(&redis.Options{
	Addr: configs.GetRedisUrl(),
})

func CacheUpdate() {

	for {
		time.Sleep(1 * time.Second)

		result, err := queue.LPop(queueName).Result()

		if err != nil {
			fmt.Println(err)
		}

		id, err := strconv.Atoi(result)

		if err != nil {
			fmt.Println(err)
			continue
		}

		friendIds, err := user.GetFriends(id)

		if err != nil {

			fmt.Println(err)
			continue

		}

		fmt.Println(friendIds)

		for _, id := range friendIds {
			_, err := feedCache(id)
			fmt.Println(err)
		}

	}

}

func feedCache(userId int) ([]entities.Post, error) {

	posts, err := post.Feed(userId)

	if err != nil {
		return nil, err
	}

	data, errJson := json.Marshal(posts)

	if errJson != nil {
		return nil, err
	}

	cacheErr := cache.Set(strconv.Itoa(userId), data, 48*time.Hour).Err()

	if cacheErr != nil {
		return nil, err
	}

	return posts, nil
}
