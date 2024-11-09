package models

import (
	"time"
)

type Subscriber struct {
	ID           uint // Standard field for the primary key
	Name         string
	RollNo       string
	PhoneNo      string
	EmailID      string
	SubscribedAt time.Time
	CreatedAt    time.Time // Automatically managed by GORM for creation time
	UpdatedAt    time.Time // Automatically managed by GORM for update time
}
