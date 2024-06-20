package main

import "fmt"

var cifra = map[rune]string{
	'a': "a",
	'b': "bac",
	'c': "cad",
	'd': "def",
	'e': "e",
	'f': "feg",
	'g': "geh",
	'h': "hij",
	'i': "i",
	'j': "jik",
	'k': "kil",
	'm': "mon",
	'n': "nop",
	'o': "o",
	'p': "poq",
	'q': "qor",
	'r': "ros",
	's': "sot",
	't': "tuv",
	'u': "u",
	'v': "vux",
	'x': "xuz",
	'z': "zuz",
}

func main () {
	var palavra string;
	nova := ""

	fmt.Scanf("%s", &palavra);

	for _, v := range palavra {
		nova += cifra[v]
	}

	fmt.Println(nova)
}