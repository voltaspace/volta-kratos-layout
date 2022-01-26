package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "sweet/api/pb/v1"
	"sweet/internal/conf"
	"sweet/internal/filter"
	"sweet/internal/middlewares"
	"sweet/internal/middlewares/metrics"
	"sweet/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.SweetService, logger log.Logger) *http.Server {
	_metricHistogramSeconds,_metricRequests,_metricGaugeSeconds := NewMetrics(c)
	var opts = []http.ServerOption{
		http.Middleware(
			middlewares.Auth(),
			recovery.Recovery(),
			metrics.Server(
				metrics.WithSeconds("_metricHistogramSeconds",prom.NewHistogram(_metricHistogramSeconds)),
				metrics.WithRequests("_metricRequests",prom.NewCounter(_metricRequests)),
				metrics.WithApiSeconds("_metricHistogramSeconds",prom.NewGauge(_metricGaugeSeconds)),
			),
		),
		http.Filter(
			filter.CorsFilter,
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	NewSwaggerServer(c,srv)
	TracerProvider("http://127.0.0.1:14268/api/traces")
	srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterSweetHTTPServer(srv, service)
	return srv
}
