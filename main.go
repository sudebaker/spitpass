package main

import (
	"fmt"
	"log"
	"os"
	"spitpass/stuff"
	"time"

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
		Name:     "SpitPass",
		Version:  "0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Jesus Cifuentes",
				Email: "jcifuentesfonseca@protonmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "suffix",
				Value: true,
				Usage: "Add a random special character at the end of password",
				Destination: &pass.Suffix,
			},
			&cli.BoolFlag{
				Name: "preffix",
				Value: true,
				Usage: "Add a random special character at the beginning of password",
				Destination: &pass.Preffix,
			},
			&cli.BoolFlag{
				Name: "eleet",
				Value: true,
				Usage: "Convert password to a semi eleet language",
				Destination: &pass.Eleet,
			},
			&cli.BoolFlag{
				Name: "capfirst",
				Value: true,
				Usage: "Capitalize the first letter of password",
				Destination: &pass.CapFirst,
			},
			&cli.BoolFlag{
				Name: "caplast",
				Value: true,
				Usage: "Capitalize the last letter of password",
				Destination: &pass.CapLast,
			},
		},
		Usage: "SpitPass",
		Action: func(*cli.Context) error {
			pass.GenerateRandomPass()
			fmt.Println(pass.Password)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
