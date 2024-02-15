package perry

import (
	"math/rand"

	"github.com/GeorgeEngland/perry/data"
)

type inMemoryAlertChecker struct {
}

func NewInMemoryAlertChecker() *inMemoryAlertChecker {
	return &inMemoryAlertChecker{}
}

func (m *inMemoryAlertChecker) GetAlerts() ([]data.AlertData, error) {
	index := rand.Intn(2)
	return []data.AlertData{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}[index:], nil
}
