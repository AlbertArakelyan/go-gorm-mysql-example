package models

import (
	"time"
)

type Registration struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID   int64     `gorm:"not null" json:"event_id"`
	Event     *Event    `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"event"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
