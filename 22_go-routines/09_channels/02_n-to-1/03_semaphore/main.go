package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {

		//此处done是为了保证两个写入c的协程运行完毕 再关闭c
		//要读出done 得先等待写入done，done写入的时候，c就已经写入完毕了
		//time.Sleep(300 * time.Millisecond)可以验证此处逻辑，等待c执行完
		//使用go 是为了不阻塞main的执行

		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
