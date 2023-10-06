package stuff

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getLeetList() [][]string {
	A := []string{"A", "4", "/\\", "@", "/-\\", "^", "aye", "(L", "Д"}
	B := []string{"B", "I3", "8", "13", "|3", "ß", "P>", "|:", "!3", "(3", "/3", ")3", "|-]", "j3"}
	C := []string{"C", "[", "¢", "<", "(", "©"}
	D := []string{"D", ")", "|)", "(|", "|o", "[)", "I>", "|>", "T)", "I7", "cl", "|}", "|]"}
	E := []string{"E", "3", "&", "£", "€", "ë", "[-", "|=-"}
	F := []string{"F", "|=", "ƒ", "|#", "ph", "/=", "v"}
	G := []string{"G", "6", "&", "(_+", "9", "C-", "gee", "(?,", "[,", "{,", "<-", "(.", "₲"}
	H := []string{"H", "#", "/-/", "[-]", "]-[", ")-(", "(-)", ":-:", "|~|", "|-|", "]~[", "}{", "!-!", "1-1", "-/", "I+I", "}-{"}
	I := []string{"I", "1", "!", "¡", "|", "eye", "3y3", "][", "]", "/me"}
	J := []string{"J", "_|", "_/", "¿", "</", "_]", "(/"}
	K := []string{"K", "X", "|<", "|{", "]{", "|X"}
	L := []string{"L", "1", "£", "7", "1_", "|", "|_", "el", "[]_", "[_"}
	M := []string{"M", "|v|", "[V]", "{V}", "\\/V\\", "em", "AA", "|\\/|", "/\\/\\", "(u)", "(V)", "(/)", "/|\\", "^^", "/|/|", "//\\", "||\\", "]\\/["}
	N := []string{"N", "^/", "|\\|", "/\\/", "[\\]", "<\\>", "en", "[]\\", "//", "[]", "/V", "1V", "И", "^", "ท"}
	O := []string{"O", "0", "()", "oh", "[]", "p", "<>", "Ø"}
	P := []string{"P", "|*", "|o", "|º", "|^(o)", "|>", "|\"", "9", "[]D", "|°", "|7"}
	Q := []string{"Q", "(_,)", "()_", "0_", "<|", "&"}
	R := []string{"R", "I2", "|`", "|~", "|?", "/2", "|^", "lz", "|9", "2", "12", "®", "[z", "Я", ".-", "|2", "|-"}
	S := []string{"S", "5", "$", "z", "§", "ehs", "es", "2"}
	T := []string{"T", "7", "+", "-|-", "']['", "†", "\\\"|\\"}
	U := []string{"U", "(_)", "|_|", "v", "L|", "µ", "บ"}
	V := []string{"V", "\\/", "|/", "\\|"}
	W := []string{"W", "\\/\\/", "vv", "\\N", "'//", "\\\\'", "\\^/", "dubya", "(n)", "\\V/", "\\X/", "\\|/", "\\_|_/", "\\_:_/", "Ш", "uu", "2u", "\\\\//\\\\//", "พ", "ω"}
	X := []string{"X", "><", "Ж", "}{", "ecks", "×", "?", ")(", "][", "}{"}
	Y := []string{"Y", "j", "`/", "Ч", "7", "\\|/", "¥"}
	Z := []string{"Z", "2", "7_", "-/_", "%", ">_", "s", "~/_", "-\\_", "-|_"}

	return [][]string{A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z}
}

func getLeetIndex(c string, leetIndex []string) int {
	for p, v := range leetIndex {
		if c == v {
			return p
		}
	}

	return -1
}

func Leet(input string) string {
	leets := getLeetList()

	var leetIndex []string
	for i := 0; i < len(leets); i++ {
		leetIndex = append(leetIndex, leets[i][0])
	}

	leetStr := ""

	for _, c := range input {
		upperC := strings.ToUpper(string([]rune{c}))

		n := getLeetIndex(upperC, leetIndex)
		if n > 0 {
			rnum := rand.Intn(len(leets[n]))
			leetStr = leetStr + leets[n][rnum]
		} else {
			leetStr = leetStr + string(c)
		}

	}

	return leetStr
}