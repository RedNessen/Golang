package main

import "fmt"

func main() {
	// cycler10()
	angryBirdsV1(3)
}

func cycler10() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v. Ты повторяешься!\n", i+1)
	}
}

func angryBirdsV1(cycles int) {
	var result int
	fmt.Println("Игра началась!")
	for i := 0; i < cycles; i++ {
		fmt.Println("Птичка преодолела препятствие! +1 очко")
		result++
	}
	fmt.Printf("Игра окончена, итоговый счет = %v\n", result)
}
