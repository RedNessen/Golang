package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

var VegMap = make(map[string]int, 10)

func main() {

	for {
		scanner := bufio.NewScanner(os.Stdin)

		pp.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			pp.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()
		fields := strings.Fields(text)

		if len(fields) == 0 {
			pp.Println("Вы ничего не ввели")
			return
		}

		if len(fields) == 3 && fields[0] == "добавить" {
			n, err := strconv.Atoi(fields[2])
			if err != nil {
				pp.Println("Ошибка, дано не число")
			} else {
				addVeg(fields[1], n)
			}
		}

		if len(fields) == 3 && fields[0] == "удалить" {
			n, err := strconv.Atoi(fields[2])
			if err != nil {
				pp.Println("Ошибка, дано не число")
			} else {
				decreaseVeg(fields[1], n)
			}
		}

		if len(fields) == 2 && fields[0] == "получить" {
			getVeg(fields[1])
		}

		if fields[0] == "help" {
			help()
		}

		if fields[0] == "выйти" {
			return
		}
	}

}

func addVeg(key string, increase int) {
	VegMap[key] += increase
}

func decreaseVeg(key string, decrease int) {
	value, ok := VegMap[key]
	if ok {
		if value-decrease < 0 {
			pp.Println("Итоговое значение меньше 0 - ошибка")
		} else {
			VegMap[key] -= decrease
		}
	} else {
		pp.Println("Такого овоща нету")
		return
	}
}

func getVeg(key string) {
	_, ok := VegMap[key]
	if ok {
		pp.Println(VegMap[key])
	} else {
		pp.Println("Такого овоща нету")
		return
	}
}

func help() {
	pp.Println("Список команд")
	pp.Println("добавить {ингредиент} {количество}")
	pp.Println("-- эта команда должна добавлять какой-то ингредиент, и запоминать сколько и чего мы добавили")
	pp.Println("удалить {ингредиент} {количество}")
	pp.Println("-- эта команда должна удалять какой-то ингредиент, если он был до этого добавлен (если идёт попытка удалить несуществующий ингредиент, или попытка уйти в минус, то писать на экран об этом уведомление)")
	pp.Println("получить {ингредиент}")
	pp.Println("-- эта команда должна выводить на экран итоговое количество ингредиента (например если было добавлено 10 а потом удалено 4 то в итоге осталось 6 какого-то ингредиента, нужно вывести это число)")
	pp.Println("help")
	pp.Println("-- эта команда должна выводить список доступных команд и их описание")
	pp.Println("выйти")
	pp.Println("-- эта команда должна выводить список доступных команд и их описание")
}
