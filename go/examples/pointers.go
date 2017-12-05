package main

import "fmt"

// * is a pointer
// *x deferences the pointer x
// &y returns *y (a pointer)

// x is a copy of the argument passed by the callsite
// so modifying x won't affect anything outside of this functions scope
func foo(x int) {
	x = 10
	fmt.Println(x)
}

// *type is a pointer
func bar(x *int) {
	*x = 10
}

func main() {
	x := 20
	foo(x)
	fmt.Println(x)

	bar(&x)
	fmt.Println(x)
}
