package mock

import (
	"log"

	"github.com/Jay-AHR/qpid_em"
)

type AlertSink struct {
}

func NewAlertSink() *AlertSink {
	return &AlertSink{}
}
func (a *AlertSink) Listen(alerts chan qpid.Notification) {
	for message := range alerts {
		log.Printf("ALERT: %#v", message)
	}
}
