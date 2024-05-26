package util

import (
	"atcoder-web-app/infra"
	"atcoder-web-app/model"
	"time"
)

type IAtcoderSubmissionUtil interface {
	GetSubmission(atcoderId string, from time.Time) *[]model.SubmissionResponse
}

type atcoderSubmissionUtil struct {
	asi infra.IAtcoderSubmissionInfra
	api infra.IAtcoderProblemInfra
}

func NewAtcoderSubmissionUtil(asi infra.IAtcoderSubmissionInfra, api infra.IAtcoderProblemInfra) IAtcoderSubmissionUtil {
	return &atcoderSubmissionUtil{asi, api}
}

func (asu *atcoderSubmissionUtil) GetSubmission(atcoderId string, from time.Time) *[]model.SubmissionResponse {
	var submissions []model.SubmissionResponse
	res, err := asu.asi.FetchAtcoderSubmission(atcoderId)
	problemIDSet := make(map[string]struct{})
	if err != nil {
		return &submissions
	}

	problems, err := asu.api.FetchAtcoderProblem()
	if err != nil {
		return &submissions
	}
	for _, submission := range *res {
		if submission.Result != "AC" {
			continue
		}
		// ProblemIDがすでに登場している場合は無視
		if _, ok := problemIDSet[submission.ProblemID]; ok {
			continue
		}
		// 提出日がfrom以前なら無視
		if submission.EpochSecond < from.Unix() {
			continue
		}
		// 追加
		submissionResponse := model.SubmissionResponse{
			AtcoderId: atcoderId,
			Time:      time.Unix(submission.EpochSecond, 0).String(),
			Problem:   submission.ProblemID,
			Diff:      problems[submission.ProblemID].Difficulty,
		}
		submissions = append(submissions, submissionResponse)
	}
	return &submissions
}
