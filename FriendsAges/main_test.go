package main

import "testing"

func TestGetAge(t *testing.T) {
	tests := []struct {
		name        string
		ages        map[string]int
		person      string
		expectedAge int
		expectedOk  bool
	}{
		{"Искомый есть в карте", map[string]int{"Pasha": 45}, "Pasha", 45, true},
		{"Искомого нет в карте", map[string]int{"Pasha": 45}, "Sasha", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ok := getAge(tt.ages, tt.person)
			if ok != tt.expectedOk {
				t.Fatalf("Ожидалось %v, получено %v\n", tt.expectedOk, ok)
			}

			if ok && res != tt.expectedAge {
				t.Errorf("Ожидалось %v, получено %v\n", tt.expectedAge, res)
			}

		})
	}
}

func TestRemovePerson(t *testing.T) {
	tests := []struct {
		name           string
		initialAges    map[string]int
		personToRemove string
		expectedOk     bool
	}{
		{"Удаление существующего", map[string]int{"Pasha": 45, "Sasha": 25, "Masha": 35}, "Sasha", false},
		{"Удаление не существующего", map[string]int{"Pasha": 45, "Masha": 35}, "Sasha", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removePerson(tt.initialAges, tt.personToRemove)
			_, ok := getAge(tt.initialAges, tt.personToRemove)

			if ok != tt.expectedOk {
				t.Errorf("Ожидалось %v, получено %v\n", tt.expectedOk, ok)
			}

		})
	}
}
