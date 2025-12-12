package main

import (
	"fmt"
	"unsafe"
)

const PI float32 = 3.14          // data Segment --> RO Segmement
const MSG string = "Hello World" // data Segment --> RO Segmement

var Counter int    //  // uninitialized package level variable or static varialbe in some programmign languages
var Global = 99999 // initialized data segment

func main() {
	// PI = 3.145
	// var any1 any
	str := "Hello Golang Minds, I am learning Golang basics! 🚗" // This is a string literal, generally stored in Data Segment
	var any1 any                                                // what is the zero value of any --> nil
	fmt.Println("any1:", any1, "address:", &any1)
	any1 = str
	fmt.Println("any1:", any1, "address:", &any1)
	{
		str1 := "Hello Golang Minds"
		{
			str2 := "Hello Golang Minds"
			{
				var a, b = 10, 20 // stack memory
				any1 = a
				println(a, b, a+b) // printed
				any1 = a + b
				println(str1)
				any1 = true
				fmt.Println("address a:", &a, "\naddress b:", &b, "\naddress str:", &str, "\naddress str1:", &str1, "\naddress str2:", &str2, "\naddress Counter", &Counter)
			} // This gets deallcated
		}
	}
	fmt.Println("Size of str:", unsafe.Sizeof(str), "Length of str:", len(str))
	//println(str1) // cannot use becase stack pointer pops .. so that it goes out of that region
}

// Process

// Code/Text segment --> Binary/Exe is loaded into Code or Text Segment
// Data Segment --> All constants and global variables are stored Data Segment

// Stack memory --> Generally variables are stored in stack memory, 2mb or 1mb or 8mb
// Heap Memory --> It depends on the memory that is required by your application and subjective to the availability

// string
//

// any
