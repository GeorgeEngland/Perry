package perry

import (
	"time"

	"github.com/GeorgeEngland/perry/data"
	"github.com/GeorgeEngland/perry/store"
)

type HandleAlertsWorkflowInput struct {
	CheckTime time.Duration
}

type AlertChecker interface {
	GetAlerts() ([]data.AlertData, error)
}

type NotificationSender interface {
	SendNotification(a data.NotificationData) (data.NotificationData, error)
}
type AlertEngine interface {
	RegisterAlerts([]data.AlertData, store.AlertStore) []data.NotificationData
}

func GetNewAlerts(checker AlertChecker) ([]data.AlertData, error) {
	return checker.GetAlerts()
}

func SendNotification(sender NotificationSender, data data.NotificationData) (data.NotificationData, error) {
	return sender.SendNotification(data)
}

const CheckAlertsQueue = "ALERT_CHECK_QUEUE"
