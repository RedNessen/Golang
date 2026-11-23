package main

import "fmt"

type Aboba struct {
	Name     string
	Location string
	Age      int
}

func NewAboba(
	name string,
	location string,
	age int,
) Aboba {
	if name == "" {
		return Aboba{}
	}

	if location == "" {
		return Aboba{}
	}

	if age <= 0 {
		return Aboba{}
	}

	return Aboba{
		Name:     name,
		Location: location,
		Age:      age,
	}
}

func (a *Aboba) ChangeName(newName string) {
	if newName != "" {
		a.Name = newName
	} else {
		fmt.Println("Error. Empty variable")
	}
}

func (a *Aboba) ChangeLocation(newLocation string) {
	if newLocation != "" {
		a.Location = newLocation
	} else {
		fmt.Println("Error. Empty variable")
	}
}

func (a *Aboba) AgeChange(ageChange int) {
	if (a.Age + ageChange) > 0 {
		a.Age += ageChange
	} else {
		fmt.Println("Error. Wrong result age")
	}
}

type Appartment struct {
	address     string
	squareSpace float64
	floorNumber int
	value       int
}

func NewAppartment(
	address string,
	squareSpace float64,
	floorNumber int,
	value int,
) Appartment {
	if address == "" {
		return Appartment{}
	}

	if squareSpace <= 10.0 {
		return Appartment{}
	}

	if floorNumber <= 0 || floorNumber > 100 {
		return Appartment{}
	}

	if value <= 0 {
		return Appartment{}
	}

	return Appartment{
		address:     address,
		squareSpace: squareSpace,
		floorNumber: floorNumber,
		value:       value,
	}
}

func (a *Appartment) ChangeValue(changeValue int) {
	if a.value+changeValue <= 0 {
		fmt.Println("Error. Wrong result after summary")
	} else {
		a.value += changeValue
	}
}

func main() {
	abobaMaster := NewAboba("Anastasiy", "Kaliningrad", 29)
	fmt.Println(abobaMaster)
	abobaMaster.ChangeName("Julia")
	abobaMaster.ChangeLocation("")
	abobaMaster.AgeChange(-3)
	fmt.Println(abobaMaster)

	myAppartment := NewAppartment("Some address 1", 11.0, 4, 10000)
	fmt.Println(myAppartment)
	myAppartment.ChangeValue(5000)
	fmt.Println(myAppartment)

}
