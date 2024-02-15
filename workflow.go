package perry

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/GeorgeEngland/perry/data"
	"github.com/GeorgeEngland/perry/store"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Initialise our engine
// Check for any alerts
// Register alert with enging
// Ask engine for any notifications
// Send notifications
func HandleAlertsWorkflow(ctx workflow.Context, input HandleAlertsWorkflowInput) ([]data.NotificationData, error) {
	logger := workflow.GetLogger(ctx)
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		// Retry indefinitely
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    math.MaxInt32,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *Activities

	logger.Info("Checking Alerts...")
	alerts := []data.AlertData{}
	// activity to get alerts (non-deterministic)
	err := workflow.ExecuteActivity(ctx, a.GetAlertsActivity, getAlertsRequest{}).Get(ctx, &alerts)
	if err != nil {
		fmt.Println("GET ALERTS: Non-retryable error or max attempts reached", err)
		return nil, err
	}
	logger.Info("Registering alerts:", alerts)
	var notifications []data.NotificationData
	err = workflow.ExecuteActivity(ctx, a.RegisterAlertsActivity, alerts).Get(ctx, &notifications)
	if err != nil {
		logger.Error("REGISTER ALERTS ERROR: non retryable error when registering alerts", err)
	}
	logger.Info("Notifications to Send", notifications)
	for _, v := range notifications {
		notificationResult := data.NotificationData{}
		// Activity to send alerts
		// non-deterministic
		err := workflow.ExecuteActivity(ctx, a.SendNotificationActivity, v).Get(ctx, &notificationResult)
		if err != nil {
			logger.Error("SEND NOTIFICATION: Non-retryable error or max attempts reached for notification", err)
			return nil, err
		}
	}
	return notifications, nil

}

type Activities struct {
	Checker AlertChecker
	Sender  NotificationSender
	Store   store.AlertStore
}

type getAlertsRequest struct {
}

func (a *Activities) RegisterAlertsActivity(ctx context.Context, data []data.AlertData) ([]data.NotificationData, error) {
	return store.RegisterAlerts(data, a.Store)
}

func (a *Activities) GetAlertsActivity(ctx context.Context, _ getAlertsRequest) ([]data.AlertData, error) {
	return GetNewAlerts(a.Checker)
}

func (a *Activities) SendNotificationActivity(ctx context.Context, data data.NotificationData) (data.NotificationData, error) {
	return SendNotification(a.Sender, data)
}
