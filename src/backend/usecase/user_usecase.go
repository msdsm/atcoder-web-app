package usecase

import (
	"atcoder-web-app/model"
	"atcoder-web-app/repository"
	"atcoder-web-app/util"
	"atcoder-web-app/validator"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	Update(userId uuid.UUID, atcoderId string) (model.UserResponse, error)
	GetAtcoderId(userId uuid.UUID) (string, error)
}

type userUsecase struct {
	ur  repository.IUserRepository
	uv  validator.IUserValidator
	auu util.IAtcoderUserUtil
	asu util.IAtcoderSubmissionUtil
}

// コンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator, auu util.IAtcoderUserUtil, asu util.IAtcoderSubmissionUtil) IUserUsecase {
	return &userUsecase{ur, uv, auu, asu}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		ID:        uuid.New(),
		Email:     user.Email,
		Password:  string(hash),
		AtcoderId: user.AtcoderId,
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		AtcoderId: newUser.AtcoderId,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) Update(userId uuid.UUID, atcoderId string) (model.UserResponse, error) {
	storedUser := model.User{}
	if err := uu.uv.AtcoderIdValidate(atcoderId); err != nil {
		return model.UserResponse{}, err
	}
	if err := uu.ur.UpdateAtcoderId(&storedUser, userId, atcoderId); err != nil {
		return model.UserResponse{}, err
	}
	fmt.Println(storedUser)
	resUser := model.UserResponse{
		ID:        storedUser.ID,
		Email:     storedUser.Email,
		AtcoderId: storedUser.AtcoderId,
	}
	return resUser, nil
}

func (uu *userUsecase) GetAtcoderId(userId uuid.UUID) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserById(&storedUser, userId); err != nil {
		return "", err
	}
	return storedUser.AtcoderId, nil
}
