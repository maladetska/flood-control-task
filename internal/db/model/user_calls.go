package model

type UserCalls struct {
	UserId     int64  `json:"user_id"`
	CallsCount uint64 `json:"calls_count"`
}
