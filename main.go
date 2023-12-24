package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var romanArr = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func Decode(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}

var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func decimalToRomanRecursive(num int) string {
	if num == 0 {
		return ""
	}
	for _, pair := range romanMap {
		if num >= pair.decVal {
			return pair.symbol + decimalToRomanRecursive(num-pair.decVal)
		}
	}
	return ""
}

func readOperation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)
	return operation
}

func splitString(str string) []string {
	sliceValues := strings.Split(str, " ")
	return sliceValues
}

func catchValues(s []string) (string, string, string) {
	if len(s) > 3 {
		panic(errors.New("Формат математической операции не удовлетворяет заданию"))
	}
	a, b, c := s[0], s[1], s[2]
	return a, b, c
}

func catchOperation(a string, operator string, b string) (int, error) {
	valueOne, err := strconv.Atoi(a)
	if err != nil {
		panic(errors.New("Это не число!"))
	}
	valueTwo, err := strconv.Atoi(b)
	if err != nil {
		panic(errors.New("Это не число!"))
	}
	var result int
	switch operator {
	case "+":
		result = valueOne + valueTwo
	case "-":
		result = valueOne - valueTwo
	case "*":
		result = valueOne * valueTwo
	case "/":
		result = valueOne / valueTwo
	default:
		return 0, errors.New("Такой операции не предусмотрено.")
	}
	return result, nil
}

func checkRoman(a, b string) (bool, error) {
	result := false
	condition := 0
	for _, romInt := range romanArr {
		if a == romInt {
			condition++
		}
		if b == romInt {
			condition++
		}
	}
	if condition == 2 {
		result = true
	}
	if condition == 1 {
		return result, errors.New("Используются одновременно разные системы счисления,вы указали число больше 10 или ввели некорректное значение")
	}
	return result, nil
}

func main() {
	operation := readOperation()
	s := splitString(operation)
	a, operator, b := catchValues(s)
	isRoman, err := checkRoman(a, b)
	if err != nil {
		panic(err)
	}
	if isRoman == true {
		aInt := Decode(a)
		aStr := strconv.Itoa(aInt)
		bInt := Decode(b)
		bStr := strconv.Itoa(bInt)
		result, err := catchOperation(aStr, operator, bStr)
		if err != nil {
			panic(err)
		}
		if result <= 0 {
			panic(errors.New("В римской системе нет отрицательных чисел"))
		}
		resultInput := decimalToRomanRecursive(result)
		fmt.Println(resultInput)
	} else {
		aStr, _ := strconv.Atoi(a)
		bStr, _ := strconv.Atoi(a)
		if aStr > 10 || bStr > 10 {
			panic(errors.New("Вы указали число больше 10"))
		}
		result, err := catchOperation(a, operator, b)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
