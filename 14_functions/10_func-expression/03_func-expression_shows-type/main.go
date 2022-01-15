package main

import "fmt"

func main() {

	//greeting := func() {
	//	fmt.Println("Hello world!")
	//}

	greeting()
	fmt.Printf("%T\n", greeting)
	fmt.Println(greet("sfd","fdsf"))
	fmt.Println(greeting())
}

func greeting() (s int){
	fmt.Println(123)
	return 
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
