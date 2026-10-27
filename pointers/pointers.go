package main

import "fmt"

func main() {
	// pointersTestFunc()
	// var intPointer *int
	// var intNilPointer *int
	// intVar := 10
	// intPointer = &intVar
	// printIfNotNil(intNilPointer)
	// printIfNotNil(intPointer)
	breastSize := 1
	sizePointer := &breastSize
	changeSize(3, sizePointer)
}

func pointersTestFunc() {
	intVar := 10
	intPointer := &intVar
	fmt.Printf("Адрес - %v, значение - %d\n", intPointer, *intPointer)
	editIntPointer(14, intPointer)

	stringVar := "Test"
	strPointer := &stringVar
	fmt.Printf("Адрес - %v, значение - %v\n", strPointer, *strPointer)
	editStrPointer("New test", strPointer)

	floatVar := 3.14
	floatPointer := &floatVar
	fmt.Printf("Адрес - %v, значение - %v\n", floatPointer, *floatPointer)
	editFloatPointer(6.66, floatPointer)

	boolVar := true
	boolPointer := &boolVar
	fmt.Printf("Адрес - %v, значение - %v\n", boolPointer, *boolPointer)
	editBoolPointer(false, boolPointer)
}

func editIntPointer(newValue int, intPointer *int) {
	fmt.Printf("Old value - %d\n", *intPointer)
	*intPointer = newValue
	fmt.Printf("New value - %d\n", *intPointer)
}

func editStrPointer(newValue string, strPointer *string) {
	fmt.Printf("Old value - %v\n", *strPointer)
	*strPointer = newValue
	fmt.Printf("New value - %v\n", *strPointer)
}

func editFloatPointer(newValue float64, floatPointer *float64) {
	fmt.Printf("Old value - %v\n", *floatPointer)
	*floatPointer = newValue
	fmt.Printf("New value - %v\n", *floatPointer)
}

func editBoolPointer(newValue bool, boolPointer *bool) {
	fmt.Printf("Old value - %v\n", *boolPointer)
	*boolPointer = newValue
	fmt.Printf("New value - %v\n", *boolPointer)
}

func printIfNotNil(pointer *int) {
	if pointer == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println(*pointer)
	}
}

func changeSize(newSize int, sizeVarPointer *int) {
	if sizeVarPointer == nil {
		fmt.Println("Pointer is nil")
	} else {
		*sizeVarPointer = newSize
		fmt.Printf("New size = %d\n", *sizeVarPointer)
	}
}
