package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Do("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Electric Ladyland added!")
}
