package model

import "time"

type Users struct {
    ID          uint `gorm:"primaryKey,autoIncrement"`
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}