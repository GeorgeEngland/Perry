package store

import (
	"fmt"

	"github.com/GeorgeEngland/perry/data"
)

type AlertStore interface {
	GetAll() ([]data.AlertData, error)
	StoreAlerts([]data.AlertData) error
}

func RegisterAlerts(alerts []data.AlertData, store AlertStore) ([]data.NotificationData, error) {
	store.StoreAlerts(alerts)
	allAlerts, err := store.GetAll()
	if err != nil {
		return nil, err
	}
	n := []data.NotificationData{}
	for k, v := range allAlerts {
		n = append(n, data.NotificationData{
			Id:    k,
			Alert: v,
		})
	}
	fmt.Println("STORE STATE: ", allAlerts)
	return n, nil
}
