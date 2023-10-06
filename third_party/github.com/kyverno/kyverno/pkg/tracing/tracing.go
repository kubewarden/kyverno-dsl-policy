package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

const limit = 256

func StringValue(value string) string {
	panic("stub")
}

func StartChildSpan(ctx context.Context, tracerName string, operationName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	panic("stub")
}

func ChildSpan(ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span), opts ...trace.SpanStartOption) {
	panic("stub")
}

func ChildSpan1[T1 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) T1, opts ...trace.SpanStartOption) T1 {
	return doFn(ctx, nil)

}

func ChildSpan2[T1 any, T2 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) (T1, T2), opts ...trace.SpanStartOption) (T1, T2) {
	return doFn(ctx, nil)

}

func ChildSpan3[T1 any, T2 any, T3 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) (T1, T2, T3), opts ...trace.SpanStartOption) (T1, T2, T3) {
	panic("stub")
}

func NewTraceConfig(log interface{}, tracerName, address, certs string, kubeClient interface{}) (func(), error) {
	panic("stub")
}

func SetSpanStatus(span trace.Span, err error) {
	panic("stub")
}

func SetStatus(ctx context.Context, err error) {
	panic("stub")
}

func SetHttpStatus(ctx context.Context, err error, code int) {
	panic("stub")
}

func IsInSpan(ctx context.Context) bool {
	panic("stub")
}

func CurrentSpan(ctx context.Context) trace.Span {
	panic("stub")
}

func SetAttributes(ctx context.Context, kv ...interface{}) {
	panic("stub")
}

func RequestFilterIsInSpan(request *http.Request) bool {
	panic("stub")
}

func Transport(base http.RoundTripper, opts ...otelhttp.Option) *otelhttp.Transport {
	panic("stub")
}

func StartSpan(ctx context.Context, tracerName string, operationName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	panic("stub")
}

func Span(ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span), opts ...trace.SpanStartOption) {
	panic("stub")
}

func Span1[T1 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) T1, opts ...trace.SpanStartOption) T1 {
	panic("stub")
}

func Span2[T1 any, T2 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) (T1, T2), opts ...trace.SpanStartOption) (T1, T2) {
	panic("stub")
}

func Span3[T1 any, T2 any, T3 any](ctx context.Context, tracerName string, operationName string, doFn func(context.Context, trace.Span) (T1, T2, T3), opts ...trace.SpanStartOption) (T1, T2, T3) {
	panic("stub")
}

type Embedme interface{}
