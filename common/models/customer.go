package models

import (
	"gorm.io/gorm"
)

//import "gorm.io/gorm"

type Customer struct {
	gorm.Model        // adds ID, created_at etc. `
	FirstName  string `json:"first_name" gorm:"not null"`
	LastName   string `json:"last_name" gorm:"not null"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile" gorm:"unique"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Status     string `json:"status" gorm:"default:ACTIVE"`
	Avatar     string `json:"avatar"`
}
