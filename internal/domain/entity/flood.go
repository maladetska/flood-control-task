package entity

import (
	"errors"
	"time"
)

var (
	ErrorFloodCountIsInvalid   = errors.New("count of method calls is zero")
	ErrorFloodEndTimeIsInvalid = errors.New("flood end time is invalid")
)

type Flood struct {
	UserID     int64
	CallsCount uint64
	EndAt      time.Time
}

func NewFlood(userID int64, callsCount uint64, endAt time.Time) (Flood, map[string]error) {
	flood := Flood{
		UserID:     userID,
		CallsCount: callsCount,
		EndAt:      endAt,
	}

	errorsMap := flood.Validate()
	if errorsMap != nil {
		return Flood{}, errorsMap
	}

	return flood, nil
}

func (flood *Flood) Validate() map[string]error {
	var errorMessages = make(map[string]error)

	if flood.CallsCount == 0 {
		errorMessages["count_required"] = ErrorFloodCountIsInvalid
	}
	if flood.EndAt.IsZero() {
		errorMessages["end_at_required"] = ErrorFloodEndTimeIsInvalid
	}

	return errorMessages
}

func (flood *Flood) CallMethod(interval time.Duration, maxCallsCount uint64) error {
	now := time.Now()

	if now.Sub(flood.EndAt) > interval {
		flood.CallsCount = 1
		return nil
	}

	flood.CallsCount++

	if flood.CallsCount > maxCallsCount {
		return ErrorFloodCountIsInvalid
	}

	return nil
}
