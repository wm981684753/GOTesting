package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person
	//rdr-1
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	//等同于1,json解析 转struct(转其他类型均可)
	rdr2 := []byte(`{"First":"James", "Last":"Bonds", "Age":30}`)
	json.Unmarshal(rdr2,&p1)

	fmt.Println(rdr,"\n",rdr2)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
