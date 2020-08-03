package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/linclaus/gomodtest/metrics"
	"github.com/linclaus/gomodtest/pkg/model"
	test "github.com/linclaus/gomodtest/test"
	myutil "github.com/linclaus/gomodtest/util"
	util "github.com/linclaus/goutil/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"rsc.io/quote"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	logrus.Infof("Listen: %s", *addr)
	go http.ListenAndServe(*addr, nil)

	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			metrics.MyMetricGauge.Inc()
			metrics.MyMetricGaugeVec.WithLabelValues("l1", "l2").Inc()
			metrics.MyMetricGaugeVec.WithLabelValues("l2", "l3").Inc()
			fmt.Println("hello world")
			util.Util()
			myutil.Util()
			fmt.Println(quote.Hello())
			test.Test()
			model.TestType()

		}
	}
}
