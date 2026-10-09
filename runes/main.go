package main

import (
	"errors"
	"strconv"
)

func main() {

	parseNumber("45,15 ")

}

func parseNumber(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New("Введены ошибочные данные")
	}
	return res, nil
}
