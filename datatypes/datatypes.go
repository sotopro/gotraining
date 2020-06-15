package main

import "fmt"

func main()  {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Daniel"
	var age int = 37
	var size float32 = 1.3
	
	const isCool = false

	// Shorthand
	lastname := "Soto"
	email := "dasg.1993@gmail.com"

	// name, email := "Daniel", "dasg.1993@gmail.com"

	fmt.Println(name, lastname, email, age, isCool, size)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)

}
