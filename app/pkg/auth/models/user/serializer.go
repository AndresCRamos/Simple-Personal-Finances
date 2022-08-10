package auth_user

import (
	"github.com/emvi/null"
	"golang.org/x/crypto/bcrypt"
)

type UserGet User

type UserRegister struct {
	ID       uint        `gorm:"primarykey" json:"-"`
	Name     null.String `gorm:"notnull" json:"name"`
	LastName null.String `gorm:"notnull" json:"last_name"`
	Email    null.String `gorm:"notnull" json:"email" validation:"type:email"`
	Password null.String `gorm:"notnull" json:"password"`
}

func (ig *UserRegister) Parse() *User {
	pass, _ := bcrypt.GenerateFromPassword([]byte(ig.Password.String), 12)
	return &User{
		ID:       ig.ID,
		Name:     ig.Name,
		LastName: ig.LastName,
		Email:    ig.Email,
		Password: pass,
	}
}
