package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/linclaus/gomodtest/metrics"
	"github.com/linclaus/gomodtest/pkg/model"
	test "github.com/linclaus/gomodtest/test"
	myutil "github.com/linclaus/gomodtest/util"
	util "github.com/linclaus/goutil/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"rsc.io/quote"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "lang",
				Value:    "english",
				Usage:    "language for greeting",
				EnvVars:  []string{"MY_LANG"},
				FilePath: "config.yaml",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			fmt.Println(c.String("lang"))
			// do()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func do() {
	test.TestOrm()
	test.TestStructTag()
	test.TestExtend()
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
