package main

import (
	"strconv"
	"strings"
)

func convertHex(text string) string {

	words := strings.Split(text, " ") //["Simply", "add", "1E", "(hex)"]

	for i, value := range words {
		if value == "(hex)" {

			back := words[i-1]

			//covert back to intger

			decimal, _ := strconv.ParseInt(back, 16, 64)

			text = strconv.Itoa(int(decimal))
		}
	}

	return text
}
