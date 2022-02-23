package main

import (
	"fmt"
)

type typeParam interface {
	int | string
}

type hoge[T typeParam] struct {
	param1 T
}

func (h *hoge[T]) SetParam1(p T) {
	h.param1 = p
}

func (h *hoge[T]) echoParam1() {
	fmt.Println(h.param1)
}

type fuga[T typeParam] interface {
	Echo(T)
}

type fugaString struct{}

func (fs *fugaString) Echo(s string) {
	fmt.Println(s)
}

type fugaInt struct{}

func (fi *fugaInt) Echo(s int) {
	fmt.Println(s)
}

func main() {
	h1 := &hoge[string]{"h1"}
	h1.echoParam1()
	h1.SetParam1("h11")
	h1.echoParam1()
	// This causes error
	// h1.SetParam1(1)

	h2 := &hoge[int]{1}
	h2.echoParam1()
	h2.SetParam1(11)
	h2.echoParam1()
	// This causes error
	// h1.SetParam1("h1")

	// int64 does not implement typeParam
	// h3 := &hoge[int64]{1}

	var fs fuga[string]
	var fi fuga[int]

	fs = &fugaString{}
	fi = &fugaInt{}

	fs.Echo("fuga")
	fi.Echo(1)

	// This causes error
	// fs = &fugaInt{}
}
