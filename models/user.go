package models

type User struct {
	ID		  uint `json:"id" gorm:"primary_key"`
	Email     string `json:"email"`
	Password  string  `gorm:"column:password;not null"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	// Bio          string  `gorm:"column:bio;size:1024"`
	// Image        *string `gorm:"column:image"`
}
