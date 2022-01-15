package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	fmt.Println(student)
	fmt.Println(students[0][0])
	fmt.Println(student == nil)
}
