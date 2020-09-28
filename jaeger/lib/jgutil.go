package lib

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)
type TracerObjectStruct struct{
	tracer opentracing.Tracer
	closer io.Closer
}

func NewTracerObjectStruct(tracer opentracing.Tracer, closer io.Closer) *TracerObjectStruct {
	return &TracerObjectStruct{tracer: tracer, closer: closer}
}
var TraceObject *TracerObjectStruct
func InitTraceConfig(){
	cfg := config.Configuration{
		ServiceName:"OrderService",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	tracer,closer,err:=cfg.NewTracer()
	if err!=nil{
		log.Fatal(err)
	}
	TraceObject=NewTracerObjectStruct(tracer,closer)
	opentracing.SetGlobalTracer(TraceObject.tracer)
}