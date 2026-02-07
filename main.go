package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Not exactly 2 arugements")
		return
	}

	inputfile := args[0]
	outputfile := args[1]

	data, err := os.ReadFile(inputfile)

	if err != nil {
		fmt.Println("Error Reading:", err)
		return
	}

	text := string(data)

	text = convertHex(text)

	err = os.WriteFile(outputfile, []byte(text), 0644)

	if err != nil {
		fmt.Println("Error Writing:", err)
		return
	}

	fmt.Println("\n successfully create a file")

}
