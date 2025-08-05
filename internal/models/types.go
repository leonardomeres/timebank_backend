package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Skills    []Skill   `json:"skills" gorm:"foreignKey:CreatedByID"`
	Balance   float64   `json:"balance"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      string    `json:"type"` // admin, user, etc.
}

type Skill struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedByID uint      `json:"created_by_id"`
	Creator     User      `json:"creator" gorm:"foreignKey:CreatedByID"`
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
	TimeAmount float64    `json:"time_amount"`
	Timestamp  time.Time  `json:"timestamp"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Rating     *float64   `json:"rating,omitempty"`
	Comment    *string    `json:"comment,omitempty"`
}
