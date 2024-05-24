package util

import (
	"atcoder-web-app/infra"
	"atcoder-web-app/model"
	"time"
)

type IAtcoderUserUtil interface {
	GetRating(atcoderId string) (int, error)
	GetStreak(atcoderId string) (int, error)
}

type atcoderUserUtil struct {
	asi infra.IAtcoderSubmissionInfra
	ahi infra.IAtcoderHistoryInfra
}

func NewAtcoderUserUtil(asi infra.IAtcoderSubmissionInfra, ahi infra.IAtcoderHistoryInfra) IAtcoderUserUtil {
	return &atcoderUserUtil{asi, ahi}
}

func (auu *atcoderUserUtil) GetStreak(atcoderId string) (int, error) {
	allSubmissions, err := auu.asi.FetchAtcoderSubmission(atcoderId)
	if err != nil {
		return 0, err
	}
	return getStreak(allSubmissions), nil
}

func (auu *atcoderUserUtil) GetRating(atcoderId string) (int, error) {
	history, err := auu.ahi.FetchAtcoderHistory(atcoderId)
	if err != nil {
		return 0, err
	}
	return getRating(history), nil
}

func getRating(history *[]model.FetchHistory) int {
	return (*history)[len(*history)-1].NewRating
}

func getStreak(allSubmissions *[]model.FetchSubmission) int {
	currentStreak := 0
	lastDate := time.Unix(0, 0)
	problemIDSet := make(map[string]struct{})

	for _, submission := range *allSubmissions {
		// 問題を正解していない場合は無視
		if submission.Result != "AC" {
			continue
		}

		// ProblemIDがすでに登場している場合は無視
		if _, ok := problemIDSet[submission.ProblemID]; ok {
			continue
		}

		//ここまできたら最低でもstreakは1
		submissionDate := time.Unix(submission.EpochSecond, 0)

		// すでにその日カウント
		if sameDate(submissionDate, lastDate) {
			continue
		}

		// 直近の正解日から1日後の日付を取得
		nextdate := lastDate.AddDate(0, 0, 1)

		// 直近の正解日と直前の正解日の日付が連続しているかどうかを確認
		if sameDate(nextdate, submissionDate) {
			currentStreak++
		} else {
			//非連続
			currentStreak = 1
		}

		// 解いた問題追加
		problemIDSet[submission.ProblemID] = struct{}{}
		// 最後の提出日更新
		lastDate = submissionDate
	}

	// 今日まで続いているかどうか
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	if sameDate(lastDate, today) {
		return currentStreak
	} else if sameDate(lastDate, yesterday) {
		return currentStreak
	} else {
		currentStreak = 0
	}
	return currentStreak
}

// 日付が等しければtrue(日時の時刻無視)
func sameDate(date1, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day()
}
