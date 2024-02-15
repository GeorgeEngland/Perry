package perry

import (
	"testing"

	"github.com/GeorgeEngland/perry/data"
	"github.com/GeorgeEngland/perry/store"
	"github.com/stretchr/testify/assert"
)

func IntegrationTest(t *testing.T) {
	db := store.NewInMemoryStore()
	checker := NewInMemoryAlertChecker()
	sender := NewLoggingNotificationSender()

	alerts, err := checker.GetAlerts()
	assert.NoError(t, err)
	ns, err := store.RegisterAlerts(alerts, db)
	assert.NoError(t, err)
	for _, v := range ns {
		_, err = SendNotification(sender, v)
		assert.NoError(t, err)
	}
}

func TestMultipleAlerts(t *testing.T) {
	var tests = []struct {
		name  string
		input []data.AlertData
		want  int
	}{
		// the table itself
		{"1 alert for 1 notification",
			[]data.AlertData{
				{Id: 1, Priority: 1}},
			1},
		{"2 alerts for 2 notifications", []data.AlertData{
			{Id: 1, Priority: 1},
			{Id: 2, Priority: 1}},
			2}}
	// The execution loop
	for _, tt := range tests {
		db := store.NewInMemoryStore()

		t.Run(tt.name, func(t *testing.T) {
			ns, err := store.RegisterAlerts(tt.input, db)
			assert.NoError(t, err)
			assert.Equal(t, len(ns), tt.want)
		})

	}
}
