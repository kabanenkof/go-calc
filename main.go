package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabicNumbers = map[string]int{
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
	"6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
}

var romanNumbers = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func toRoman(number int) string {
	roman := ""
	for number > 0 {
		if number >= 10 {
			roman += "X"
			number -= 10
		} else if number >= 9 {
			roman += "IX"
			number -= 9
		} else if number >= 5 {
			roman += "V"
			number -= 5
		} else if number >= 4 {
			roman += "IV"
			number -= 4
		} else {
			roman += "I"
			number -= 1
		}
	}
	return roman
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	splitInput := strings.Split(input, " ")

	if len(splitInput) != 3 {
		panic("Ошибка: недопустимое выражение")
		return
	}

	firstNumber, err := strconv.Atoi(splitInput[0])
	if err != nil {
		panic("Ошибка: недопустимое число")
		return
	}

	secondNumber, err := strconv.Atoi(splitInput[2])
	if err != nil {
		panic("Ошибка: недопустимое число")
		return
	}

	if firstNumber < 1 || firstNumber > 10 || secondNumber < 1 || secondNumber > 10 {
		panic("Ошибка: числа должны быть от 1 до 10")
		return
	}

	operator := splitInput[1]

	if _, ok := arabicNumbers[splitInput[0]]; ok {
		if _, ok := romanNumbers[splitInput[2]]; ok {
			panic("Ошибка: числа должны быть либо только арабскими, либо только римскими")
			return
		} else {
			result := calculate(firstNumber, secondNumber, operator)
			fmt.Println("Результат:", result)
		}
	} else if _, ok := romanNumbers[splitInput[0]]; ok {
		if _, ok := arabicNumbers[splitInput[2]]; ok {
			panic("Ошибка: числа должны быть либо только арабскими, либо только римскими")
			return
		} else {
			result := calculate(firstNumber, secondNumber, operator)
			if result < 0 {
				panic("Ошибка: римские числа не могут быть отрицательными")
				return
			}
			fmt.Println("Результат:", toRoman(result))
		}
	} else {
		panic("Ошибка: недопустимое выражение")
		return
	}
}
