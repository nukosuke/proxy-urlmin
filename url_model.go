package main

import (
    _ "os"
)
//URL ID = 10桁
//使用できる文字59文字 => 59進数表示
const URLCharacter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.!'()"

type URL struct {
    index [10]int8 //0~58
}

func NewURL() *URL {
    u := new(URL)
    return u
}

func (this *URL) Save(url string) {
    isCarry := false
    
    this.index[0]++
    if this.index[0] >= 59 {
        this.index[0] -= 59
        isCarry = true
    }
    
    for i:=1; i<10; i++ {
        if isCarry {
            this.index[i]++
            isCarry = false
        }
        
        if this.index[i] >= 59 {
            this.index[i] -= 59
            isCarry = true
        }
    }
    
    
}

func (this *URL) Find(url string) {
    //find from redis
    //return 
}