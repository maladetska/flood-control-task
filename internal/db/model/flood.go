package model

import (
	"time"
)

type Flood struct {
	UserID int64     `json:"user_id"`
	EndAt  time.Time `json:"end_at"`
}
