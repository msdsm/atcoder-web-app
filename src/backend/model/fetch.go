package model

import "time"

type FetchSubmission struct {
	ID            int64   `json:"id"`
	EpochSecond   int64   `json:"epoch_second"`
	ProblemID     string  `json:"problem_id"`
	ContestID     string  `json:"contest_id"`
	UserID        string  `json:"user_id"`
	Language      string  `json:"language"`
	Point         float64 `json:"point"`
	Length        int     `json:"length"`
	Result        string  `json:"result"`
	ExecutionTime int     `json:"execution_time"`
}

type FetchProblem struct {
	Slope            float64 `json:"slope"`
	Intercept        float64 `json:"intercept"`
	Variance         float64 `json:"variance"`
	Difficulty       int     `json:"difficulty"`
	Discrimination   float64 `json:"discrimination"`
	IRTLogLikelihood float64 `json:"irt_loglikelihood"`
	IRTUsers         int     `json:"irt_users"`
	IsExperimental   bool    `json:"is_experimental"`
}

type FetchHistory struct {
	IsRated           bool      `json:"IsRated"`
	Place             int       `json:"Place"`
	OldRating         int       `json:"OldRating"`
	NewRating         int       `json:"NewRating"`
	Performance       int       `json:"Performance"`
	InnerPerformance  int       `json:"InnerPerformance"`
	ContestScreenName string    `json:"ContestScreenName"`
	ContestName       string    `json:"ContestName"`
	ContestNameEn     string    `json:"ContestNameEn"`
	EndTime           time.Time `json:"EndTime"`
}
