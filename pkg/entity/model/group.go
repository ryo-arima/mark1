package model

import "time"

type Groups struct {
	ID        uint       `gorm:"primaryKey,autoIncrement"`
	UUID      string     `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
}
