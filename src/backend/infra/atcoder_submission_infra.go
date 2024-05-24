package infra

import (
	"atcoder-web-app/model"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type IAtcoderSubmissionInfra interface {
	FetchAtcoderSubmission(atcoderId string) (*[]model.FetchSubmission, error)
}
type atcoderSubmissionInfra struct{}

func NewAtcoderSubmissionInfra() IAtcoderSubmissionInfra {
	return &atcoderSubmissionInfra{}
}

// ユーザーIDから提出全て取得
func (asi *atcoderSubmissionInfra) FetchAtcoderSubmission(atcoderId string) (*[]model.FetchSubmission, error) {
	fromSecond := 0
	baseURL := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/submissions?user=%s&from_second=%d"
	var allSubmissions []model.FetchSubmission
	for {
		fmt.Println(fmt.Sprintf(baseURL, atcoderId, fromSecond))
		submissions, err := fetchSubmissions(fmt.Sprintf(baseURL, atcoderId, fromSecond))
		if err != nil {
			return nil, err
		}
		if len(*submissions) == 0 {
			break
		}
		allSubmissions = append(allSubmissions, *submissions...)
		lastSubmission := (*submissions)[len(*submissions)-1]
		fromSecond = int(lastSubmission.EpochSecond) + 1
	}
	// EpochSecondをキーとしてソート
	sort.Slice(allSubmissions, func(i, j int) bool {
		return allSubmissions[i].EpochSecond < allSubmissions[j].EpochSecond
	})
	return &allSubmissions, nil
}

// ユーザーIDから提出取得(上限500)
func fetchSubmissions(url string) (*[]model.FetchSubmission, error) {
	var submissions []model.FetchSubmission
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GETリクエストの送信中にエラーが発生しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ネットワークレスポンスがOKではありませんでした")
	}

	err = json.NewDecoder(resp.Body).Decode(&submissions)
	if err != nil {
		return nil, fmt.Errorf("JSONの解析中にエラーが発生しました: %v", err)
	}

	return &submissions, nil
}
