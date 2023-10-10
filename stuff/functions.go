package stuff

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const special_chars = "!@#$%&?,;.:-_<>"

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

//var nonAlphanumericRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

type pass_struct struct {
	Password       string
	Length         int
	Preffix        bool
	Preffix_lenght int
	Suffix         bool
	Suffix_lenght  int
	Eleet          bool
	CapFirst       bool
	CapLast        bool
}

func (p *pass_struct) CapitalizePass() {
	//capitalize by default first letter except lastletter = True
	// Convert the string to lowercase

	// Check the provided option and capitalize the corresponding letter
	if p.CapFirst {
		if len(p.Password) > 0 {
			p.Password = strings.ToUpper(string(p.Password[0])) + p.Password[1:]
		}
	} else if p.CapLast {
		if len(p.Password) > 0 {
			p.Password = p.Password[:len(p.Password)-1] + strings.ToUpper(string(p.Password[len(p.Password)-1]))
		}
	}
}

func (p *pass_struct) PreffixPass(lenght int) {
	//add a random preffix with lenght <lenght>
	p.Password = chooseRandomChar(special_chars, lenght) + p.Password
}

func (p *pass_struct) SuffixPass(lenght int) {
	//add a random suffix with lenght <lenght>
	p.Password = p.Password + chooseRandomChar(special_chars, lenght)
}

func (p *pass_struct) ConvertEleet() {
	//convert string password to eleet language
	fmt.Println("Convert to leet function", p)
}

func NewPassword() pass_struct {
	password := pass_struct{
		Password:       "",
		Length:         10,
		Preffix:        false,
		Preffix_lenght: 0,
		Suffix:         false,
		Suffix_lenght:  0,
		Eleet:          false,
		CapFirst:       false,
		CapLast:        false,
	}
	return password
}

func chooseRandomChar(s string, n int) string {

	rand.Seed(time.Now().UnixNano())
	// Convert the string to a slice of runes
	chars := []rune(s)
	// Create a new slice to store the randomly chosen characters
	randomChars := make([]rune, n)

	// Generate n random indices in the range 0 to len(chars)-1
	// and select the corresponding characters
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(chars))
		randomChars[i] = chars[randomIndex]
	}

	// Convert the slice of runes back to a string
	return string(randomChars)
}

func GetWords(filepath string, wordlenght int) []string {
	wordlist := []string{}
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the file: %v\n", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		// Split the line into words
		words := strings.Fields(scanner.Text())

		// Iterate over each word
		for _, word := range words {
			// Check if the word length is greater than x
			if len(word) > wordlenght {
				word = strings.ToLower(word)
				word = nonAlphanumericRegex.ReplaceAllString(word, "")
				wordlist = append(wordlist, word)
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan the file: %v\n", err)
	}
	return wordlist
}
