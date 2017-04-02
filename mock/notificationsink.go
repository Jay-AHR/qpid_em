package mock

import (
	"log"

	"github.com/Jay-AHR/qpid_em"
)

// NotificationSink is a mock qpid.NotificationSink
type NotificationSink struct{}

// NewNotificationSink returns an initialized NotificationSink
func NewNotificationSink() *NotificationSink {
	return &NotificationSink{}
}

// Listen processes grill statuses
func (n *NotificationSink) Listen(notifications chan qpid.GrillStatus) {
	for message := range notifications {
		log.Printf("GrillStatus: %#v", message)
	}
}
