package commands

import (
	"fmt"
	"strconv"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserFloatInput(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}
		fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
	}
}

func getUserIntInput(prompt string) int {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
		fmt.Println("Geçersiz giriş. Lütfen bir tam sayı girin.")
	}
}
