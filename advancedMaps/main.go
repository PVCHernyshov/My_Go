package main

import (
	"errors"
	"fmt"
)

func main() {

	address := map[string]map[string]map[string]struct{}{}

	inputData := []struct {
		city   string
		street string
		home   string
	}{
		{"Липецк", "Циолковского", "25"},
		{"Липецк", "Смородина", "25"},
		{"Липецк", "Смородина", "2"},
		{"Москва", "Центральная", "12"},
		{"Москва", "Угловая", "12"},
		{"Санкт Петербург", "Невский проспект", "12"},
		{"Грязи", "", ""},
		{"", "", "15"},
		{"", "Петровская", ""},
		{"Воронеж", "", "12"},
	}

	for _, addr := range inputData {
		err := addAddress(address, addr.city, addr.street, addr.home)
		if err != nil {
			fmt.Println(err)
		}
	}

	printAddresses(address)

	dataToBeDeleted := []struct {
		city   string
		street string
		home   string
	}{
		{"Санкт Петербург", "", ""},
		{"Липецк", "Циолковского", ""},
		{"Липецк", "Смородина", "2"},
		{"Москва", "Угловая", "12"},
		{"Воронеж", "центральная", "2"},
		{"Москва", "Окружная", "2"},
		{"Москва", "Угловая", "10"},
	}

	for _, delAddr := range dataToBeDeleted {
		err := removeAddress(address, delAddr.city, delAddr.street, delAddr.home)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("\n\n\n")
	printAddresses(address)

}

func addAddress(addresses map[string]map[string]map[string]struct{}, city, street, house string) error {

	err := addressValidation(city, street, house)
	if err != nil {
		return err
	}

	if _, ok := addresses[city]; !ok { //в новом городе
		addresses[city] = map[string]map[string]struct{}{}
	}

	if street != "" {
		if _, ok := addresses[city][street]; !ok {
			addresses[city][street] = map[string]struct{}{}
		}
	}

	if house != "" {
		addresses[city][street][house] = struct{}{}
	}

	return nil
}

func removeAddress(addresses map[string]map[string]map[string]struct{}, city, street, house string) error {

	err := addressValidation(city, street, house)
	if err != nil {
		return err
	}

	if _, ok := addresses[city]; !ok {
		return fmt.Errorf("Запрашиваемый на удаление город %s не найден", city)
	}

	if house == "" && street == "" {
		delete(addresses, city)
		return nil
	}

	if _, ok := addresses[city][street]; !ok {
		return fmt.Errorf("Запрашиваемая на удаление улица %s в городе %s не найдена", street, city)
	}

	if house == "" && street != "" {
		delete(addresses[city], street)
		return nil
	}

	if _, ok := addresses[city][street][house]; !ok {
		return fmt.Errorf("Запрашиваемый на удаление дом %s на улице %s в городе %s не найден", house, street, city)
	}

	delete(addresses[city][street], house)
	return nil
}

func printAddresses(addresses map[string]map[string]map[string]struct{}) {
	for cityName, street := range addresses {
		fmt.Println(cityName, ":")
		for streetName, house := range street {
			fmt.Print("\t", streetName, ": ")
			for houseNumber := range house {
				fmt.Print(houseNumber, "; ")
			}
			fmt.Println()
		}
	}
}

func addressValidation(city, street, house string) error {
	if city == "" || (street == "" && house != "") {
		if city == "" {
			return errors.New("получены невалидные данные: отсутствует название города")
		}
		return errors.New("получены невалидные данные: отсутствует название улицы")
	}
	return nil
}
