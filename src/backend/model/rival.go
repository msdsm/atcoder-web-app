package model

import "github.com/google/uuid"

type Rival struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	RivalAtcoderId string    `json:"rival" gorm:"not null"`
	User           User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId         uuid.UUID `json:"user_id" gorm:"not null"`
}

type RivalResponse struct {
	ID             uuid.UUID `json:"id"`
	RivalAtcoderId string    `json:"rival"`
}
