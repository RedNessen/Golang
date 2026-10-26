package main

import "fmt"

func main() {
	varShort()
	varLong()
	varNil()
	VarIntegerOperations()
	VarStringOperations()
	VarFloatOperations()
	VarBooleanOperation()
}

func varShort() {
	integer := 5
	stringVar := "Hello world!"
	float := 3.14
	boolean := true

	fmt.Printf("integer = %v\n", integer)
	fmt.Printf("stringVar = %q\n", stringVar)
	fmt.Printf("float = %v\n", float)
	fmt.Printf("boolean = %v\n", boolean)
	fmt.Println()
}

func varLong() {
	var integer int = 5
	var stringVar string = "Hello world!"
	var float float64 = 3.14
	var boolean bool = true

	fmt.Printf("integer = %v\n", integer)
	fmt.Printf("stringVar = %q\n", stringVar)
	fmt.Printf("float = %v\n", float)
	fmt.Printf("boolean = %v\n", boolean)
	fmt.Println()
}

func varNil() {
	var integer int
	var stringVar string
	var float float64
	var boolean bool

	fmt.Printf("integer = %v\n", integer)
	fmt.Printf("stringVar = %q\n", stringVar)
	fmt.Printf("float = %v\n", float)
	fmt.Printf("boolean = %v\n", boolean)
	fmt.Println()
}
