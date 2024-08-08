package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4;primaryKey;unique;AutoIncrement:false"`
	Name     string    `json:"name" gorm:"type:string"`
	Email    string    `json:"email" gorm:"type:string;index"`
	Password string    `json:"password" gorm:"type:string;index"`
	Phone    string    `json:"phone" gorm:"type:string;index"`
}
