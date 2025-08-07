package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strs := []string{"neet", "code", "love", "you"}
	var concatString strings.Builder

	for _, str := range strs {
		concatString.WriteString(strconv.Itoa(len(str)))
		concatString.WriteString("#")
		concatString.WriteString(str)
	}

	fmt.Println("Encoded string:", concatString.String())
	stringToDecode := concatString.String()
	i := 0
	var result []string

	for i < len(stringToDecode) {
		j := i
		for j < len(stringToDecode) && stringToDecode[j] != '#' {
			j++
		}

		length, err := strconv.Atoi(stringToDecode[i:j])
		if err != nil {
			fmt.Println("Error decoding length:", err)
			return
		}

		i = j + 1
		if i+length > len(stringToDecode) {
			fmt.Println("Error: not enough characters for the specified length")
			return
		}

		str := stringToDecode[i : i+length]
		fmt.Printf("Decoded string: %s\n", str)
		result = append(result, str)
		i += length
	}

	fmt.Println("Decoding completed successfully.", result)

}
