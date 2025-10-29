package server

import (
	"github.com/google/uuid"
	"sync"
)

type MemberStore struct {
	Mu      sync.RWMutex
	Members map[uuid.UUID]*Member
}

type Channel struct {
	Id          uuid.UUID
	Name        string
	Tag         string
	MemberStore MemberStore
}
