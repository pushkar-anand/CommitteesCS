package handlers

import (
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type recoveryHandler struct {
	handler    http.Handler
	logger     *logrus.Logger
	production bool
}

// RecoveryHandler is HTTP middleware that recovers from a panic,
// logs the panic, writes http.StatusInternalServerError, and
// continues to the next handler.
func RecoveryHandler(logger *logrus.Logger, production bool) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		r := &recoveryHandler{
			handler:    h,
			logger:     logger,
			production: production,
		}

		return r
	}
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			h.log(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
	}()

	h.handler.ServeHTTP(w, r)
}

func (h recoveryHandler) log(v ...interface{}) {
	if h.production {
		h.logger.WithField("stack", string(debug.Stack())).Error(v...)
	} else {
		h.logger.Error(v...)
		debug.PrintStack()
	}
}
