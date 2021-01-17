package middleware

import (
	"committees/validation"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type uuidHandler struct {
	params []string
	next   http.Handler
	logger *logrus.Logger
}

func (uh *uuidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)

	for _, param := range uh.params {
		if id, ok := ids[param]; ok && !validation.IsValidUUID(id) {
			uh.logger.Debugf("%s is not a valid UUID", id)

			return
		}
	}

	uh.next.ServeHTTP(w, r)
}

// UUIDValidator provides a middleware to validate UUID in request URI.
// It receives a list of params that need to be checked if it's a valid UUID
func UUIDValidator(logger *logrus.Logger, params ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return &uuidHandler{
			params: params,
			next:   next,
			logger: logger,
		}
	}
}
