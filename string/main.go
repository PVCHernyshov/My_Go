package main

import (
	"errors"
)

func main() {

	substring("Hello World!", 2, 8)

}

func substring(s string, start, end int) (string, error) {
	if start < 0 || start >= end || end > len(s) {
		return "", errors.New("Ошибка. Указано не верное значение индексов.")
	}
	return s[start:end], nil
}
