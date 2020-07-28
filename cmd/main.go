package main

import (
	"fmt"
	"net/http"
	"time"

	test "github.com/linclaus/gomodtest/test"
	myutil "github.com/linclaus/gomodtest/util"
	util "github.com/linclaus/goutil/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"rsc.io/quote"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	logrus.Info("listen: 8080")
	go http.ListenAndServe(":8080", nil)

	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:

			fmt.Println("hello world")
			util.Util()
			myutil.Util()
			fmt.Println(quote.Hello())
			test.Test()

		}
	}
}
