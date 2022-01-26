package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"sweet/internal/conf"
)

func NewMetrics(c *conf.Server) (*prometheus.HistogramVec,*prometheus.CounterVec,*prometheus.GaugeVec) {
	_metricHistogramSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "http_durations_histogram_seconds",
		Help:      "请求耗时统计.",
		Buckets:  []float64{0.001,0.01,0.03,0.05,0.07,0.1,0.2,0.5,1,2,5,10,30},
	}, []string{"kind", "operation"})
	_metricRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "接口请求数与延迟统计",
	}, []string{"kind", "operation","code", "reason"})
	_metricGaugeSeconds := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "http_durations",
		Help:      "各接口实时延迟统计",
	},[]string{"kind", "operation"})

	prometheus.MustRegister(_metricHistogramSeconds, _metricRequests, _metricGaugeSeconds)
	return _metricHistogramSeconds,_metricRequests,_metricGaugeSeconds
}
