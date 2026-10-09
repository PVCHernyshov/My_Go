package main

import (
	"fmt"
	"math"
)

/*
func main() {

		division(10, 2) 10,0; 0,5; -10,2

		extractingSquare(4) -4; 0

		factorial(3) -5; 0

	fmt.Println(isPrime(9))
}
*/

func division(x, y float64) (float64, error) {
	if y != 0 {
		return float64(x) / float64(y), nil
	}
	return float64(y), fmt.Errorf("Ошибка! Деление %f на %f.", x, y)
}

func extractingSquare(n float64) (float64, error) {
	if n >= 0 {
		return math.Sqrt(n), nil
	}
	return 0, fmt.Errorf("Ошибка! Невозможно извлечь корень из %.2f.", n)
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Факториал для числпа %d не определён.", n)
	}
	if n < 2 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("Число должно быть больше 1. Получено %d.", n)
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil

}
