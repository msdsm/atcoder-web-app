package repository

import (
	"atcoder-web-app/model"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRivalRepository interface {
	GetAllRivals(rivals *[]model.Rival, userId uuid.UUID) error
	GetRivalById(rival *model.Rival, userId uuid.UUID, ID uuid.UUID) error
	CreateRival(rival *model.Rival) error
	// UpdateRival(rival *model.Rival, userId uuid.UUID, ID uuid.UUID) error
	DeleteRival(userId uuid.UUID, ID uuid.UUID) error
}

type rivalRepository struct {
	db *gorm.DB
}

func NewRivalRepository(db *gorm.DB) IRivalRepository {
	return &rivalRepository{db}
}

func (rr *rivalRepository) GetAllRivals(rivals *[]model.Rival, userId uuid.UUID) error {
	if err := rr.db.Joins("User").Where("user_id=?", userId).Find(rivals).Error; err != nil {
		return err
	}
	return nil
}

func (rr *rivalRepository) GetRivalById(rival *model.Rival, userId uuid.UUID, ID uuid.UUID) error {
	if err := rr.db.Joins("User").Where("user_id=?", userId).First(rival, ID).Error; err != nil {
		return err
	}
	return nil
}

func (rr *rivalRepository) CreateRival(rival *model.Rival) error {
	if err := rr.db.Create(rival).Error; err != nil {
		return err
	}
	return nil
}

/*
func (rr *rivalRepository) UpdateRival(rival *model.Rival, userId uuid.UUID, ID uuid.UUID) error {
	result := rr.db.Model(rival).Where("id=? AND user_id=?", ID, userId).Update("rival_atcoder_id", rival.RivalAtcoderId)
	if result.Error != nil {
		return result.Error
	}
	// 更新されかたどうか
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
*/

func (rr *rivalRepository) DeleteRival(userId uuid.UUID, ID uuid.UUID) error {
	result := rr.db.Where("id=? AND user_id=?", ID, userId).Delete(&model.Rival{})
	if result.Error != nil {
		return result.Error
	}
	// 削除されたかどうか
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
