package main

import "github.com/k0kubun/pp"

type Dog struct {
	name        string
	judgeScore  int
	isDogPolite bool
}

func NewDog(
	name string,
	judgeScore int,
	isDogPolite bool,
) Dog {
	return Dog{
		name:        name,
		judgeScore:  judgeScore,
		isDogPolite: isDogPolite,
	}
}

func main() {

	// arr := [7]int{1, 2, 3, 4, 5, 6}
	// sliceExample := []int{}
	// mapExample := map[int]int{}

	// sliceExample2 := make([]int, 0, 10)
	// holidaysMap := make(map[int]bool, 31)

	// arr[6] = 7
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i]%2 == 1 {
	// 		arr[i] *= 2
	// 	}
	// }

	// sliceExample2 = append(sliceExample2, 13)
	// sliceExample2 = append(sliceExample2, 22)
	// sliceExample2 = append(sliceExample2, 131)
	// sliceExample2 = append(sliceExample2, 2)
	// sliceExample2 = append(sliceExample2, 111)
	// sliceExample2 = append(sliceExample2, 7)

	// for i := range sliceExample2 {
	// 	pp.Println(sliceExample2[i])
	// }
	// for i, v := range arr {
	// 	pp.Println(i, v)
	// }

	// holidaysMap[26] = true
	// holidaysMap[23] = false
	// holidaysMap[27] = true
	// holidaysMap[24] = false
	// delete(holidaysMap, 24)

	// for k, v := range holidaysMap {
	// 	pp.Println(k, v)
	// }

	// v, ok := holidaysMap[1]
	// pp.Println(v, ok)

	dog1 := NewDog("Lacy", 7, true)
	dog2 := NewDog("Jack", 6, false)
	dog3 := NewDog("Monty", 8, true)

	dogSlice := []Dog{dog1, dog2, dog3}

	pp.Println(dogSlice)

	for i, v := range dogSlice {
		if v.isDogPolite {
			dogSlice[i].judgeScore += 1
		}
	}

	pp.Println(dogSlice)

}
