package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type Employee struct {
	Person
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}

func main() {

	listOfEmployees := []Employee{}

	list := []struct {
		Name     string
		Age      int
		Position string
		Salary   float64
	}{
		{"Паша", 45, "Оператор ЧПУ", 522},
		{"Никита", 0, "Наладчик", 700},
		{"Сероб", 0, "", 522},
	}

	for _, tt := range list {
		employee, err := createEmployee(tt.Name, tt.Age, tt.Position, tt.Salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		listOfEmployees, err = addEmployeeInList(listOfEmployees, employee)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	result, err := preparationForPreservation(listOfEmployees)

	if err != nil {
		fmt.Println(err)
		return
	}

	listOfEmployees = []Employee{} // процедура записи результата в файл. пока я этого не умею, по этому такое действие.
	// fmt.Printf("%v\n%s\n", result, string(result)) // просто интересно было посмотреть как это выглядит.
	// процедура загрузки результата из файла. Пока я этого не умею

	err1 := dataExtraction(result, &listOfEmployees)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	printListEmployees(listOfEmployees)
}

func createEmployee(name string, age int, position string, salary float64) (Employee, error) {
	if name == "" {
		return Employee{}, errors.New("Не указан работник")
	}
	e := Employee{
		Person: Person{
			Name: name,
			Age:  age},
		Position: position,
		Salary:   salary,
	}

	return e, nil
}

func addEmployeeInList(list []Employee, emp Employee) ([]Employee, error) {
	for _, empList := range list {
		if (empList.Name == emp.Name) && (empList.Age == emp.Age) {
			return list, fmt.Errorf("сотрудник %s %d лет уже есть в списке", empList.Name, empList.Age)
		}
	}
	return append(list, emp), nil

}

func preparationForPreservation(list []Employee) ([]byte, error) {
	l, err := json.Marshal(list)
	return l, err
}

func dataExtraction(result []byte, list *[]Employee) error {

	err := json.Unmarshal(result, list)
	return err
}

func printListEmployees(list []Employee) {
	for num, emp := range list {
		fmt.Printf("Сотрудник №%d.\n\tИмя: %s. Возраст: %d.\n\t\tДолжность: %s. З/п в час: %.2f\n", num+1, emp.Name, emp.Age, emp.Position, emp.Salary)
	}
}
