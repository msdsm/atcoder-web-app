package model

import "time"

type TableResponse struct {
	AtcoderId string `json:"atcoder_id"`
	Rating    int    `json:"rating"`
	Streak    int    `json:"streak"`
}

type SubmissionResponse struct {
	AtcoderId string    `json:"atcoder_id"`
	Time      time.Time `json:"time"`
	Problem   string    `json:"problem"`
	Diff      int       `json:"diff"`
}
