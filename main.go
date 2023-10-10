package main

import (
	"fmt"
	_ "fmt"
	"log"
	"os"
	"spitpass/stuff"

	"github.com/urfave/cli/v2"
)

func main() {

	//     wordlist := stuff.GetWords("./texto.txt",10)

	//     for _,word := range(wordlist){
	//         password:=stuff.NewPassword()
	//         password.Password = word
	//         fmt.Println(password.Password)
	//     }
	//

	pass := stuff.NewPassword()
	// pass.GenerateRandomPass()
    // pass.SuffixPass(pass.Suffix_lenght)
    // pass.CapFirst = true
    // pass.CapitalizePass()
	// fmt.Println(pass.Password)
    app := &cli.App{
        Name: "SpitPass",
        Usage: "SpitPass",
        Action: func(*cli.Context) error {
            pass.GenerateRandomPass()
            fmt.Println(pass.Password)
            return nil
        },

    }
    if err := app.Run(os.Args); err !=nil {
        log.Fatal(err)
    }
}
