//go:build wireinject
// +build wireinject

package di

import (
	"examples/kahootee/config"
	"examples/kahootee/internal/delivery/http"
	"examples/kahootee/internal/repository"
	service "examples/kahootee/internal/service/jwthelper"
	"examples/kahootee/internal/usecase"
	"examples/kahootee/pkg/httpserver"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var useCaseSet = wire.NewSet(config.NewConfig, provideGormDB, provideKahootUseCase, provideKahootRepo, provideGroupUseCase, provideGroupRepo, provideAuthUseCase, provideAuthRepo, provideJWTService,provideConfig)

func InitializeHttpServer() (*httpserver.Server, func(), error) {
	panic(wire.Build(
		useCaseSet,
		gin.New,
		provideHttpServer,
		http.NewRouter,
	))
}

// func InitializeGRPCServer() (*grpcserver.GRPCServer, func(), error) {
// 	panic(wire.Build(
// 		useCaseSet,
// 		grpc.NewServer,
// 		provideGRPCServerOptions,
// 		provideGRPCServer,
// 		provideGRPCCustomerService,
// 	))
// }

// func provideGRPCCustomerService(u usecase.Customer) customer.CustomerServiceServer {
// 	return grpcDelivery.NewCustomerService(u)
// }

// func provideGRPCServerOptions() []grpc.ServerOption {
// 	return nil
// }

// func provideGRPCServer(cfg *config.Config, server *grpc.Server, delivery customer.CustomerServiceServer) *grpcserver.GRPCServer {
// 	customer.RegisterCustomerServiceServer(server, delivery)
// 	return grpcserver.New(server, cfg.GRPC.Address)
// }

func provideKahootRepo(db *gorm.DB) usecase.KahootRepo {
	return repo.NewKahootRepo(db)
}

func provideKahootUseCase(k usecase.KahootRepo) usecase.KahootUsecase {
	return usecase.NewKahootUsecase(k)
}

func provideGroupRepo(db *gorm.DB) usecase.GroupRepo {
	return repo.NewGroupRepo(db)
}

func provideGroupUseCase(g usecase.GroupRepo) usecase.GroupUsecase {
	return usecase.NewGroupUsecase(g)
}

func provideAuthRepo(db *gorm.DB) usecase.AuthRepo {
	return repo.NewAuthRepo(db)
}

func provideAuthUseCase(g usecase.AuthRepo, jwtService service.JWTHelper) usecase.AuthUsecase {
	return usecase.NewAuthUsecase(g, jwtService)
}
func provideJWTService(s *config.Specification) service.JWTHelper {
	return service.NewJWTService(s)
}
func provideConfig() *config.Specification {
	return config.LoadEnvConfig()
}

func provideGormDB(cfg *config.Config) (*gorm.DB, func(), error) {
	db, err := gorm.Open(postgres.Open(cfg.PG.URL), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		conn, err := db.DB()
		if err != nil {
			log.Printf("failed to get db connection, %v", err)
			return
		}
		conn.Close()
	}, nil
}

func provideHttpServer(router *http.Router, handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	router.Register()
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}
