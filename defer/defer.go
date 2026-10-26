package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("Program shut down")
	}()
	defer func() {
		fmt.Println("Program finished")
	}()

	fmt.Println("Start working. Pending DB connection")
	dbConnection()
	fmt.Println("Program working")

}

func dbConnection() {
	defer func() {
		fmt.Println("Connection stopped")
	}()
	fmt.Println("Connection started")
	fmt.Println("Connection working")
}
