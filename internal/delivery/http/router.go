// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"examples/kahootee/internal/delivery/http/auth"
	v1 "examples/kahootee/internal/delivery/http/v1"
	service "examples/kahootee/internal/service/jwthelper"
	"examples/kahootee/internal/usecase"
)

type Router struct {
	handler *gin.Engine
	j       service.JWTHelper
	a       usecase.AuthUsecase
	k       usecase.KahootUsecase
	g       usecase.GroupUsecase
	p       usecase.User
}

func (r *Router) Register() {
	// Options
	r.handler.Use(gin.Logger())
	r.handler.Use(gin.Recovery())
	r.handler.Use(CORSMiddleware())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	r.handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	r.handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	r.handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	v1.NewRouter(r.handler.Group("/v1"), r.j, r.k, r.g, r.p)
	auth.NewAuthRouter(r.handler.Group("/"), r.a)
}

func NewRouter(handler *gin.Engine, jwtHelper service.JWTHelper, a usecase.AuthUsecase, k usecase.KahootUsecase, g usecase.GroupUsecase, p usecase.User) *Router {
	return &Router{
		handler: handler,
		j:       jwtHelper,
		a:       a,
		k:       k,
		g:       g,
		p:       p,
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-TB-Access-Token, accept, origin, Cache-Control, X-Requested-With, Content-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
