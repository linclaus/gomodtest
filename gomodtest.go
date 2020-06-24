package main

import(
	"fmt"
	util "github.com/linclaus/goutil/util"
	"rsc.io/quote"
	myutil "github.com/linclaus/gomodtest/util"
)

func main(){
	fmt.Println("hello world")
	util.Util()
	myutil.Util()
	fmt.Println(quote.Hello())
}
