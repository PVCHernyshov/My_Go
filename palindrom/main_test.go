package main

import "testing"

func TestPalindrom(t *testing.T) {
	tests := []struct {
		name         string
		inputData    string
		expectedData bool
	}{
		{"палиндром кириллица + разный регистр", "А роза упала на лапУ азора", true},
		{"палиндром латиница + разный регистр + знак припенания", "Was it a car or a cat I saw?", true},
		{"пустая строка", "", true},
		{"один символ", "А", true},
		{"строка пробелов", "   ", true},
		{"многоточие", "...", true},
		{"знаки препинания + пробелы + симметрично", " ! . ! ", true},
		{"не палиндром кирилица", "Привет", false},
		{"не палиндром латиница", "Hello", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := palindrom(tt.inputData)

			if result != tt.expectedData {
				t.Errorf("ожидалось %v получено %v", tt.expectedData, result)
			}
		})
	}

}
