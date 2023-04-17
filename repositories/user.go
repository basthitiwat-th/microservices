package repositories

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `db:"username" gorm:"unique;not null"`
	FirstName string `db:"first_name" gorm:"not null"`
	LastName  string `db:"last_name" gorm:"not null"`
	Phone     string `db:"phone" gorm:"unique;not null"`
	Email     string `db:"email" gorm:"unique;not null"`
	Password  string `db:"password" gorm:"not null"`
}
