package action

import (
	"fmt"
)

func DataTypeInt() {
	var age uint64 = 10
	fmt.Printf("my age is %d\n", age)
}

func DataTypeSting() {
	var name string = "wt"
	fmt.Printf("my name is %s\n", name)
}

func DataTypeFloat() {
	var height float64 = 1.74
	fmt.Printf("my height is %fm\n", height)
}

func DataTypeChar() {
	var bloodType byte = 'O'
	var character rune = '中'
	fmt.Printf("My bloodType is %c\n", bloodType)
	fmt.Printf("My favorite character is %c\n", character)
}

func DataTypeEscape() {
	var sentence string = "It\"s my name called WT"
	fmt.Println(sentence)
}

func DataTypeMultiLine() {
	var lines string = `
name
age
height
weight
`
	fmt.Println(lines)
}

func DataTypeBool() {
	var result bool = true
	fmt.Printf("The result is %t\n", result)
}

func DataTypeZero() {
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
}
