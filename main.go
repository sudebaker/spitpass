package main

import (
    "fmt"
    "spitpass/stuff"
)
func main() {

    wordlist := stuff.GetWords("./texto.txt",8)

    for _,word := range(wordlist){
        password:=stuff.NewPassword()
        password.Password = word
        password.ClearString()
        password.ConvertEleet()
        fmt.Println(password)
    }
}
