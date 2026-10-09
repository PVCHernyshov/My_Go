package main

import (
	"testing"
)

// Это table-driven тест для функции division
func TestDivision(t *testing.T) {
	// 1. Создаем "срез тестовых случаев" (список сценариев)
	// Мы используем анонимную структуру, описывая поля прямо здесь
	tests := []struct {
		name     string  // Название теста (чтобы видеть, какой упал)
		a, b     float64 // Входные данные
		expected float64 // Ожидаемый результат
		wantErr  bool    // Ожидаем ли мы ошибку? (true = да, false = нет)
	}{
		// 2. Заполняем список конкретными случаями
		{"Деление 10 на 2", 10, 2, 5, false},
		{"Деление на ноль", 10, 0, 0, true},
		{"Деление 0 на 5", 0, 5, 0, false},
		{"Деление отрицательного", -10, 2, -5, false},
	}

	// 3. "Итерация по тестам" из плана
	for _, tt := range tests {
		// t.Run запускает каждый случай как отдельный подтест с именем tt.name
		t.Run(tt.name, func(t *testing.T) {

			// Вызываем тестируемую функцию с данными из текущего случая
			result, err := division(tt.a, tt.b)

			// Проверяем ошибку: если мы ждали ошибку, а её нет (или наоборот)
			if (err != nil) != tt.wantErr {
				t.Errorf("division() error = %v, wantErr %v", err, tt.wantErr)
				return // Прерываем этот конкретный тест, идем к следующему
			}

			// Проверяем результат (только если ошибки не было)
			if !tt.wantErr && result != tt.expected {
				t.Errorf("division() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractingSquare(t *testing.T) {

	test := []struct {
		name    string  // имя теста
		a       float64 // входные данные
		res     float64 // ожидаемый результат
		errRetu bool    // ожидаем ли ошибку (true = да, false = нет)
	}{
		{"Извлечение корня из положительного числа", 4, 2, false},
		{"Извлечение корня из отрицательного числа", -4, 0, true},
		{"Извлечение корня из 0", 0, 0, false},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractingSquare(tt.a)

			if (err != nil) != tt.errRetu {
				t.Errorf("Результат теста %v. Ожидаемый результат %v", err, tt.errRetu)
				return
			}
			if !tt.errRetu && result != tt.res {
				t.Errorf("Результат теста %f. Ожидаемый результат %f", result, tt.res)
			}

		})
	}
}

func TestFactorial(t *testing.T) {

	test := []struct {
		name    string
		a       int
		res     int
		errRetu bool
	}{
		{"Факториал отрицательного числа", -1, 0, true},
		{"0!", 0, 1, false},
		{"1!", 1, 1, false},
		{"Факториал положительного числа", 2, 2, false},
	}
	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {

			result, err := factorial(tt.a)

			if (err != nil) != tt.errRetu {
				t.Errorf("Ожидаемый результат %v, получен %v", tt.errRetu, err)
				return
			}
			if !tt.errRetu && result != tt.res {
				t.Errorf("Ожидаемый результат %d, получен %d", tt.res, result)
			}
		})
	}

}

func TestIsPrime(t *testing.T) {

	test := []struct {
		name    string
		inpData int
		res     bool
		err     bool
	}{
		{"=Проверка с отрицательным числом -1 =", -1, false, true},
		{"=Исключение число 1 =", 1, false, true},
		{"=Простое число 2 =", 2, true, false},
		{"=Простое число 7 =", 7, true, false},
		{"=Составное число 6 =", 6, false, false},
		{"=Составное число 9 =", 9, false, false},
	}
	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {

			result, err := isPrime(tt.inpData)

			if (err != nil) != tt.err {
				t.Errorf("Ожидаемый результат %v полученый результат %v", tt.err, err)
				return
			}
			if !tt.err && result != tt.res {
				t.Errorf("Ожидаемый результат %v полученый результат %v", tt.res, result)
			}

		})

	}
}
