package model

import "github.com/google/uuid"

type TableResponse struct {
	Id        uuid.UUID `json:"id"`
	AtcoderId string    `json:"atcoder_id"`
	Rating    int       `json:"rating"`
	Streak    int       `json:"streak"`
}

type SubmissionResponse struct {
	AtcoderId string `json:"atcoder_id"`
	Time      string `json:"time"`
	Problem   string `json:"problem"`
	Diff      int    `json:"diff"`
}
