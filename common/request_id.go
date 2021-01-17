package common

import (
	"context"
	"errors"
	"net/http"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// contextKeyRequestID is the ContextKey for RequestID
const contextKeyRequestID ContextKey = "requestID"

// ErrorReqIDRetrieval is returned when request ID couldn't be retrieved
var ErrorReqIDRetrieval = errors.New("request ID couldn't be retrieved")

func AddRequestIDToCTX(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, contextKeyRequestID, reqID)
}

func GetRequestID(r *http.Request) (string, error) {
	return GetRequestIDFromCTX(r.Context())
}

func GetRequestIDFromCTX(ctx context.Context) (string, error) {
	reqID := ctx.Value(contextKeyRequestID)
	if ret, ok := reqID.(string); ok {
		return ret, nil
	}

	return "", ErrorReqIDRetrieval
}
