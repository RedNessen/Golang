package main

import (
	"fmt"
	"math"
)

var GlobalTemp int

func main() {
	printingTextX(1)
	greetingGuests("Динамо", 42)
	squareGuess(1, -5, 6)
	squareGuess(1, -4, 4)
	squareGuess(1, 2, 5)

	sqrExp(2)

	increaseGlobalTemp(11)
	decreaseGlobalTemp(12)
}

func printingTextX(count int) {
	for range count {
		fmt.Println("Отель Баронесса - лучший отель в Нью-Йорке")
	}
}

func greetingGuests(name string, roomNumber int) {
	fmt.Printf("Добро пожаловать в Баронессу, уважаемый %s, позвольте проводить вас в комнату номер %d\n", name, roomNumber)
}

func squareGuess(a float64, b float64, c float64) {
	discriminant := (b * b) - (4 * a * c)

	if discriminant > 0 {
		result1 := (-b + (math.Sqrt(discriminant))) / (2 * a)
		result2 := (-b - (math.Sqrt(discriminant))) / (2 * a)
		fmt.Printf("%v %v\n", result1, result2)
	} else if discriminant == 0 {
		result := -b / (2 * a)
		fmt.Println(result)
	} else {
		fmt.Println("D < 0 - действительных корней нет")
	}
}

func sqrExp(x int) {
	result := x * x
	fmt.Println(result)
}

func increaseGlobalTemp(increment int) {
	GlobalTemp += increment
	fmt.Printf("Температура увеличилась и составляет %d\n", GlobalTemp)
}

func decreaseGlobalTemp(decrement int) {
	GlobalTemp -= decrement
	fmt.Printf("Температура уменьшилась и составляет %d\n", GlobalTemp)
}
