package main

import (
	"fmt"

	myutil "github.com/linclaus/gomodtest/util"
	util "github.com/linclaus/goutil/util"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")
	util.Util()
	myutil.Util()
	fmt.Println(quote.Hello())
}
