package model

type User struct {
	Id    uint   `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email" binding:"required"`
	//Password          string `json:"password,omitempty" binding:"required"`
	EncryptedPassword string `json:"password"`
}
