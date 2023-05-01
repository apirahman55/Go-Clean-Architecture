// gorm.Model definition
package entities

import "time"

type User struct {
	UUID              string    `gorm:"type:VARCHAR(255);NOT NULL;primary_key" json:"uuid"`
	Email             string    `gorm:"type:VARCHAR(255);NOT NULL;unique" json:"email"`
	Phone             string    `gorm:"type:VARCHAR(255);NOT NULL;unique" json:"phone"`
	Password          string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	ResetPasswordCode string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	OTPCode           string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	VerifyAt          time.Time `json:"verify_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Users []User
