package stringutil

type myStruct struct {
	name string
	pwd string
}

//小写仅可内部调用
var pwd = "pwd"

//注意大小写，大写才可外部调用
var Person = myStruct{
	"weiming",
	pwd,
}