package auth_token

import "time"

type Token struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	Token     []byte    `gorm:"notnull" json:"token"`
	User_id   uint      `gorm:"notnull" json:"-"`
	ExpiresAt time.Time `gorm:"notnull" json:"-"`
}
