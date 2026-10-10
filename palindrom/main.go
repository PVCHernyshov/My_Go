package main

import (
	"strings"
	"unicode"
)

func main() {

	str := "А роза упала на лапу азора"

	palindrom(str)

}

func palindrom(s string) bool {
	res := strings.Map(sorting, s)
	result := []rune(res)
	size := len(result)

	for i := 0; i < size/2; i++ {
		if result[i] != result[size-1-i] {
			return false
		}
	}

	return true
}

func sorting(r rune) rune {
	if unicode.IsSpace(r) || unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}
