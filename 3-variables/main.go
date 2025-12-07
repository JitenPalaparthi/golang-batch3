package main

import "fmt"

var (
	Counter int = 1 // initialized global variable
	Max     int     // uninitialized global variable
) // static variables , package level variables

const (
	SECOND int = 1
	MIN    int = SECOND * 60                   // GetMins() // This is an expression, this expression is evaluated by the compiler
	HOUR       = MIN * 60                      // type inference works, automatically takes as int
	EXPR       = 10 + 20 + SECOND*MIN*500/HOUR // it is evaluated at compiletime
)

// function arguments are generally stack allocated
func GetMins() int {
	return SECOND * 60
}

// func someFunc() {
// 	//println(age) // cant use age here
// 	println(Counter)
// }

func main() {

	println(EXPR)

	// There is an another way to declare variables

	//var num int = 10

	var age = 41 // uint8 but it takes as int

	var num1 = 10 // Type inference , int

	var num2 = 10.5 // Type inference , float64

	var ok1 = true // Type inference . bool

	var str1 = "Hello World" // Type inference string

	var num3 uint // zero value --> 0

	var ok2 bool // zero value --> false

	var str2 string // zero value -->""

	println(age, num1, num2, num3, ok1, ok2, str1, str2)

	// shorthand declaration
	a := 10
	str3 := "Hello World"
	ok3 := true

	println(a, str3, ok3)

	// expressions vs statements

	// arithmetic operators, comparision, logical operators and bitwise operators

	//a,b := 10,20

	a, b := 10, 20

	expr1 := a + b*(a+b)*10 + (a * 2) - (b / 2) + 10%2 + Counter // What kind of expressions--> arithmetic
	// () + - * /

	fmt.Printf("expr value: %v type: %T", expr1, expr1)

	expr2 := a+b*(a+b)*10+(a*2)-(b/2)+10%2 >= 16020

	fmt.Printf("\nexpr value: %v type: %T", expr2, expr2)

	expr3 := a+b*(a+b)*10+(a*2)-(b/2)+10%2 >= 16020 || true
	fmt.Printf("\nexpr value: %v type: %T", expr3, expr3)

	var b1 = 0b1111 // binary number
	var b2 = 0b1010 //
	//       & 1010
	//       | 1111
	b3 := b1 & b2 // bitwise operator

	fmt.Printf("\nbinary:%b decimal:%d", b3, b3)

	b4 := b1 | b2 // bitwise operator

	fmt.Printf("\nbinary:%b decimal:%d", b4, b4)

	// const

	//const PI float32 = 3.14 // constants are created in data segment
	// PI = 3.143
}

// for any type, if you dont assign a value , autoimatically it takes zero value
// zero value of number is 0
// zero value of bool is false
// zero value of string ""
