package messagebroker

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client
var lock = &sync.Mutex{}

// GetClient implements the design pattern singleton
// to just have one redis client
func GetClient() *redis.Client {
	if rdb == nil {
		lock.Lock()
		defer lock.Unlock()
		if rdb == nil {
			connect()
			fmt.Println("Connected with redis.")
		}
	}

	return rdb
}

func connect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // password set
		DB:       0,  // use default DB
	})
}
