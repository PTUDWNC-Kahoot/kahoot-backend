package main

import (
	"examples/identity/config"
	"examples/identity/db"
	service "examples/identity/internal/service/jwthelper"
	"examples/identity/internal/usecase"

	authRouter "examples/identity/internal/delivery/auth"
	v1 "examples/identity/internal/delivery/v1"
	repo "examples/identity/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	s := config.LoadEnvConfig()
	db := db.InitDatabase(s)

	authRepo := repo.NewAuthRepo(db)
	jwtHandler := service.NewJWTService(s)
	authUsecase := usecase.NewAuthUsecase(authRepo, jwtHandler)
	authRouter := authRouter.NewAuthRouter(authUsecase)
	authRouter.Register(r)

	kahootRepo := repo.NewKahootRepo(db)
	kahootUsecase := usecase.NewKahootUsecase(kahootRepo)

	groupRepo := repo.NewKahootRepo(db)
	groupUsecase := usecase.NewGroupUsecase(groupRepo)
	router := v1.NewRouter(jwtHandler, kahootUsecase, groupUsecase)
	router.Register(r)

	r.Run(":" + s.Port)
}
