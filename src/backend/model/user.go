package model

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"primalyKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	AtcoderId string    `json:"atcoder_id"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	AtcoderId string    `json:"atcoder_id"`
}
