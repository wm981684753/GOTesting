package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//-指针也可以直接调用结构体内容
	p1 := &person{"James", 20}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
