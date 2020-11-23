package api

import (
	"net/http"
	"time"

	"weather_service/adapters/domain/metrics"
)

func MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			metrics.TotalRequestCounter.WithLabelValues(r.URL.Path).Inc()
			metrics.RequestLatencySummary.WithLabelValues(r.URL.Path).Observe(time.Since(begin).Seconds())
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}
