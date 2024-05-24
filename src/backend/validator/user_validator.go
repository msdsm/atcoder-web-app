package validator

import (
	"atcoder-web-app/infra"
	"atcoder-web-app/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
	AtcoderIdValidate(atcoderId string) error
}
type userValidator struct {
	ui infra.IAtcoderUserInfra
}

func NewUserValidator(ui infra.IAtcoderUserInfra) IUserValidator {
	return &userValidator{ui}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}

func (uv *userValidator) AtcoderIdValidate(atcoderId string) error {
	return uv.ui.FetchAtcoderId(atcoderId)
}
