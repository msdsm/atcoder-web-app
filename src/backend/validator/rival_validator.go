package validator

import (
	"atcoder-web-app/infra"
	"atcoder-web-app/model"
)

type IRivalValidator interface {
	RivalValidate(rival model.Rival) error
}
type rivalValidator struct {
	ui infra.IAtcoderUserInfra
}

func NewRivalValidator(ui infra.IAtcoderUserInfra) IRivalValidator {
	return &rivalValidator{ui}
}

func (rv *rivalValidator) RivalValidate(rival model.Rival) error {
	return rv.ui.FetchAtcoderId(rival.RivalAtcoderId)
}
