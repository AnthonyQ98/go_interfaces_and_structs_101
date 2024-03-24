package main

import (
	"fmt"
	"golang_interfaces/sounds"
)

func main() {

	// Test interface
	dog := sounds.NewDog("")
	sounds.CallSound(dog)
	phone := sounds.NewPhone("")
	sounds.CallSound(phone)

	// Test getters
	fmt.Println(phone.Owner())
	fmt.Println(phone.Model())

	fmt.Println(dog.Breed())
	fmt.Println(dog.Owner())

	// Test setters
	dog.SetBreed("Whippet")
	dog.SetOwner("Eamon")

	phone.SetOwner("Jim")
	phone.SetModel("Redmi Note 10")

	// Verify setters
	fmt.Println(phone.Owner())
	fmt.Println(phone.Model())

	fmt.Println(dog.Breed())
	fmt.Println(dog.Owner())
}
