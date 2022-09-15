package models

type Email struct {
	Email string `form:"email" binding:"required"`
}
