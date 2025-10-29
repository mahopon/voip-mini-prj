package server

import (
	"github.com/google/uuid"
	"time"
)

type Status int

const (
	StatusDisconnected Status = iota
	StatusIdle
	StatusBusy
	StatusConnected
)

type Member struct {
	Id         uuid.UUID
	Name       string
	Status     Status
	CreatedAt  time.Time
	ModifiedAt time.Time
}
