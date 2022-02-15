package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println("a:", n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println("b:", n)
		}
		done <- true
	}()

	//done是为了保证全部接收后 才结束main主程序
	<-done
	<-done

}
