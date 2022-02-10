/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-10 13:20:12
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-10 13:20:12
 */
package common

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

//创建jaeger链路追踪实例
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	return tracer, closer, err
}
