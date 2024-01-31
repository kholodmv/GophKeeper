package models

import "gorm.io/gorm"

// Auth - registration/authentication data
type Auth struct {
	Login    string `json:"login"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// User - user model
type User struct {
	gorm.Model
	Login    string `gorm:"size:255;not null;unique"`
	Password []byte `gorm:"size:255;not null"`
}

// Secret - stored secret model
type Secret struct {
	gorm.Model `json:"-"`
	UserID     uint   `gorm:"not null;index:,unique,composite:uid" json:"-"`
	Title      string `gorm:"size:255;not null;index:,unique,composite:uid" json:"title"`
	Comment    string `gorm:"size:255;" json:"comment"`

	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Text     string `json:"text,omitempty"`
	Number   string `json:"number,omitempty"`
	Date     string `json:"date,omitempty"`
	Cvv      string `json:"cvv,omitempty"`
	Binary   []byte `json:"binary,omitempty"`
	FilePath string `json:"path,omitempty"`
}
