package main

import "testing"

func TestAddAddress(t *testing.T) {
	tests := []struct {
		name            string
		address         map[string]map[string]map[string]struct{}
		inputDataCity   string
		inputDataStreet string
		inputDataHouse  string
		returnError     bool
	}{
		{"полный, валидный адрес: город+улица+дом", map[string]map[string]map[string]struct{}{}, "Липецк", "Циолковского", "14", false},
		{"неполный, валидный адрес: город+улица", map[string]map[string]map[string]struct{}{}, "Липецк", "Циолковского", "", false},
		{"неполный, валидный адрес: город", map[string]map[string]map[string]struct{}{}, "Липецк", "", "", false},
		{"невалидный адрес: улица+дом", map[string]map[string]map[string]struct{}{}, "", "Циолковского", "14", true},
		{"невалидный адрес: только дом", map[string]map[string]map[string]struct{}{}, "", "", "14", true},
		{"невалидный адрес: только улица", map[string]map[string]map[string]struct{}{}, "", "Циолковского", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := addAddress(tt.address, tt.inputDataCity, tt.inputDataStreet, tt.inputDataHouse)

			if (err != nil) != tt.returnError {
				t.Fatalf("ожидалось %v получено %v", tt.returnError, err)
			}

			if _, ok := tt.address[tt.inputDataCity]; ok == tt.returnError { // если не ждём ошибку значит должно быть записано. Если ждём == пусто.
				t.Fatalf("ожидалось %s получено иное значение", tt.inputDataCity)
			}

			if tt.inputDataStreet != "" {
				if _, ok := tt.address[tt.inputDataCity][tt.inputDataStreet]; ok == tt.returnError { // если не ждём ошибку значит должно быть записано. Если ждём == пусто.
					t.Fatalf("ожидалось %s получено иное значение", tt.inputDataStreet)
				}
			}

			if tt.inputDataHouse != "" {
				if _, ok := tt.address[tt.inputDataCity][tt.inputDataStreet][tt.inputDataHouse]; ok == tt.returnError { // если не ждём ошибку значит должно быть записано. Если ждём == пусто.
					t.Errorf("ожидалось %s получено иное значение", tt.inputDataHouse)
				}
			}

		})
	}
}

func TestRemoveAddress(t *testing.T) {
	tests := []struct {
		name            string
		address         map[string]map[string]map[string]struct{}
		inputDataCity   string
		inputDataStreet string
		inputDataHouse  string
		returnError     bool
	}{
		{"полный, валидный адрес: город+улица+дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Липецк", "Циолковского", "14", false},
		{"неполный, валидный адрес: город+улица", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Липецк", "Циолковского", "", false},
		{"неполный, валидный адрес: город", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Липецк", "", "", false},
		{"невалидный адрес: улица+дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "", "Циолковского", "14", true},
		{"невалидный адрес: только дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "", "", "14", true},
		{"невалидный адрес: только улица", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "", "Циолковского", "", true},
		{"валидный адрес: отсутствуют в map удаляемые город+улица+дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Воронеж", "Верхняя", "4", true},
		{"валидный адрес: отсутствуют в map удаляемые улица+дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Липецк", "Верхняя", "4", true},
		{"валидный адрес: отсутствуют в map удаляемый дом", map[string]map[string]map[string]struct{}{"Липецк": {"Циолковского": {"14": {}}}}, "Липецк", "Циолковского", "4", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := removeAddress(tt.address, tt.inputDataCity, tt.inputDataStreet, tt.inputDataHouse)

			if (err != nil) != tt.returnError {
				t.Fatalf("ожидалось %v получено %v", tt.returnError, err)
			}

			if tt.inputDataStreet == "" && tt.inputDataHouse == "" && !tt.returnError {
				if _, ok := tt.address[tt.inputDataCity]; ok {
					t.Fatalf("ошибка удаления города %s", tt.inputDataCity)
				}
			}

			if tt.inputDataStreet != "" && tt.inputDataHouse == "" && !tt.returnError {
				if _, ok := tt.address[tt.inputDataCity][tt.inputDataStreet]; ok {
					t.Fatalf("ошибка удаления улицы %s", tt.inputDataStreet)
				}

				if _, ok := tt.address[tt.inputDataCity]; !ok {
					t.Fatalf("город %s отсутствует", tt.inputDataCity)
				}
			}

			if tt.inputDataHouse != "" && !tt.returnError {
				if _, ok := tt.address[tt.inputDataCity][tt.inputDataStreet][tt.inputDataHouse]; ok {
					t.Fatalf("ошибка удаления дома %s", tt.inputDataHouse)
				}

				if _, ok := tt.address[tt.inputDataCity][tt.inputDataStreet]; !ok {
					t.Fatalf("улица %s отсутствует", tt.inputDataStreet)
				}
			}

		})
	}

}
