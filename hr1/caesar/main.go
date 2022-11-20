package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}

	fmt.Println(string(ret))
}

func cipher(r rune, delta int) rune {
	if isUpper(r) {
		return rotate(r, 'A', delta)
	}

	if isLower(r) {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
