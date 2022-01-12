package model

type User struct {
	Id                uint   `json:"id" binding:"required"`
	Name              string `json:"name"`
	Email             string `json:"email" binding:"required" gorm:"unique" `
	EncryptedPassword string `json:"password"`
}
