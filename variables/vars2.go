package main

import "fmt"

func VarIntegerOperations() {
	integer := 10

	fmt.Printf("integer = %v\n", integer)

	fmt.Println("+1")
	integer = integer + 1
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("-1")
	integer = integer - 1
	fmt.Printf("inintegert = %v\n", integer)

	fmt.Println("*2")
	integer = integer * 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("/2")
	integer = integer / 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("%2")
	integer = integer % 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("+= 3")
	integer += 3
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("-= 2")
	integer -= 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("*= 2")
	integer *= 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("/= 2")
	integer /= 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("%= 2")
	integer %= 2
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("++")
	integer++
	fmt.Printf("integer = %v\n", integer)

	fmt.Println("--")
	integer--
	fmt.Printf("integer = %v\n", integer)
}

func VarStringOperations() {
	string := "Hello"

	fmt.Printf("string = %q\n", string)

	fmt.Println("+ world!")
	string = string + "world"
	fmt.Printf("int = %q\n", string)

	fmt.Println("+= 3")
	string += "!"
	fmt.Printf("int = %q\n", string)
}

func VarFloatOperations() {
	float := 3.14

	fmt.Printf("float = %v\n", float)

	fmt.Println("+1")
	float = float + 1
	fmt.Printf("float = %v\n", float)

	fmt.Println("-1")
	float = float - 1
	fmt.Printf("float = %v\n", float)

	fmt.Println("*2")
	float = float * 2
	fmt.Printf("float = %v\n", float)

	fmt.Println("/2")
	float = float / 2
	fmt.Printf("float = %v\n", float)

	fmt.Println("+= 3")
	float += 3
	fmt.Printf("float = %v\n", float)

	fmt.Println("-= 2")
	float -= 2
	fmt.Printf("float = %v\n", float)

	fmt.Println("*= 2")
	float *= 2
	fmt.Printf("float = %v\n", float)

	fmt.Println("/= 2")
	float /= 2
	fmt.Printf("float = %v\n", float)

	fmt.Println("++")
	float++
	fmt.Printf("float = %v\n", float)

	fmt.Println("--")
	float--
	fmt.Printf("float = %v\n", float)
}

func VarBooleanOperation() {
	boolean := true

	fmt.Printf("boolean = %v\n", boolean)

	fmt.Println("+1")
	boolean = !boolean
	fmt.Printf("boolean = %v\n", boolean)
}
