package main

// 1. any type
// 2. type cast
// 3. type conversion using strconv package
import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var any1 any = 10 // any and interface{} are same . any is also statically typed and so type safe

	any1 = true

	any1 = 123.123

	any1 = "Hello World"

	fmt.Println("any1:", any1)

	var any2 interface{} = 10 // any and interface{} are same

	any2 = true

	any2 = 123.123

	any2 = "Hello Golang Minds"

	any2 = any1

	fmt.Println("any2:", any2)

	// type casting

	// there is no implicit type casting in Go

	var num1 uint8 = 234
	//var num2 uint64 = num1 // It is not allowed, unless type casted
	var num2 uint64 = uint64(num1)

	num2++

	println(num2)

	float1 := 432.343 // float64

	num2 = uint64(float1) // obbiously it lose data after decimal but type casting can be done

	println(num2)

	float2 := float32(432.343) // type casted while creating

	fmt.Println("Value of float2", float2, " type of float2:", reflect.TypeOf(float2))

	// var b1 uint8 = 1

	// var ok1 bool = bool(b1) // This is not accepted, bool to number or number to bool is not accepeted

	// can convert any type value to string

	num2 = 312312312

	str1 := strconv.Itoa(int(num2))

	fmt.Printf("str1: %s Type:%T \n", str1, str1)

	str2 := fmt.Sprint(num2) // kind of a formatter
	fmt.Printf("str2: %s Type:%T \n", str2, str2)

	// everything can be convereted to string using Sprint, Sprintln, Sprintf kind of formatters

	str3 := fmt.Sprint("-->", num1, " ", float1, " ", float2, " ", str1, " ", str2, "<--")

	println(str3)
	a, b, c, d := Calc(823, 345)

	println("add:", a, "sub", b, "mul", c, "div", d)

	a1, _, c1, _ := Calc(823, 345) // _ blank identifier

	println("add:", a1, "mul", c1)

	// from string to int

	str4 := "a69862322"

	val, err := strconv.Atoi(str4) // some fields are nil in Go , nil is similar to null but much more matured

	if err != nil { // why check error with nil is error is an interface
		println(err.Error())
	} else {
		println("successfully converted to int", val)
	}

}

// please go through these functions from strconv package
// ParseInt,
// ParseBool
// ParseFloat

// thre is one type of variable that is capable of storing any kind of data.

// system.object --> Object --> any kind of data in C# and Java

// any or interface{} is a data type using that we can store any kind of data.

// any number type can be type casted to any number type

func Calc(a, b int) (add int, sub int, mul int, div int) { // named return types
	add, sub, mul, div = a+b, a-b, a*b, a/b
	return add, sub, mul, div
}

func Calc1(a, b int) (int, int, int, int) { // unnamed return types
	add, sub, mul, div := a+b, a-b, a*b, a/b
	return add, sub, mul, div
}
