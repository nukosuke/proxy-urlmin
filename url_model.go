package main

import (
	"errors"
	validator "github.com/asaskevich/govalidator"
	"github.com/garyburd/redigo/redis"
	"log"
)

//URL ID = 10桁
//使用できる文字69文字 => 69進数表示
const URLCharacter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.!'()"
const URLAvailable = 69

type URL struct {
	index [10]int8 //0~68
	conn  redis.Conn
}

func NewURL(redis_connection redis.Conn) *URL {
	if redis_connection == nil {
		log.Fatal("Invalid Redis Connection")
	}

	//TODO
	// load start_id from redis to this.index

	this := new(URL)
	this.conn = redis_connection
	return this
}

func (this *URL) Save(url string) (string, error) {
	if !this.validateUrl(url) {
		return "", errors.New("URL validation error")
	}

	// 次のURL IDの払い出し
	id := ""
	for i := 9; i >= 0; i-- {
		id += string(URLCharacter[this.index[i]])
	}

	isCarry := false

	this.index[0]++
	if this.index[0] >= URLAvailable {
		this.index[0] -= URLAvailable
		isCarry = true
	}

	for i := 1; i < 10; i++ {
		if isCarry {
			this.index[i]++
			isCarry = false
		}

		if this.index[i] >= URLAvailable {
			this.index[i] -= URLAvailable
			isCarry = true
		}
	}

	_, err := this.conn.Do("SET", id, url)
	if err != nil {
		return "", errors.New("Failed to save url to KVS")
	} else {
		return id, nil
	}

}

func (this *URL) Find(id string) (string, error) {
	//TODO: validate id
	res, err := redis.String(this.conn.Do("GET", id))
	if err != nil {
		return "", errors.New("Failed to GET url")
	} else {
		return res, nil
	}
}

func (this *URL) validateUrl(url_string string) bool {
	return validator.IsURL(url_string)
}

func (this *URL) validateId(url_id string) bool {
	//TODO RegExp("^(0-9a-zA-Z)*$")
	return true
}
