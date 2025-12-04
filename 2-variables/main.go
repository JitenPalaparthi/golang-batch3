package main

import "fmt"

func main() {
	println("你好 ⭐ 😍") // chinees characters
	fmt.Println("హలో వరల్డ్ \nஹலோ வேர்ல்டு\nಹೆಲೋ ವರ್ಲ್ಡ್")

	println("హలో World")

	var num1 int = 100
	var num2 uint8 = 233

	// group of variables, can be created at once
	var (
		num3   uint16  = 12321
		ok1    bool    = true
		float1 float32 = 232.232312
		float2 float64 = 2312312321312.123231
	)

	println(num1, num2)
	fmt.Println("num3:", num3, "ok1:", ok1, "float1:", float1, "float2", float2)
	fmt.Printf("num3:%d ok1: %v float1: %.2f float2: %.3f", num3, ok1, float1, float2)

}

// numbers
// int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64,float32,float64, rune, byte
// the size of int is on 64 bit machine it is 8 bytes, on 32but machine it is 4 bytes
// uinsigned --> you cannot give negative values

// bool -> true or false

// string --> collections of bytes, can give utf8 chars, ~65k chars

// complex --> complex64, complex128
// any --> interface{} or any
