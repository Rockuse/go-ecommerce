package user

import "time"

type User struct {
	id           int
	Username     string `json:"usename" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	IsDeleted    int
	CreatedDate  time.Time
	ModifiedDate time.Time
	CreatedBy    int
	ModifiedBy   int
}

type EmailInput struct {
	Email string `json:"email" binding:"required"`
}
