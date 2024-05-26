package controller

import (
	"atcoder-web-app/model"
	"atcoder-web-app/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IRivalController interface {
	CreateRival(c echo.Context) error
	DeleteRival(c echo.Context) error
	GetTable(c echo.Context) error
	GetSubmission(c echo.Context) error
}

type rivalController struct {
	ru usecase.IRivalUsecase
}

func NewRivalController(ru usecase.IRivalUsecase) IRivalController {
	return &rivalController{ru}
}

func (rc *rivalController) CreateRival(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	rival := model.Rival{}
	if err := c.Bind(&rival); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	rival.UserId = userId.(uuid.UUID)
	rivalRes, err := rc.ru.CreateRival(rival)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, rivalRes)
}

func (rc *rivalController) DeleteRival(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(uuid.UUID)
	idstr := c.Param("id")
	id, _ := uuid.Parse(idstr)
	if err := rc.ru.DeleteRival(userId, id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (rc *rivalController) GetTable(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(uuid.UUID)
	res, err := rc.ru.GetTable(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func (rc *rivalController) GetSubmission(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(uuid.UUID)
	res, err := rc.ru.GetSubmission(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
