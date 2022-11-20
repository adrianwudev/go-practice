package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := camelCase(input)
	fmt.Println("Answer: ", answer)
}

const lowerCase = 97

func camelCase(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	var howManyWords int32
	howManyWords = 1
	for _, rune := range s {
		if rune < lowerCase {
			howManyWords++
		}
	}
	return howManyWords
}
