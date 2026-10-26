package main

import "fmt"

func main() {
	ageChecker(29)
	ageCheckerV2(14)
	beerChecker(2.14)
	croutonnsAndBeerChecker(true, false)
	weatherChecker(4)
	premiumChecker(false)
}

func ageChecker(age int) {
	if age > 30 {
		fmt.Println("Тебе уже больше тридцатки!")
	} else {
		fmt.Println("Ты ещё молодой!")
	}
}

func ageCheckerV2(age int) {
	if age >= 18 {
		fmt.Println("Я продам тебе пиво.")
	} else if age < 12 {
		fmt.Println("Ты чего малой совсем берега попутал пиво покупать?!")
	} else {
		fmt.Println("Алкоголь только от 18-ти лет")
	}
}

func beerChecker(liters float64) {
	if liters > 3 || liters < 1 {
		fmt.Println("Ты какой-то странный...")
	} else {
		fmt.Println("А ты знаешь золотую середину!")
	}
}

func croutonnsAndBeerChecker(croutonsDelicious bool, beerDelicious bool) {
	if croutonsDelicious && beerDelicious {
		fmt.Println("Мы идем гулять!")
	} else {
		fmt.Println("Мы никуда не пойдем!")
	}
}

func weatherChecker(weather int) {
	if weather != 3 {
		fmt.Println("Сегодня я буду участвовать в гонке!")
	}
}

func premiumChecker(isPremium bool) {
	if !isPremium {
		fmt.Println("Бро, купи себе премиум")
	}
}
