package main

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var ErrNoAlbum = errors.New("no album found")

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func FindAlbum(id string) (*Album, error) {

	conn := pool.Get()

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "album:"+id))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, ErrNoAlbum
	}

	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		return nil, err
	}

	return &album, nil
}
