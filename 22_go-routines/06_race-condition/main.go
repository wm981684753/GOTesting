package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x

		fmt.Println(s,counter)
		zz := rand.Intn(3)
		fmt.Println(s, i, "Counter:", zz, counter,zz,time.Duration(zz))
	}
	wg.Done()
}
// 数据溢出了，两个协程存在竞争，不安全
// go run -race main.go //-race用于检测数据的安全性，可以检测数据竞争
// vs
// go run main.go
