package main

import (
	"strconv"
	"strings"
)

func convertHex(text string) string {

	words := strings.Split(text, " ") //["Simply", "add", "1E", "(hex)"]

	for i, value := range words {
		if value == "(hex)" && i > 0 {

			back := words[i-1]

			decimal, _ := strconv.ParseInt(back, 16, 64)

			words[i-1] = strconv.Itoa(int(decimal))
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
