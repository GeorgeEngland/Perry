package store

import (
	"github.com/GeorgeEngland/perry/data"
)

type inMemoryStore struct {
	Alerts []data.AlertData
}

func NewInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		Alerts: []data.AlertData{},
	}
}

func (s *inMemoryStore) GetAll() ([]data.AlertData, error) {
	return s.Alerts, nil
}

func (s *inMemoryStore) StoreAlerts(a []data.AlertData) error {
	s.Alerts = append(s.Alerts, a...)
	return nil
}
