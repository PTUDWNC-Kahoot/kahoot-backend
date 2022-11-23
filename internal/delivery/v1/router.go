package v1

import (
	service "examples/identity/internal/service/jwthelper"
	"examples/identity/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(g *gin.Engine)
}

type router struct {
	jwtHelper service.JWTHelper
	u         usecase.KahootUsecase
	g         usecase.GroupUsecase
}

func NewRouter(s service.JWTHelper, u usecase.KahootUsecase, g usecase.GroupUsecase) Router {
	return &router{
		jwtHelper: s,
		u:         u,
		g:         g,
	}
}

func (r *router) Register(g *gin.Engine) {

	kahoot := g.Group("/kahoots")
	kahoot.Use(r.verifyToken())
	{
		kahoot.GET("", getKahoots)
	}
	group := g.Group("/groups")
	group.Use(r.verifyToken())
	{
		group.POST("", getKahoots)
	}
}

func getKahoots(c *gin.Context) {
	c.JSON(http.StatusOK, "get kahoots")
}

func (r *router) verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request IdentityRequest
		c.ShouldBindJSON(&request)

		claims, err := r.jwtHelper.ValidateJWT(request.TokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error_message": err.Error(),
			})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, claims)
	}
}
