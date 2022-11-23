package main

import (
	"examples/identity/config"
	"examples/identity/db"
	service "examples/identity/internal/service/jwthelper"
	"examples/identity/internal/usecase"
	"os"

	authRouter "examples/identity/internal/delivery/auth"
	v1 "examples/identity/internal/delivery/v1"
	repo "examples/identity/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	s := config.LoadEnvConfig()
	db := db.InitDatabase(s)

	authRepo := repo.NewAuthRepo(db)
	jwtHandler := service.NewJWTService(s)
	authUsecase := usecase.NewAuthUsecase(authRepo, jwtHandler)
	authRouter := authRouter.NewAuthRouter(authUsecase)
	authRouter.Register(r)

	kahootRepo := repo.NewKahootRepo(db)
	kahootUsecase := usecase.NewKahootUsecase(kahootRepo)

	groupRepo := repo.NewGroupRepo(db)
	groupUsecase := usecase.NewGroupUsecase(groupRepo)
	router := v1.NewRouter(jwtHandler, kahootUsecase, groupUsecase)
	router.Register(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = s.Port
	}
	r.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
