package components

import (
	"github.com/astaxie/beego"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
	"os"
	"time"
)

var Trace opentracing.Tracer

func InitTrace() opentracing.Tracer {
	endp := os.Getenv("JAEGER_ENDPOINT")

	cfg := config.Configuration{
		ServiceName: "MicroTestService", //自定义服务名称
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   endp,
		},
	}

	metricsFactory := metrics.NullFactory
	tracer, _, err := cfg.NewTracer(config.Metrics(metricsFactory))
	if err != nil {
		beego.Error("cannot initialize Jaeger Tracer", err.Error())
	}

	Trace = tracer

	return tracer
}
