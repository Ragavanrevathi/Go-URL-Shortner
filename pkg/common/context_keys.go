package common

type ContextKey string

// defining own type to avoid collision in context
const (
	LoggerKey  ContextKey = "logger"
	TraceIDKey ContextKey = "traceId"
)
