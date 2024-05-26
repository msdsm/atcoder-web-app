package main

import (
	"atcoder-web-app/controller"
	"atcoder-web-app/db"
	"atcoder-web-app/infra"
	"atcoder-web-app/repository"
	"atcoder-web-app/router"
	"atcoder-web-app/usecase"
	"atcoder-web-app/util"
	"atcoder-web-app/validator"
)

func main() {
	// db
	db := db.NewDB()

	// repository
	userRepository := repository.NewUserRepository(db)
	rivalRepository := repository.NewRivalRepository(db)

	// infra
	atcoderHistoryInfra := infra.NewAtcoderHistoryInfra()
	atcoderProblemInfra := infra.NewAtcoderProblemInfra()
	atcoderSubmissionInfra := infra.NewAtcoderSubmissionInfra()
	atcoderUserInfra := infra.NewAtcoderUserInfra()

	// validator
	userValidator := validator.NewUserValidator(atcoderUserInfra)
	rivalValidator := validator.NewRivalValidator(atcoderUserInfra)

	// util
	atcoderSubmissionUtil := util.NewAtcoderSubmissionUtil(atcoderSubmissionInfra, atcoderProblemInfra)
	atcoderUserUtil := util.NewAtcoderUserUtil(atcoderSubmissionInfra, atcoderHistoryInfra)

	// usecase
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		userValidator,
		atcoderUserUtil,
		atcoderSubmissionUtil,
	)
	rivalUsecase := usecase.NewRivalUsecase(
		userRepository,
		rivalRepository,
		rivalValidator,
		atcoderSubmissionUtil,
		atcoderUserUtil,
	)

	// controller
	userController := controller.NewUserController(userUsecase)
	rivalController := controller.NewRivalController(rivalUsecase)

	// router
	e := router.NewRouter(userController, rivalController)

	// port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
