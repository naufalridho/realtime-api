package utilctx

import (
	"context"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
)

func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "start_time", time.Now())
	return ctx
}

func NewContextFromRequest(r *http.Request) context.Context {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "start_time", time.Now())
	ctx = context.WithValue(ctx, "path", r.URL.Path)
	ctx = context.WithValue(ctx, "uri", r.URL.Path)
	ctx = context.WithValue(ctx, "method", r.Method)
	return ctx
}

func StartSpanFromContext(ctx context.Context, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, "tax_calculator", opts...)
}
