package model

type Agents struct {
	ID        uint     `gorm:"primaryKey,autoIncrement"`
	UUID      string   `gorm:"type:char(36);not null;"`
	Shells    []string `gorm:"shells"`
	Name      string   `gorm:"not null"`
	IP        string   `gorm:"not null"`
	Status    string   `gorm:"type:enum('UP','ERROR','UNKNOWN');default:'DOWN'"`
	CreatedAt *string  `gorm:"type:datetime(0);not null"`
	UpdatedAt *string  `gorm:"type:datetime(0);not null"`
	DeletedAt *string  `gorm:"type:datetime(0)"`
}
