package main

import (
	validator "github.com/asaskevich/govalidator"
	"errors"
)

//URL ID = 10桁
//使用できる文字69文字 => 69進数表示
const URLCharacter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.!'()"
const URLAvailable = 69

type URL struct {
	index [10]int8 //0~68
}

func NewURL() *URL {
	this := new(URL)
	return this
}

func (this *URL) Save(url string) (string, error) {
	if !this.validate(url) {
		return "", errors.New("URL validation error")
	}
	
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

	return id, nil
}

func (this *URL) Find(url string) {
	//TODO: validate id
	//find from redis
	//return
}

func (this *URL) validate(url_string string) bool {
	return validator.IsURL(url_string)
}
