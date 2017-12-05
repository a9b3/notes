package main

import "fmt"

// func <name>(<...argument types>) (<...return types>) { }
func average(arr []int) float64 {
	var sum float64
	for _, val := range arr {
		sum += float64(val)
	}
	return sum/float64(len(arr))
}

// variadic argument, can only be used on last parameter
// printThis("this will be foo", "this will be str", ...)
func printThis(foo string, strs ...string) {
	fmt.Println("foo", foo)
	fmt.Println(strs)
}

// functions can return functions
func addMore() func() {
	return func() {
		fmt.Println("hi")
	}
}

// Defer
func somethingDeferred() {
	fmt.Println("before")
	defer fmt.Println("done") // this runs at the END of the function
	fmt.Println("before")
}

// Panic/Recover
// defer a closure with recover to handle panics
func panicRecover() {
	defer func() {
		str := recover()
		fmt.Println("something wrong but its ok", str)
	}()
	panic("PANICING")
}

func main() {
	avg := average([]int{1,2,3,4})
	fmt.Println(avg)

	printThis("hi", "bye", "foo")

	bar := addMore()
	bar()

	somethingDeferred()

	panicRecover()
}
