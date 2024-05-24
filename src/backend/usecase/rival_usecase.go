package usecase

import (
	"atcoder-web-app/model"
	"atcoder-web-app/repository"
	"atcoder-web-app/validator"

	"github.com/google/uuid"
)

type IRivalUsecase interface {
	GetAllRivals(userId uuid.UUID) ([]model.RivalResponse, error)
	GetRivalById(userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error)
	CreateRival(rival model.Rival) (model.RivalResponse, error)
	// UpdateRival(rival model.Rival, userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error)
	DeleteRival(userId uuid.UUID, ID uuid.UUID) error
}

type rivalUsecase struct {
	rr repository.IRivalRepository
	rv validator.IRivalValidator
}

func NewTaskUsecase(rr repository.IRivalRepository, rv validator.IRivalValidator) IRivalUsecase {
	return &rivalUsecase{rr, rv}
}

func (ru *rivalUsecase) GetAllRivals(userId uuid.UUID) ([]model.RivalResponse, error) {
	rivals := []model.Rival{}
	if err := ru.rr.GetAllRivals(&rivals, userId); err != nil {
		return nil, err
	}
	resRivals := []model.RivalResponse{}
	for _, v := range rivals {
		r := model.RivalResponse{
			ID:             v.ID,
			RivalAtcoderId: v.RivalAtcoderId,
		}
		resRivals = append(resRivals, r)
	}
	return resRivals, nil
}

func (ru *rivalUsecase) GetRivalById(userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error) {
	rival := model.Rival{}
	if err := ru.rr.GetRivalById(&rival, userId, ID); err != nil {
		return model.RivalResponse{}, err
	}
	resRival := model.RivalResponse{
		ID:             rival.ID,
		RivalAtcoderId: rival.RivalAtcoderId,
	}
	return resRival, nil
}

func (ru *rivalUsecase) CreateRival(rival model.Rival) (model.RivalResponse, error) {
	if err := ru.rv.RivalValidate(rival); err != nil {
		return model.RivalResponse{}, err
	}
	if err := ru.rr.CreateRival(&rival); err != nil {
		return model.RivalResponse{}, err
	}
	resRival := model.RivalResponse{
		ID:             rival.ID,
		RivalAtcoderId: rival.RivalAtcoderId,
	}
	return resRival, nil
}

// func (ru *rivalUsecase) UpdateRival(rival model.Rival, userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error)

func (ru *rivalUsecase) DeleteRival(userId uuid.UUID, ID uuid.UUID) error {
	if err := ru.rr.DeleteRival(userId, ID); err != nil {
		return err
	}
	return nil
}
