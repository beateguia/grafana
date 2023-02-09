package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Historian struct {
	TransitionsTotal      *prometheus.CounterVec
	TransitionsFailed     *prometheus.CounterVec
	WritesTotal           prometheus.Counter
	WritesFailed          prometheus.Counter
	ActiveWriteGoroutines prometheus.Gauge
	PersistDuration       prometheus.Histogram
}

func NewHistorianMetrics(r prometheus.Registerer) *Historian {
	return &Historian{
		TransitionsTotal: promauto.With(r).NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "state_history_transitions_total",
			Help:      "The total number of state transitions processed by the state historian.",
		}, []string{"org"}),
		TransitionsFailed: promauto.With(r).NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "state_history_transitions_failed_total",
			Help:      "The total number of state transitions that failed to be written.",
		}, []string{"org"}),
		WritesTotal: promauto.With(r).NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "state_history_batch_writes_total",
			Help:      "The total number of state history batches that were attempted to be written.",
		}),
		WritesFailed: promauto.With(r).NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "state_history_batch_writes_failed_total",
			Help:      "The total number of failed writes of state history batches.",
		}),
		PersistDuration: promauto.With(r).NewHistogram(prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "state_history_persist_duration_seconds",
			Help:      "Histogram of write times to the state history store.",
			Buckets:   prometheus.DefBuckets,
		}),
	}
}
