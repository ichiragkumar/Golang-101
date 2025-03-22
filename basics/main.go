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

	// if else in go
	var marks int = 80
	if marks > 60 {
		if marks > 70 {
			fmt.Println("you are good")
			if marks > 70 && marks < 80 {
				fmt.Println("you are very good")
			}
			if marks > 80 {
				fmt.Println("you are excellent")
			}
			if marks > 90 {
				fmt.Println("you are master")
			}
			fmt.Println("you are good 2")

		}
		fmt.Println("you are first class")
	} else {
		if marks > 50 {
			fmt.Println("you are average")
		}
		fmt.Println("you are bad")
	}

	//Learn loop in go

	var counter int = 10
	for i := 0; i <= counter; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for counter < 20 {
		fmt.Println(counter)
		counter++
	}

	// invfinte for loopg in golang
	var counter2 int = 0
	for {
		fmt.Println("infinite loop", counter2)
		if counter2 > 20 {
			break
		}
		counter2++
	}

	var counter4 int = 10
	for p := 0; p <= counter4; p++ {
		fmt.Println("i am p", p)
		if p == 5 {
			fmt.Println("i will exit now", p)
			break
		}

	}
	// format specifier in go lang
	var lax string
	name = "ichirag"
	fmt.Printf("Hello %s\n", lax)

	var point int = 10
	fmt.Println("point is %d\n", point)

	var pi float32 = 3.14
	pi = 3.14
	fmt.Println("pi is %f\n", pi)

	// take input from user

	fmt.Println("enter username")
	var username string
	fmt.Scan(&username)

	// take input from user
	fmt.Println("enter password")
	var password2 string
	fmt.Scan(&password2)

	// print the input
	fmt.Println("username is ", username)
	fmt.Println("password is ", password2)

	// use bufio paclage to accept input from user

}
