package models

import "time"

type Student struct {
	Id          int       `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name" binding:"required,min=2,max=30"`
	LastName    string    `json:"last_name" db:"last_name" binding:"required,min=2,max=30"`
	UserName    string    `json:"username" db:"user_name" binding:"required,min=2,max=30"`
	Email       string    `json:"email" db:"email" binding:"required,email"`
	PhoneNumber string    `json:"phone_number"  db:"phone_number" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreateStudent struct {
	FirstName   string `json:"first_name" db:"first_name" binding:"required,min=2,max=30"`
	LastName    string `json:"last_name" db:"last_name" binding:"required,min=2,max=30"`
	UserName    string `json:"username" db:"user_name" binding:"required,min=2,max=30"`
	Email       string `json:"email" db:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" db:"phone_number" binding:"required"`
}
