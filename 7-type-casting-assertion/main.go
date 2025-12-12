package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 10000
	var b uint16 = uint16(a) // casting

	var f1 float32 = 12332.23

	var num1 uint32 = uint32(f1)

	println(a, b, f1, num1)

	str1 := fmt.Sprint("Hello World", true, num1)
	fmt.Println("Type of str1:", reflect.TypeOf(str1))

	var (
		n1 int     = 12312
		n2 uint    = 232323
		n3 uint64  = 1231232232
		n4 float32 = 2321.232
		n5 float64 = 44545454.45
		n6 uint8   = 255
		n7 int16   = -12312
	)

	var any1 any = 56567 // int

	// var n9 int = any1.(int) // any.(T) type assertion

	// var n9 float64 = float64(any1.(int)) // any.(T) type assertion

	// Dataptr, TypePtr
	// Type Ptr of any1 is int

	//var n8 int = any1

	var any2 any = n5 // The TypePtr of any1 is float64

	var sum float64 = float64(any1.(int)) + any2.(float64) + float64(n1) + float64(n2) + float64(n3) + float64(n4) + n5 + float64(n6) + float64(n7)
	fmt.Printf("sum:%.3f\n", sum)
	// var num = 1231231
	// fmt.Printf("sum:%d\n", num)

	var any3 any /// What is the value and what is the type of any
	// value is nil and also type is nil

	fmt.Printf("Value: %v Type: %T\n", any3, any3)

	any3 = 100
	fmt.Printf("Value: %v Type: %T\n", any3, any3)

	println(any3.(int) + 200)

	any3 = true
	fmt.Printf("Value: %v Type: %T\n", any3, any3)

	//any3 = 12312312
	v, ok := any3.(int) // can return two values , ok is true means successfully asserted, else failed to assert

	if ok {
		println(v + 200)
	} //else {
	// 	println("failied to assert to int")
	// }
	//println(any3.(int) + 200) // runtime panic
	// true + 200

	any3 = "Hello World"
	fmt.Printf("Value: %v Type: %T", any3, any3)

}

// a number type can be casted to another number type.
// may lose data based on the size of the variable

// take variables of all number types and perform an arithmetic operations +
