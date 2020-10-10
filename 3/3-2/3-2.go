package main

import (
	"fmt"
)

//func1 is demo of for loop
func func1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}

//func2 is demo of break
func func2() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

//func3 is demo of continue
func func3() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

//func4 is demo of dead loop
func func4() {
	for {
		fmt.Println("Hello World")
	}
}

func main() {
	func1()
	func2()
	func3()
	func4()
}
