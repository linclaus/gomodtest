package main

import (
	"fmt"
	"time"

	test "github.com/linclaus/gomodtest/test"
	myutil "github.com/linclaus/gomodtest/util"
	util "github.com/linclaus/goutil/util"
	"rsc.io/quote"
)

func main() {
	for {
		time.Sleep(10000)
		fmt.Println("hello world")
		util.Util()
		myutil.Util()
		fmt.Println(quote.Hello())
		test.Test()
	}
}
