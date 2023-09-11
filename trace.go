package ptrace

import (
	"github.com/goinbox/pcontext"
	"go.opentelemetry.io/otel/trace"
)

type StartTraceFunc[T pcontext.Context] func(ctx T, spanName string, opts ...trace.SpanStartOption) (T, trace.Span)
