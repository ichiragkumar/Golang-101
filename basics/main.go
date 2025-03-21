package main

import (
	getstring "basics/utils"
	"fmt"
)

func main() {
	fmt.Println("hi ichirag how are you doning ")
	fmt.Println("welcome to go lang learning ")

	getstring.PrintString("hii chirag")
	getstring.PrintHii()

	fmt.Println("bye bye ")
	fmt.Println("bye bye ")

	fmt.Println("welcome back to go lang serires ")

	fmt.Println("hi hcirag")

	// runes in go lonag

	fmt.Println('a')

	// number is golnag
	fmt.Println(123)

	// number can be of float
	fmt.Println(123.456)

	// how to creaet variable in go lang

	var name2 string
	name2 = "ichirag"
	fmt.Println(name2)

	// create variable for int data type
	var name3 int
	name3 = 123
	fmt.Println(name3)

	// create variable for float64 data type
	var name4 float64
	name4 = 123.456
	fmt.Println(name4)

	// create variable for boolean data type
	var name5 bool
	name5 = true
	fmt.Println(name5)

	// create variable for rune data type
	var name6 rune
	name6 = 'a'
	fmt.Println(name6)

	// create variable for each data type
	var name7 byte
	name7 = byte(123)
	fmt.Println(name7)

	// create variable for each data type
	var name8 uint
	name8 = uint(123)
	fmt.Println(name8)

	// create variable for each data type
	var name9 uint8
	name9 = uint8(123)
	fmt.Println(name9)

	// create variable for each data type
	var n1ame uint16
	n1ame = uint16(123)
	fmt.Println(n1ame)

	// create variable for each data type
	var n2ame uint32
	n2ame = uint32(123)
	fmt.Println(n2ame)

	// create variable for each data type
	var n3ame uint64
	n3ame = uint64(123)
	fmt.Println(n3ame)

	// creaet more than one variavle at a time
	var name, age string
	name = "ichirag"
	age = "123"
	fmt.Println(name, age)

}
