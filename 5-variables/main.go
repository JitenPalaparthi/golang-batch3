package main

// 1. rune and byte

func main() {

	type integer = int // this is called as alias.. means just an another name to the same type

	var num1 integer = 1231

	println(num1)

	// similarly

	// rune is also an alias for int32

	//------- rune ------
	// rune --> rune is not a separate type but it is a alias of an existing type int32
	// why to use rune --> generally used to store chars. It is more of a convention
	// bcz in Go char is nothing but an integer with the limits of utf8 char sequence . ~65000

	// rune is nothing but an integer

	var char1 int32 = 65

	println(string(char1), "-->", char1)

	char1 = 'A'

	println(string(char1)) // casted as string, otherwise char is nothing but int32

	var char2 rune = 'A' + 1

	println(string(char2), "-->", char2) // casted as string, otherwise char is nothing but int32

	// 字 体; 字體; zìtǐ)

	char2 = '字'

	println(string(char2), "-->", char2) // casted as string, otherwise char is nothing but int32

	var char3 uint8 = 'A'
	println(string(char3), "-->", char3) // casted as string, otherwise char is nothing but int32

	//char3 = '体' //  20307

	var char4 rune = 20307

	println(string(char4), "-->", char4) // casted as string, otherwise char is nothing but int32

	str1 := "Hello World" // The whole hello world chars are 1 bytes chars

	println("len:", len(str1))

	str2 := "Hello 字体" // 5+ +3+3
	println("len:", len(str2))

	str3 := "Hello Police car --> 🚗"
	println("len:", len(str3))

}

// So what are utf8 chars are

// 0-255 are 1 byte chars are normal ASCII chars
// 256 -- XXXX 2 byte chars
// 3 bytes
// 4 bytes chars
