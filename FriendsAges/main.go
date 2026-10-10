package main

import "fmt"

func main() {

}
func addPerson(ages map[string]int, name string, age int) {
	ages[name] = age
}

func getAge(ages map[string]int, name string) (int, bool) {
	age, exist := ages[name]
	return age, exist
}
func removePerson(ages map[string]int, name string) {
	delete(ages, name)
}

func printAge(ages map[string]int, name string) {
	age, ok := getAge(ages, name)
	if ok {
		fmt.Printf("Возраст %s %d лет\n", name, age)
	} else {
		fmt.Printf("Данные для %s не найдены.\n", name)
	}
}
