package main

import "fmt"

// type <name> struct { ...fields }
// Now Dog can be used as a type
type Dog struct {
	name string
	age  int
}

// Methods are basically instance functions
// func <receiver> <name>(<arguments>) <returns> {}
func (dog *Dog) growOlder() {
	dog.age += 1
}

// Inheritance
// ShibaInu has Dog
type ShibaInu struct {
	Dog
}

func (shiba ShibaInu) bark() {
	fmt.Println(shiba.Dog.name, "barks")
	shiba.Dog.name = "yolo"
}

// Interface defines methods instead of variables a type should have
// You can pass DogsBehavior into arguments
type DogBehavior interface {
	bark()
}

func printDog(dog Dog) {
	fmt.Println(dog)
}

func changeDogAge(dog *Dog) {
	dog.age += 1
}

func dogBark(dog DogBehavior) {
	dog.bark()
}

func main() {
	// initialize type Dog
	dog := Dog{
		name: "Fido",
		age:  10,
	}

	fmt.Println(dog.name)
	printDog(dog)

	changeDogAge(&dog)
	fmt.Println(dog)

	dog.growOlder()
	fmt.Println(dog)

	shiba := ShibaInu{
		Dog{name: "Shiba", age: 20},
	}
	shiba.growOlder()
	fmt.Println(shiba)

	dogBark(shiba)

	fmt.Println(shiba)
}
