package main

import (
	"fmt"
	"math"
)

func main() {

	Valid(complexCalculation(10, 2))

	Valid(complexCalculation(10, -2))

	Valid(complexCalculation(10, 0))

}

func complexCalculation(x, y float64) (float64, error) {

	return division(x, y)
}

func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("Ошибка! Деление на 0.")
	}
	return extractingSquare(x / y)
}

func extractingSquare(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Ошибка! Получено отрицательное значение.")
	}
	return math.Sqrt(n), nil
}

func Valid(n float64, e error) {
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(n)
	}
}
