package models

import (
	"time"
)

type User struct {
	ID              string     `json:"id" db:"id"`
	Username        string     `json:"username" db:"username"`
	Email           string     `json:"email" db:"email"`
	PasswordHash    string     `json:"-" db:"password_hash"`
	FullName        *string    `json:"full_name" db:"full_name"`
	BirthDate       *time.Time `json:"birth_date" db:"birth_date"`
	Phone           *string    `json:"phone" db:"phone"`
	ProfileImageURL *string    `json:"profile_image_url" db:"profile_image_url"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	IsActive        bool       `json:"is_active" db:"is_active"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"max=100"`
}

type UpdateUserRequest struct {
	FullName        *string    `json:"full_name" binding:"omitempty,max=100"`
	BirthDate       *time.Time `json:"birth_date"`
	Phone           *string    `json:"phone" binding:"omitempty,max=20"`
	ProfileImageURL *string    `json:"profile_image_url"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type EmergencyContact struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	Name         string    `json:"name" db:"name"`
	Phone        string    `json:"phone" db:"phone"`
	Relationship *string   `json:"relationship" db:"relationship"`
	IsPrimary    bool      `json:"is_primary" db:"is_primary"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateContactRequest struct {
	Name         string  `json:"name" binding:"required,max=100"`
	Phone        string  `json:"phone" binding:"required,max=20"`
	Relationship *string `json:"relationship" binding:"omitempty,max=50"`
	IsPrimary    bool    `json:"is_primary"`
}

type UpdateContactRequest struct {
	Name         *string `json:"name" binding:"omitempty,max=100"`
	Phone        *string `json:"phone" binding:"omitempty,max=20"`
	Relationship *string `json:"relationship" binding:"omitempty,max=50"`
	IsPrimary    *bool   `json:"is_primary"`
}
