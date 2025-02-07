package model

import "time"

type User struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Username        string     `gorm:"size:255;unique;not null" json:"username"`
	FirstName       string     `gorm:"size:255;not null" json:"firstName"`
	LastName        string     `gorm:"size:255;not null" json:"lastName"`
	Email           string     `gorm:"size:255;unique;not null" json:"email"`
	Password        string     `gorm:"size:511;not null" json:"-"` // Ignore in JSON
	Contact         string     `gorm:"size:255;not null" json:"contact"`
	DOB             *time.Time `gorm:"type:date" json:"dob,omitempty"`
	EmailVerifiedYn string     `gorm:"not null" json:"emailVerifiedYn"`
	CreatedAt       *time.Time `json:"-"`              // Ignore in JSON
	UpdatedAt       *time.Time `json:"-"`              // Ignore in JSON
	DeletedAt       *time.Time `gorm:"index" json:"-"` // Ignore in JSON
}
