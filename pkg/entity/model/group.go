package model

type Groups struct {
	ID        uint    `gorm:"primaryKey,autoIncrement"`
	UUID      string  `gorm:"type:char(36);not null;"`
	Name      string  `gorm:"not null"`
	CreatedAt *string `gorm:"type:datetime(0);not null"`
	UpdatedAt *string `gorm:"type:datetime(0);not null"`
	DeletedAt *string `gorm:"type:datetime(0)"`
}
