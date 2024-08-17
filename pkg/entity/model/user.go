package model

type Users struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UUID      string  `gorm:"type:char(36);not null;"`
	Email     string  `gorm:"type:varchar(191);not null;uniqueIndex"`
	Name      string  `gorm:"type:longtext;not null"`
	Password  string  `gorm:"type:longtext;not null"`
	Status    string  `gorm:"type:enum('EmailVerified','EmailNotVerified','Locked');default:'EmailNotVerified';not null"`
	CreatedAt *string `gorm:"type:datetime(0);not null"`
	UpdatedAt *string `gorm:"type:datetime(0);not null"`
	DeletedAt *string `gorm:"type:datetime(0)"`
}
