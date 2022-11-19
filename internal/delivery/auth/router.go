package auth

import (
	"net/http"

	"examples/identity/internal/entity"
	"examples/identity/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthRouter interface {
	Register(gr *gin.Engine)
}
type router struct {
	u usecase.AuthUsecase
}

func NewAuthRouter(u usecase.AuthUsecase) AuthRouter {
	return &router{
		u: u,
	}
}

func (r *router) Register(gr *gin.Engine) {
	user := gr.Group("/user")
	{
		user.POST("/login", r.login)
		user.POST("/register", r.register)
	}
}

func (r *router) login(c *gin.Context) {
	var request AuthenRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "error",
		})
		return
	}
	result, err := r.u.Login(&entity.User{Email: request.Email, Password: request.Password})
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, &AuthenResponse{Token: result})
}

func (r *router) register(c *gin.Context) {
	var request AuthenRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "error",
		})
		return
	}
	result, err := r.u.Register(&entity.User{Email: request.Email, Password: request.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &AuthenResponse{Token: result})
}
