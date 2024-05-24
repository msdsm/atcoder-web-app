package infra

import (
	"fmt"
	"net/http"
)

type IAtcoderUserInfra interface {
	FetchAtcoderId(atcoderId string) error
}
type atcoderUserInfra struct{}

func NewAtcoderUserInfra() IAtcoderUserInfra {
	return &atcoderUserInfra{}
}

func (ui *atcoderUserInfra) FetchAtcoderId(atcoderId string) error {
	url := fmt.Sprintf("https://kenkoooo.com/atcoder/atcoder-api/v3/user/ac_rank?user=%s", atcoderId)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("user ID does not exist")
	}
}
