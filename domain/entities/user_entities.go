package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID `gorm:"primary_key" json:"id"`
	Email             string    `gorm:"type:VARCHAR(255);NOT NULL;unique" json:"email"`
	Phone             string    `gorm:"type:VARCHAR(255);NOT NULL;unique" json:"phone"`
	Password          string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	ResetPasswordCode string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	OTPCode           string    `gorm:"type:VARCHAR(255);NOT NULL" json:"-"`
	VerifyAt          string    `json:"verify_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Users []User

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

func (u *User) BeforeCreate(scope *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}
