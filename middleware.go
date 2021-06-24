package zaper

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

func LogMiddleware(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		newLogger := wrapReqFields(logger, r)

		r.WithContext(context.WithValue(r.Context(), "logger", newLogger))

		h.ServeHTTP(w, r)

		newLogger.Info("finished handle a request")

	})

}

func wrapReqFields(logger *zap.Logger, r *http.Request) *zap.Logger {
	// TODO 日志中间件，统计请求
	urlQuery := r.URL.Query()
	return logger.With(
		zap.String("logid", urlQuery.Get("")),
		zap.String("caller_ip", r.Header.Get("")),
		zap.String("host_ip", r.Host),
		zap.String("product", urlQuery.Get("")),
		zap.String("module", urlQuery.Get("")),
		zap.String("service_id", urlQuery.Get("")),
		zap.String("instance_id", urlQuery.Get("")),
		zap.String("uri_path", urlQuery.Get("")),
		zap.String("tag", urlQuery.Get("")),
		zap.String("env", urlQuery.Get("")),
		zap.String("sversion", urlQuery.Get("")),
		zap.String("stag", urlQuery.Get("")),
	)
}
