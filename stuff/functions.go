package stuff

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	conso = [...]rune{
		'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}
	vowels        = [...]rune{'a', 'e', 'i', 'o', 'u'}
	special_chars = "!@#$%&?,;.:-_<>"
	alphabet      = map[rune]rune{
		'a': '4',
		'b': '8',
		'e': '3',
		'g': '6',
		'h': '#',
		'l': '1',
		'o': '0',
		's': '5',
		't': '7',
		'z': '2',
	}
	//nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)
)

var nonAlphanumericRegex = regexp.MustCompile("^[a-zA-Z]+$")

// struct to store password and options
type pass_struct struct {
	Password       string
	File           string
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
	//capitalize password first letter, last letter or both

	// Check the provided option and capitalize the corresponding letter
	if p.CapFirst {
		if len(p.Password) > 0 {
			p.Password = strings.ToUpper(string(p.Password[0])) + p.Password[1:]
		}
	}
	if p.CapLast {
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
	//convert string password to semi eleet language
	var ptemp string
	for _, char := range p.Password {
		if replacement, ok := alphabet[char]; ok {
			ptemp += string(replacement)
		} else {
			ptemp += string(char)
		}
	}
	p.Password = ptemp
}

func (p *pass_struct) GenerateRandomPass() {
	//totally human readable random pass
	//rand.Seed(time.Now().UTC().UnixNano())
	tsize := int(p.Length / 2)
	tpass := ""
	index := 0
	for item := 0; item <= tsize; item++ {
		index = rand.Intn(len(conso))
		tpass = tpass + string(conso[index])
		index = rand.Intn((len(vowels)))
		tpass = tpass + string(vowels[index])
	}
	p.Password = tpass
}

func (p *pass_struct) GeneratePassword() {
	//generate password from file or random one and modify it according to args

	if p.File != "" {
		p.GenerateFromFile()
	} else {
		p.GenerateRandomPass()
	}

	//order of checking args is important
	if p.Eleet {
		p.ConvertEleet()
	}
	if p.CapFirst || p.CapLast {
		p.CapitalizePass()
	}
	if p.Preffix {
		p.PreffixPass(p.Preffix_lenght)
	}
	if p.Suffix {
		p.SuffixPass(p.Suffix_lenght)
	}
}

func NewPassword() pass_struct {
	password := pass_struct{
		Password:       "",
		Length:         8,
		Preffix:        false,
		Preffix_lenght: 2,
		Suffix:         false,
		Suffix_lenght:  2,
		Eleet:          false,
		CapFirst:       false,
		CapLast:        false,
	}
	return password
}

func chooseRandomChar(s string, n int) string {

	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	// Convert the string to a slice of runes
	chars := []rune(s)
	// Create a new slice to store the randomly chosen characters
	randomChars := make([]rune, n)

	// Generate n random indices in the range 0 to len(chars)-1
	// and select the corresponding characters
	for i := 0; i < n; i++ {
		randomIndex := randGenerator.Intn(len(chars))
		randomChars[i] = chars[randomIndex]
	}

	// Convert the slice of runes back to a string
	return string(randomChars)
}

func (p *pass_struct) GenerateFromFile() {
	// Open the file
	file, err := os.Open(p.File)
	if err != nil {
		fmt.Printf("Failed to open the file: %v", file)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process the file
	var passwordComponents []string // To store potential components
	currentComponent := ""
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, word := range fields {
			// Check if the word can be added to the current component
			if nonAlphanumericRegex.MatchString(word) {
				// If the word is alphanumeric
				if len(word) == p.Length {
					// we got it
					passwordComponents = append(passwordComponents, word)
					continue
				}
				newComponent := currentComponent + word // Combine with the previous component

				if len(newComponent) <= p.Length {
					currentComponent = newComponent
				} else {
					// If component gets too long...
					currentComponent = newComponent[:p.Length] // Trim it
					passwordComponents = append(passwordComponents, currentComponent) // Store it
					currentComponent = ""                                                                  // Start a new component
				}
			}
		}

		// // Handle the last potential component
		// if currentComponent != "" {
		// 	passwordComponents = append(passwordComponents, currentComponent)
		// }

		for _, word := range passwordComponents {
			p.Password = word
			if p.Eleet {
				p.ConvertEleet()
			}
			if p.CapFirst || p.CapLast {
				p.CapitalizePass()
			}
			if p.Preffix {
				p.PreffixPass(p.Preffix_lenght)
			}
			if p.Suffix {
				p.SuffixPass(p.Suffix_lenght)
			}
			fmt.Println(p.Password)
		}

	}
}
