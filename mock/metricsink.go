package mock

import (
	"log"

	"github.com/Jay-AHR/qpid_em"
)

type MetricSink struct{}

func NewMetricSink() *MetricSink {
	return &MetricSink{}
}

func (m *MetricSink) Listen(metrics chan qpid.GrillStatus) {
	for message := range metrics {
		log.Printf("METRIC: %#v", message)
	}
}
