package main

import "testing"

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name          string
		inputData     string
		expectedData  float64
		expectedError bool
	}{
		{"Валидные данные", "14.15", 14.15, false},
		{"Запятая вместо точки", "14,15", 0, true},
		{"Пробел вместо точки", "14 15", 0, true},
		{"Буквы вместо цифр", "abs", 0, true},
		{"Пустая строка", "", 0, true},
		{"Смешанная строка", "14abs", 0, true},
		{"Пробел перед цифрами", " 14.15", 0, true},
		{"Пробел после цифр", "14.15 ", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseNumber(tt.inputData)

			if (err != nil) != tt.expectedError {
				t.Fatalf("Ожидалось %v Получено %v", tt.expectedError, err)
			}

			if !tt.expectedError && result != tt.expectedData {
				t.Errorf("Ожидалось %f Получено %f", tt.expectedData, result)
			}
		})
	}
}
