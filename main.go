package main

import (
	"fmt"
	"spitpass/stuff"
)
func main() {

    wordlist := stuff.GetWords("./texto.txt",10)

    for _,word := range(wordlist){
        password:=stuff.NewPassword()
        password.Password = word
        fmt.Println(password.Password)
    }
}
