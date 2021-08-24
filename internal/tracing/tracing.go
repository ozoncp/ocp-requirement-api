package tracing

import (
	"github.com/opentracing/opentracing-go"
	config2 "github.com/ozoncp/ocp-requirement-api/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func InitTracer(cfg config2.Config) io.Closer {
	cfgMetrics := &config.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Address,
		},
	}
	tracer, closer, err := cfgMetrics.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Err(err)
	}
	opentracing.SetGlobalTracer(tracer)

	return closer
}
