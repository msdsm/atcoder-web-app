package infra

import (
	"atcoder-web-app/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type IAtcoderHistoryInfra interface {
	FetchAtcoderHistory(atcoderId string) (*[]model.FetchHistory, error)
}
type atcoderHistoryInfra struct{}

func NewAtcoderHistoryInfra() IAtcoderHistoryInfra {
	return &atcoderHistoryInfra{}
}

// ユーザーIDから提出全て取得
func (ahi *atcoderHistoryInfra) FetchAtcoderHistory(atcoderId string) (*[]model.FetchHistory, error) {
	baseURL := "https://kenkoooo.com/atcoder/proxy/users/%s/history/json"
	res, err := http.Get(fmt.Sprintf(baseURL, atcoderId))
	if err != nil {
		return nil, fmt.Errorf("HTTP GETリクエストの送信中にエラーが発生しました: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ネットワークレスポンスがOKではありませんでした")
	}

	var history []model.FetchHistory
	err = json.NewDecoder(res.Body).Decode(&history)
	if err != nil {
		return nil, fmt.Errorf("JSONの解析中にエラーが発生しました: %v", err)
	}

	return &history, nil
}
