package metrics

import "github.com/prometheus/client_golang/prometheus"

const nameSpace = "weather_service"

var (
	TotalRequestCounter = func() *prometheus.CounterVec {
		m := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: nameSpace,
				Name:      "request_count",
				Help:      "Service requests total counter",
			},
			[]string{"path"},
		)
		prometheus.MustRegister(m)

		return m
	}()

	RequestLatencySummary = func() *prometheus.SummaryVec {
		m := prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace: nameSpace,
				Name:      "request_latency_microseconds",
				Help:      "response latency by path",
			},
			[]string{"path"},
		)
		prometheus.MustRegister(m)

		return m
	}()
)
