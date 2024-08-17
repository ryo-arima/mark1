package model

type Members struct {
	ID        uint    `gorm:"primaryKey,autoIncrement"`
	UUID      string  `gorm:"type:char(36);not null;"`
	GroupUUID string  `gorm:"type:char(36);not null;"`
	UserUUID  string  `gorm:"type:char(36);not null;"`
	Name      string  `gorm:"type:longtext;not null"`
	Role      string  `gorm:"type:enum('group_admin','group_editor','group_viewer');default:'group_viewer'"`
	CreatedAt *string `gorm:"type:datetime(0);not null"`
	UpdatedAt *string `gorm:"type:datetime(0);not null"`
	DeletedAt *string `gorm:"type:datetime(0)"`
}
