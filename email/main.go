package main

import (
	"errors"
	"strings"
)

func main() {
	validateEmail("q@gmail.com")

}
func validateEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("Пустой email")
	}

	if strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") ||
		strings.HasPrefix(email, ".") || strings.HasSuffix(email, ".") ||
		strings.Count(email, "@") < 1 || strings.Count(email, "@") > 1 {
		return false, errors.New("Не верный формат")
	}

	parts := strings.Split(email, "@")
	if strings.Count(parts[1], ".") < 1 {
		return false, errors.New("Не верный формат")
	}
	return true, nil
}
