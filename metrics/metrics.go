package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MyMetricGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_metric_gauge",
		Help: "metric test",
	})
	MyMetricGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "my_metric_gauge_vec",
		Help: "metric_vec test",
	}, []string{"label1", "label2"})
)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	MyMetricGauge.Set(0)
	prometheus.MustRegister(MyMetricGauge)
	MyMetricGaugeVec.WithLabelValues("l1", "l2").Set(1)
	MyMetricGaugeVec.WithLabelValues("l3", "l2").Set(2)
	prometheus.MustRegister(MyMetricGaugeVec)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}
