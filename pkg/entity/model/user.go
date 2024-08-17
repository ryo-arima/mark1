package model

import "time"

type Users struct {
	ID        uint       `gorm:"primaryKey,autoIncrement"`
	UUID      string     `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string     `gorm:"unique;not null"`
	Name      string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
	Status    string     `gorm:"type:enum('EmailVerified','EmailNotVerified','Locked');default:'EmailNotVerified'"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
}
