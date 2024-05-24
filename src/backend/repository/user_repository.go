package repository

import (
	"atcoder-web-app/model"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateAtcoderId(user *model.User, userId uuid.UUID, atcoderId string) error
}
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateAtcoderId(user *model.User, userId uuid.UUID, atcoderId string) error {
	result := ur.db.Model(user).Where("id=?", userId).Update("atcoder_id", atcoderId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
