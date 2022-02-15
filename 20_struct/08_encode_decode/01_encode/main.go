package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}

	dat := make(map[string]string, 5)
	dat["First"] = "James"
	dat["Last"] = "Bond"

	//这里的用途暂时不明白
	json.NewEncoder(os.Stdout).Encode(p1)
	json.NewEncoder(os.Stdout).Encode(dat)

	p1j, _ := json.Marshal(p1)
	fmt.Println("struct转json:", string(p1j))
	mj, _ := json.Marshal(dat)
	fmt.Println("map转json:", string(mj))

	dat2 := make(map[string]string, 5)
	json.Unmarshal(mj, &dat2)
	fmt.Println("json转map", dat2)
}
