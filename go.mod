module ruok

go 1.13

replace sortssyn => ./sorts

require (
	github.com/HdrHistogram/hdrhistogram-go v0.9.0 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.6.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/halfrost/LeetCode-Go v0.0.0-20200817031238-7989003450a2
	github.com/influxdata/influxdb v1.8.3
	github.com/jinzhu/gorm v1.9.16
	github.com/knqyf263/go-deb-version v0.0.0-20190517075300-09fca494f03d
	github.com/mitchellh/mapstructure v1.3.3
	github.com/nsqio/go-nsq v1.0.8
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
