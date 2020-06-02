package database

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

var Cache redis.Conn

func InitCache() {
	conn, err := redis.DialURL("redis://localhost")

	if err != nil {
		fmt.Println("error with redis")
		log.Fatal(err)
	}
	Cache = conn
}
