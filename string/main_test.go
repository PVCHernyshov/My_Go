package main

import "testing"

func TestSubstring(t *testing.T) {

	test := []struct {
		name string
		s, e int
		res  string
		err  bool
	}{
		{"Старт=[2] Конец=[8]", 2, 8, "llo wo", false},
		{"Старт=[0] Конец=[1]", 0, 1, "H", false},
		{"Старт=[11] Конец=[12]", 11, 12, "!", false},
		{"Старт=[11] Конец=[13]", 11, 13, "", true},
		{"Старт=[-1] Конец=[8]", -1, 8, "", true},
		{"Старт=[5] Конец=[2]", 5, 2, "", true},
	}

	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {
			result, err := substring("Hello world!", tt.s, tt.e)
			if (err != nil) != tt.err {
				t.Fatalf("Ожидалось %v получено %v", tt.err, err)
			}
			if !tt.err && result != tt.res {
				t.Errorf("Ожидалось %s получено %s", tt.res, result)
			}
		})

	}

}
