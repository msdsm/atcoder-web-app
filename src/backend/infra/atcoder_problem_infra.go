package infra

import (
	"atcoder-web-app/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type IAtcoderProblemInfra interface {
	FetchAtcoderProblem() (map[string]model.FetchProblem, error)
}
type atcoderProblemInfra struct{}

func NewAtcoderProblemInfra() IAtcoderProblemInfra {
	return &atcoderProblemInfra{}
}

// 問題すべて取得
func (ahi *atcoderProblemInfra) FetchAtcoderProblem() (map[string]model.FetchProblem, error) {
	url := "https://kenkoooo.com/atcoder/resources/problem-models.json"
	var problems map[string]model.FetchProblem
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GETリクエストの送信中にエラーが発生しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ネットワークレスポンスがOKではありませんでした")
	}

	err = json.NewDecoder(resp.Body).Decode(&problems)
	if err != nil {
		return nil, fmt.Errorf("JSONの解析中にエラーが発生しました: %v", err)
	}

	return problems, nil
}
