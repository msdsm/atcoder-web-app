package usecase

import (
	"atcoder-web-app/model"
	"atcoder-web-app/repository"
	"atcoder-web-app/util"
	"atcoder-web-app/validator"

	"github.com/google/uuid"
	"github.com/jinzhu/now"
)

type IRivalUsecase interface {
	GetAllRivals(userId uuid.UUID) ([]model.RivalResponse, error)
	GetRivalById(userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error)
	CreateRival(rival model.Rival) (model.RivalResponse, error)
	// UpdateRival(rival model.Rival, userId uuid.UUID, ID uuid.UUID) (model.RivalResponse, error)
	DeleteRival(userId uuid.UUID, ID uuid.UUID) error
	GetSubmission(userId uuid.UUID) ([]model.SubmissionResponse, error)
	GetTable(userId uuid.UUID) ([]model.TableResponse, error)
}

type rivalUsecase struct {
	ur  repository.IUserRepository
	rr  repository.IRivalRepository
	rv  validator.IRivalValidator
	asu util.IAtcoderSubmissionUtil
	auu util.IAtcoderUserUtil
}

func NewRivalUsecase(ur repository.IUserRepository, rr repository.IRivalRepository, rv validator.IRivalValidator, asu util.IAtcoderSubmissionUtil, auu util.IAtcoderUserUtil) IRivalUsecase {
	return &rivalUsecase{ur, rr, rv, asu, auu}
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

func (ru *rivalUsecase) GetSubmission(userId uuid.UUID) ([]model.SubmissionResponse, error) {
	var submissions []model.SubmissionResponse
	rivalResponse, err := ru.GetAllRivals(userId)
	if err != nil {
		return submissions, err
	}
	// 今日の提出だけ取得
	today := now.BeginningOfDay()

	// 自分
	var storedUser model.User
	if err := ru.ur.GetUserById(&storedUser, userId); err != nil {
		return submissions, err
	}
	atcoderId := storedUser.AtcoderId
	res := ru.asu.GetSubmission(atcoderId, today)
	submissions = append(submissions, (*res)...)

	// ライバル
	for _, rival := range rivalResponse {
		res := ru.asu.GetSubmission(rival.RivalAtcoderId, today)
		submissions = append(submissions, (*res)...)
	}
	return submissions, nil
}

func (ru *rivalUsecase) GetTable(userId uuid.UUID) ([]model.TableResponse, error) {
	var tables []model.TableResponse
	rivalResponse, err := ru.GetAllRivals(userId)
	if err != nil {
		return tables, err
	}

	// 自分追加
	var storedUser model.User
	if err := ru.ur.GetUserById(&storedUser, userId); err != nil {
		return tables, err
	}
	atcoderId := storedUser.AtcoderId
	rating, err := ru.auu.GetRating(atcoderId)
	if err != nil {
		return tables, err
	}
	streak, err := ru.auu.GetRating(atcoderId)
	if err != nil {
		return tables, err
	}
	tables = append(tables, model.TableResponse{
		AtcoderId: atcoderId,
		Rating:    rating,
		Streak:    streak,
	})

	// ライバル追加
	for _, rival := range rivalResponse {
		rating, err := ru.auu.GetRating(rival.RivalAtcoderId)
		if err != nil {
			return nil, err
		}
		streak, err := ru.auu.GetStreak(rival.RivalAtcoderId)
		if err != nil {
			return nil, err
		}
		table := model.TableResponse{
			AtcoderId: rival.RivalAtcoderId,
			Rating:    rating,
			Streak:    streak,
		}
		tables = append(tables, table)
	}
	return tables, nil
}
