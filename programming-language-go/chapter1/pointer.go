package main

import "fmt"

func IncrementV(v int) {
	v++
}

func IncrementVByPointer(v *int) {
	(*v)++
}

func main() {
	v := 1

	IncrementV(v)
	fmt.Println(v) // 1

	IncrementVByPointer(&v)
	fmt.Println(v) // 2
}
