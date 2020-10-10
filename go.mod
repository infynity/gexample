module ruok

go 1.13

replace sortssyn => ./sorts

require (
	github.com/HdrHistogram/hdrhistogram-go v0.9.0 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shenyisyn/goft-expr v0.3.0
	github.com/stretchr/testify v1.4.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.3.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	sortssyn v0.0.0-00010101000000-000000000000
)
