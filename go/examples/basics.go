// Declare what package this file belongs to
package main

import "fmt"

// - Variables
// - For Loop
// - If
// - Switch
// - Array
// - Slice
// - make
// - Copy, Append
// - Map

func main() {
	/********************************
	* Variables
	*******************************/

	// var <name> <type> = <type value>
	var x string = "Hello"

	// shorthand for x = x + <type x value>
	x += " World"

	// check equality
	// ==
	fmt.Println(x == x)

	// initialize by inferring type
	y := "Hello"
	fmt.Println(y)
	// y is type String

	// Constants
	const cannotBeChanged = "Cartoon Character"
	fmt.Println(cannotBeChanged)

	// Defining multiple variables
	var (
		a = 1
		b = 2
	)
	const (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	/********************************
	* Control Flow
	*******************************/

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
	}

	// for loop as a while loop
	var cancel bool = false
	for cancel == false {
		fmt.Println("don't cancel me")
		cancel = true
	}

	// if
	if 1 == 1 {
		fmt.Println("wow")
	} else {
		fmt.Println("even more wow")
	}

	// switch
	someType := "dog"
	switch someType {
	case "dog":
		fmt.Println("its a dog")
	case "cat":
		fmt.Println("its a cat")
	default:
		fmt.Println("its a nothing")
	}

	/********************************
	* Arrays
	*******************************/

	var arr [5]int
	for i:=0; i<len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr, len(arr), cap(arr))

	// initialize inferring type
	arr2 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(arr2, len(arr2), cap(arr2))

	// A slice is a segment of an array, length is allowed to change
	var slice [3]int
	for i:=0; i<3; i++ {
		slice[i] = i
	}
	fmt.Println(slice, len(slice), cap(slice))

	// use make to initialize a slice
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

	arr3 := [10]int{1,2,3,4,5,6,7,8,9,10}
	sliceOfArr3 := arr3[1:4]
	fmt.Println(arr3, sliceOfArr3)

	sliceRestOfArr3 := arr3[5:]
	fmt.Println(sliceRestOfArr3)

	appendedSlice := append(sliceRestOfArr3, 200, 300)
	fmt.Println(appendedSlice)

	// copy slice
	copySlice1 := []int{1,2,3}
	copySlice2 := []int{100}
	copy(copySlice1, copySlice2)
	fmt.Println(copySlice1, copySlice2)

	/********************************
	* Maps
	*******************************/

	// map[<key type>] <value type>
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	fmt.Println(map1, len(map1))
	delete(map1, "one")
	fmt.Println(map1, len(map1))

	// check for existence of key value
	value, ok := map1["three"]
	fmt.Println(value, ok)

	if value, ok := map1["three"]; ok {
		fmt.Println(value)
	}
}
