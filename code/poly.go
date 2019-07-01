package code

import "fmt"

type myString struct {
	a []byte
}

type voidStruct struct {
	myString
}

type strInt interface {
	strChange([]byte)
}

func (ins *myString) strChange(str []byte) {
	ins.a = str
}

func Poly() {
	var int strInt
	s := &voidStruct{myString{[]byte("123")}}
	int = s
	int.strChange([]byte("321"))
	fmt.Printf("%v\n", s.a)
}

// interface只有一个比较奇怪的点,接受者是指针的时候,int赋值是一个值的时候,不可以调用
