package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) //结构体转换为json []byte格式
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs)) //[]byte转为string
}
