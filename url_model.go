package main

import (
	"net/url"
	_ "os"
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

func (this *URL) Save(url string) string {
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

	return id
}

func (this *URL) Find(url string) {
	//find from redis
	//return
}

func (this *URL) validate(url_string string) bool {
	var isValid bool

	_, err := url.Parse(url_string)

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid
}
