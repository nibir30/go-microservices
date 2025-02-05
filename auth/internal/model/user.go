package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Contact   string    `gorm:"size:255;not null" json:"contact"`
	DOB       time.Time `gorm:"not null" json:"dob"`
	ImageID   uint      `gorm:"not null" json:"image_id"`
	RegistrationDate time.Time `gorm:"not null" json:"registration_date"`
	EmailVerifiedYn  string      `gorm:"not null" json:"email_verified"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
