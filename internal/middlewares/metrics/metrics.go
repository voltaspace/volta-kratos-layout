package metrics

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metrics"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"strconv"
	"time"
)

type DC struct {
	Start int64
	End int64
}

var DCS = map[string]DC{
	"0-30ms":{0,30},
	"30-70ms":{30,70},
	"70-100ms":{70,100},
	"100-300ms":{100,300},
	"300-1000ms":{300,1000},
}

// Option is metrics option.
type Option func(*options)

// WithRequests with requests counter.
func WithRequests(rk string,c metrics.Counter) Option {
	return func(o *options) {
		o.requests = c
		o.rk = rk
	}
}

// WithSeconds with seconds histogram.
func WithSeconds(rk string,c metrics.Observer) Option {
	return func(o *options) {
		o.seconds = c
		o.rk = rk
	}
}

// WithSeconds with seconds histogram.
func WithApiSeconds(rk string,c metrics.Gauge) Option {
	return func(o *options) {
		o.apiSeconds = c
		o.rk = rk
	}
}

type options struct {
	rk string
	// counter: <client/server>_requests_code_total{kind, operation, code, reason}
	requests metrics.Counter
	// histogram: <client/server>_requests_seconds_bucket{kind, operation}
	seconds metrics.Observer

	apiSeconds metrics.Gauge
}

func Server(opts ...Option) middleware.Middleware {
	op := options{}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				op.requests.With(kind, operation, strconv.Itoa(code), reason).Inc()
			}
			if op.seconds != nil {
				op.seconds.With(kind, operation).Observe(time.Since(startTime).Seconds())
			}
			if op.apiSeconds != nil {
				ms,_ := strconv.ParseFloat(strconv.FormatInt(time.Since(startTime).Milliseconds(),10),64)
				op.apiSeconds.With(kind, operation).Set(ms)
			}

			return reply, err
		}
	}
}

// Client is middleware client-side metrics.
func Client(opts ...Option) middleware.Middleware {
	op := options{}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				op.requests.With(kind, operation, strconv.Itoa(code), reason).Inc()
			}
			if op.seconds != nil {
				op.seconds.With(kind, operation).Observe(time.Since(startTime).Seconds())
			}
			return reply, err
		}
	}
}

func delayClassify(t int64) string{
	for k,v := range DCS {
		if t >= v.Start && t < v.End {
			return k
		}
	}
	return ">1000ms"
}