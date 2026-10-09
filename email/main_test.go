package main

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	test := []struct {
		name  string
		email string
		err   bool
	}{
		{"Пустая строка", "", true},
		{"Без имени", "@mail.ru", true},
		{"Опечатка. Нет \"@\"", "abcdemail.ru", true},
		{"Опечатка. Много \"@@\"", "abcde@@mail.ru", true},
		{"Опечатка. Нет \".\"", "abcde@mailru", true},
		{"Опечатка. Много \".\"", "abcde@mail.ru.", true},
		{"Полностью валидный", "abcde@mail.ru", false},
		{"Полностью валидный c поддоменом", "abcde@mail.example.ru", false},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result, err := validateEmail(tt.email)

			if (err != nil) != tt.err {
				t.Fatalf("Ожидалось %v получено %v", tt.err, err)
			}
			if !tt.err && !result {
				t.Errorf("Ожидалось %v получено %v", tt.err, result)
			}
		})
	}
}
