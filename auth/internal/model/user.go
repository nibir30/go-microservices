package model

import "time"

type User struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	FirstName        string     `gorm:"size:255;not null" json:"first_name"`
	LastName         string     `gorm:"size:255;not null" json:"last_name"`
	Email            string     `gorm:"size:255;unique;not null" json:"email"`
	Password         string     `gorm:"size:255;not null" json:"password"`
	Contact          string     `gorm:"size:255;not null" json:"contact"`
	DOB              *time.Time `gorm:"type:date" json:"dob,omitempty"` 
	RegistrationDate *time.Time `gorm:"type:date" json:"registration_date,omitempty"` 
	EmailVerifiedYn  string     `gorm:"not null" json:"email_verified"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
