package v1

import (
	service "examples/identity/internal/service/jwthelper"
	"examples/identity/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(g *gin.Engine)
}
type router struct {
	jwtHelper service.JWTHelper
	u         usecase.KahootUsecase
}

func NewRouter(jwtHelper service.JWTHelper, u usecase.KahootUsecase) Router {
	return &router{
		jwtHelper: jwtHelper,
		u:         u,
	}
}

func (r *router) verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request IdentityRequest
		c.ShouldBindJSON(&request)

		token, _ := r.jwtHelper.ValidateJWT(request.TokenString)
		if token.Valid {
			c.JSON(http.StatusOK, IdentityResponse{
				IsValid: true,
			})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func (r *router) Register(g *gin.Engine) {

	internal := g.Group("/internal")
	{
		internal.POST("/VerifyToken", r.verifyToken())
	}
}
func NewInternalRouter(s service.JWTHelper, u *usecase.KahootUsecase) Router {
	return &router{
		jwtHelper: s,
		u:         u,
	}
}
