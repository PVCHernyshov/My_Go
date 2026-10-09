package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var newGame bool = true

	for newGame {
		n := rand.IntN(51)

		fmt.Println("Угадайте число в диапазоне от 0 до 50. У вас 5 попыток.")
		var usInp int
		var flag bool = true
		att := 5
		var re rune
		var re2 rune

		for i := 1; i <= att; i++ {

			fmt.Println("Попытка №", i)

			fmt.Scan(&usInp)

			if usInp < n {
				fmt.Println("Загаданное число больше, попробуйте ещё раз.")
				continue
			}

			if usInp > n {
				fmt.Println("Загаданное число меньше, попробуйте ещё раз.")
				continue
			}

			fmt.Println("Поздравляю! Вы угадали!")
			flag = false
			break
		}
		if flag {

			fmt.Println("К сожалению, Вы исчерпали свои попытки.")
		}
		fmt.Println("Хотите попробовать ещё? д/н")
		fmt.Scanf("%c%c", &re, &re2)
		if re == 'н' || re == 'n' {
			newGame = false
			break
		}

	}
	fmt.Println("Досвидания!")

}
