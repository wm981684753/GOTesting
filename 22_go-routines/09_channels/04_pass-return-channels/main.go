package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)

	// chan为双向通道，此处还可以写入数据
	go func() {
		fmt.Println("a")
		cSum <- 5
	}()

	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		fmt.Println("a")
		close(out)
	}()
	return out
}
