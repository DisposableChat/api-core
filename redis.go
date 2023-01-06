package core

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Host     string
	Port     int16
	Password string
	Client   *redis.Client
}

func (r *Redis) Init() (*redis.Client) {
	r.Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB: 0,
	})

	_, err := r.Client.Ping(r.Client.Context()).Result()
	if err != nil {
		panic(err)
	}

	return r.Client
} 