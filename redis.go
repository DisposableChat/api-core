package core

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Host     string
	Port     int16
	Username  string
	Password string
	Client   *redis.Client
}

func (r *Redis) Init() (*redis.Client) {
	if r.Host == "" || r.Port < 0 || r.Username == "" || r.Password == "" {
		panic("REDIS_ENVIROMENT_VARIABLES not set up correctly")
	}

	r.Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.Host, r.Port),
		Username: r.Username,
		Password: r.Password,
		DB: 0,
	})

	_, err := r.Client.Ping(r.Client.Context()).Result()
	if err != nil {
		panic(err)
	}

	return r.Client
} 