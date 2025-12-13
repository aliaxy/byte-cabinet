package model

import (
	"time"
)

// User represents the admin user of the blog
type User struct {
	ID           int64     `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"` // Never expose in JSON
	DisplayName  string    `db:"display_name" json:"display_name"`
	Avatar       string    `db:"avatar" json:"avatar,omitempty"`
	Bio          string    `db:"bio" json:"bio,omitempty"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// UserLogin represents login request payload
type UserLogin struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserUpdate represents update profile request payload
type UserUpdate struct {
	DisplayName string `json:"display_name" validate:"omitempty,max=100"`
	Email       string `json:"email" validate:"omitempty,email"`
	Avatar      string `json:"avatar" validate:"omitempty,url"`
	Bio         string `json:"bio" validate:"omitempty,max=500"`
}

// PasswordChange represents password change request payload
type PasswordChange struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=72"`
}

// UserResponse represents user data in API responses
type UserResponse struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		Avatar:      u.Avatar,
		Bio:         u.Bio,
		CreatedAt:   u.CreatedAt,
	}
}
