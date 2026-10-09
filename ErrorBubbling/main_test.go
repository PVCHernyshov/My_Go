package main

import (
	"math"
	"testing"
)

func TestComplexCalculation(t *testing.T) {

	test := []struct {
		name   string  // имя теста
		x, y   float64 // входные данные
		resRet float64 // ожидаемый результат
		errRet bool    // должна ли вернутся ошибка
	}{
		{"Деление на отрицательное число", 10, -5, 0, true},
		{"Деление на 0", 10, 0, 0, true},
		{"деление на положительное число", 10, 5, math.Sqrt(10 / 5), false},
	}
	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {
			result, err := complexCalculation(tt.x, tt.y)
			if (err != nil) != tt.errRet {
				t.Errorf("Ожидаемый результат %v полученый %v", tt.errRet, err)
				return
			}
			if !tt.errRet && result != tt.resRet {
				t.Errorf("Ожидаемый результат %.2f полученый %.2f", tt.resRet, result)
			}
		})
	}
}
