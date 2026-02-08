package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertHex(text string) string {

	words := strings.Split(text, " ")

	for i, value := range words {
		if value == "(hex)" && i > 0{
			words[i] = ""

			back := words[i-1] 

			decimal, err := strconv.ParseInt(back, 16, 64)
			if err != nil {
				fmt.Println("Error Hex decimal:", err)
			}

			words[i-1] = strconv.Itoa(int(decimal))
		}
	}
	return strings.Join(words, " ")
}
