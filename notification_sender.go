package perry

import (
	"fmt"

	"github.com/GeorgeEngland/perry/data"
)

type loggingNotificationSender struct {
}

func NewLoggingNotificationSender() *loggingNotificationSender {
	return &loggingNotificationSender{}
}

func (n *loggingNotificationSender) SendNotification(a data.NotificationData) (data.NotificationData, error) {
	fmt.Println("SENT NOTIFICATION: ", a)
	return a, nil
}
