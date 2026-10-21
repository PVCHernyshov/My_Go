package main

import (
	"reflect"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	tests := []struct {
		NameTest       string
		NameEmplo      string
		Age            int
		Position       string
		Salary         float64
		expectedStruct Employee
		expectedError  bool
	}{
		{"все данные", "Паша", 45, "Оператор", 522, Employee{Person: Person{Name: "Паша", Age: 45}, Position: "Оператор", Salary: 522}, false},
		{"без возраста", "Паша", 0, "Оператор", 522, Employee{Person: Person{Name: "Паша", Age: 0}, Position: "Оператор", Salary: 522}, false},
		{"без возраста и должности", "Паша", 0, "", 522, Employee{Person: Person{Name: "Паша", Age: 0}, Position: "", Salary: 522}, false},
		{"без возраста, должности, з/п ставки", "Паша", 0, "", 0, Employee{Person: Person{Name: "Паша", Age: 0}, Position: "", Salary: 0}, false},
		{"без имени", "", 45, "Оператор", 522, Employee{}, true},
	}

	for _, tt := range tests {

		t.Run(tt.NameTest, func(t *testing.T) {
			result, err := createEmployee(tt.NameEmplo, tt.Age, tt.Position, tt.Salary)

			if (err != nil) != tt.expectedError {
				t.Fatalf("ожидалось %v получено %v", tt.expectedError, err)
			}

			if result != tt.expectedStruct {

				if tt.expectedStruct.Name != result.Name {
					t.Errorf("ожидалось %s получено %s ", tt.expectedStruct.Name, result.Name)
				}
				if tt.expectedStruct.Age != result.Age {
					t.Errorf("ожидалось %d получено %d ", tt.expectedStruct.Age, result.Age)
				}
				if tt.expectedStruct.Position != result.Position {
					t.Errorf("ожидалось %s получено %s ", tt.expectedStruct.Position, result.Position)
				}
				if tt.expectedStruct.Salary != result.Salary {
					t.Errorf("ожидалось %v получено %v ", tt.expectedStruct.Salary, result.Salary)
				}
			}

		})
	}
}

func TestAddEmployeeInList(t *testing.T) {

	pasha := Employee{Person: Person{Name: "Паша", Age: 45}, Position: "Оператор", Salary: 522}
	sasha := Employee{Person: Person{Name: "Cаша", Age: 40}, Position: "Наладчик", Salary: 600}
	pashaName := Employee{Person: Person{Name: "Паша", Age: 40}, Position: "Оператор", Salary: 522}
	sashaAge := Employee{Person: Person{Name: "Cаша", Age: 45}, Position: "Оператор", Salary: 522}

	tests := []struct {
		Name          string
		inputSlice    []Employee
		inputStruct   Employee
		expectedSlice []Employee
		expectedErr   bool
	}{
		{"в пустой срез новые данные", []Employee{}, pasha, []Employee{pasha}, false},
		{"в не пустой срез новые данные", []Employee{pasha}, sasha, []Employee{pasha, sasha}, false},
		{"в не пустой срез повторные данные", []Employee{pasha}, pasha, []Employee{pasha}, true},
		{"в не пустой срез совпадает только имя", []Employee{pasha}, pashaName, []Employee{pasha, pashaName}, false},
		{"в не пустой срез совпадает только возраст", []Employee{pasha}, sashaAge, []Employee{pasha, sashaAge}, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := addEmployeeInList(tt.inputSlice, tt.inputStruct)

			if (err != nil) != tt.expectedErr {
				t.Fatalf("ожидалось %v получено %v", tt.expectedErr, err)
			}
			if !reflect.DeepEqual(result, tt.expectedSlice) {
				t.Error("ошибка сравнения ожидаемого и полученного срезов")
			}
		})
	}

}
