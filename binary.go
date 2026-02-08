package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertBin(text string) string {

	words := strings.Split(text, " ")

	for i, value := range words {
		if value == "(bin)" && i > 0 {
			words[i] = ""

			back := words[i-1]

			decminal, err := strconv.ParseInt(back, 2, 64)
			if err != nil {
				fmt.Println("Error Binary decmial", err)
			}
			words[i-1] = strconv.Itoa(int(decminal))
		}
	}
	return strings.Join(words, " ")
}