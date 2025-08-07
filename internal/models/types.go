package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Skills    []Skill    `json:"skills" gorm:"foreignKey:CreatedByID"`
	Balance   float64    `json:"balance"`
	IsActive  bool       `json:"is_active"`
	Type      string     `json:"type"` // admin, user, etc.
	TimeBanks []TimeBank `json:"time_banks" gorm:"many2many:time_bank_users;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TimeBank struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `json:"description"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID"`
	Members     []User    `json:"members" gorm:"many2many:time_bank_users;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Skill struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	AreaID      uint      `json:"area_id"`
	Area        Area      `json:"area" gorm:"foreignKey:AreaID"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Area struct {
	AreaID uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"unique"`
}

type Offer struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeAmount  float64   `json:"time_amount"`
	ProviderID  uint      `json:"provider_id"`
	Provider    User      `json:"provider" gorm:"foreignKey:ProviderID"`
	SkillID     uint      `json:"skill_id"`
	Skill       Skill     `json:"skill" gorm:"foreignKey:SkillID"`
	TimeBankID  uint      `json:"time_bank_id"`
	TimeBank    TimeBank  `json:"time_bank" gorm:"foreignKey:TimeBankID"`
	IsActive    bool      `json:"is_active"`
	IsRequest   bool      `json:"is_request"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	FromUserID uint       `json:"from_user_id"`
	FromUser   User       `json:"from_user" gorm:"foreignKey:FromUserID"`
	ToUserID   uint       `json:"to_user_id"`
	ToUser     User       `json:"to_user" gorm:"foreignKey:ToUserID"`
	SkillID    uint       `json:"skill_id"`
	Skill      Skill      `json:"skill" gorm:"foreignKey:SkillID"`
	OfferID    uint       `json:"offer_id"`
	Offer      Offer      `json:"offer" gorm:"foreignKey:OfferID"`
	TimeBankID uint       `json:"time_bank_id"`
	TimeBank   TimeBank   `json:"time_bank" gorm:"foreignKey:TimeBankID"`
	TimeAmount float64    `json:"time_amount"`
	Timestamp  time.Time  `json:"timestamp"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Rating     *float64   `json:"rating,omitempty"`
	Comment    *string    `json:"comment,omitempty"`
}
