package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Integers
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is", reflect.TypeOf(3220))

	// unsigned for positives numbers ... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Bytes is ", reflect.TypeOf(b))

	// with sinal for positive numbers ... int8 int16 int32 int64
	maxInt := math.MaxInt64
	fmt.Printf("The maximun value of int is %v and type is %v \n",
		maxInt,
		reflect.TypeOf(maxInt))

	// 'rune' maps Unicode table (int32)
	var unicodeValue rune = 'a'
	fmt.Println("The unicodeValue is a", reflect.TypeOf(unicodeValue))
	fmt.Println("The value of 'a' on Unicode table is ", unicodeValue)

	// Real numbers ... float32, float64
	var f float32 = 49.99
	fmt.Println("The type of f is ", reflect.TypeOf(f))
	fmt.Println("and the type of literal 49.99 is ", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("The type os bo is ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello"
	fmt.Println(s1 + "!")
	fmt.Println("Length of s1 is ", len(s1))

	// String with multiples lines
	s2 := `Hello!
	That
	is
	cool!
	`
	fmt.Println("Length of s2 is ", len(s2))

	// So, we don't have char guys
	char := 'a'
	fmt.Println("char is a ", reflect.TypeOf(char))
	fmt.Println("Your value is  ", char)
	fmt.Println("Do you remember of unicode values? Say TRUE! ", char == unicodeValue)

	// And default values of types
	// You must declare types if you don't initialize
	var (
		i   int
		f64 float64
		bl  bool
		s   string
		n   *int
	)

	fmt.Printf("The default value for i, f64, bl, s and n are %v %v %v %v %v", i, f64, bl, s, n)
	fmt.Printf("The default value for i, f64, bl, s and n are %v %v %v %q %v", i, f64, bl, s, n)

}
