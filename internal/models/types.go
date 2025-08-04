package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Skills    []Skill `gorm:"foreignKey:CreatedByID"`
	Balance   float64
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Type      string //User, Admin
}

type Skill struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Type        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID uint
	Creator     User `gorm:"foreignKey:CreatedByID"`
}

type Offer struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	TimeAmount  float64
	ProviderID  uint
	Provider    User `gorm:"foreignKey:ProviderID"`
	SkillID     uint
	Skill       Skill `gorm:"foreignKey:SkillID"`
	IsRequest   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	FromUserID uint
	FromUser   User `gorm:"foreignKey:FromUserID"`
	ToUserID   uint
	ToUser     User `gorm:"foreignKey:ToUserID"`
	SkillID    uint
	Skill      Skill `gorm:"foreignKey:SkillID"`
	OfferID    uint
	Offer      Offer `gorm:"foreignKey:OfferID"`
	TimeAmount float64
	Timestamp  time.Time
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FinishedAt *time.Time // pointer to allow null
	Rating     *float64
	Comment    *string
}
