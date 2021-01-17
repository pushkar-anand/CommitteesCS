package handlers

import (
	c "committees/common"
	"net/http"

	"github.com/google/uuid"
)

const (
	// HTTPHeaderNameRequestID has the name of the header for request ID
	HTTPHeaderNameRequestID = "X-Request-ID"
)

// assignRequestID will attach a brand new request ID to an incoming http request
func assignRequestID(r *http.Request) *http.Request {
	ctx := c.AddRequestIDToCTX(r.Context(), uuid.New().String())
	return r.WithContext(ctx)
}

// AssignRequestIDHandler is handler to assign request ID to each incoming request
// Make sure this is the last handler to ensure request ID is assigned as early as possible
func AssignRequestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = assignRequestID(r)
		id, _ := c.GetRequestID(r)
		w.Header().Set(HTTPHeaderNameRequestID, id)
		next.ServeHTTP(w, r)
	})
}
