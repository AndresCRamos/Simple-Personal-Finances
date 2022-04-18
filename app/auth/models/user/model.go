package auth_user

import "github.com/emvi/null"

type User struct {
	ID       uint        `gorm:"primarykey" json:"-"`
	Name     null.String `gorm:"notnull" json:"name"`
	LastName null.String `gorm:"notnull" json:"last_name"`
	Email    null.String `gorm:"notnull;unique" json:"email"`
	Password []byte      `gorm:"notnull" json:"-"`
}
