package model

import "time"

type Members struct {
	ID        uint   `gorm:"primaryKey,autoIncrement"`
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4()"`
	GroupUUID string
	UserUUID  string
	Role      string     `gorm:"type:enum('group_admin','group_editor','group_viewer');default:'group_viewer'"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
}
