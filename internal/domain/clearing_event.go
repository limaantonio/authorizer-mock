package domain

type ClearingEvent struct {
	AuthorizationID string
	Amount          float64
	CardNumber      string
	Status          string
}
