package util

import "context"

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	DebugContextKey = contextKey("debug")
	PortContextKey  = contextKey("port")
)

func GetDebugContextKey(ctx context.Context) (bool, bool) {
	debug, ok := ctx.Value(DebugContextKey).(bool)
	return debug, ok
}

func GetPortContextKey(ctx context.Context) (string, bool) {
	port, ok := ctx.Value(PortContextKey).(string)
	return port, ok
}
