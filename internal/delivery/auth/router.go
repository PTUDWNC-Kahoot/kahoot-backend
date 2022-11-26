package auth

import (
	"fmt"
	"net/http"

	"examples/identity/config"
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

func (r *router) Register(g *gin.Engine) {
	auth := g.Group("/auth")
	{
		auth.POST("/login", r.login)
		auth.POST("/register", r.register)
	}
	googleAuth := g.Group("/google")
	{
		googleAuth.GET("/login", r.googleLogin)
		googleAuth.GET("/callback", r.googleCallback)
	}
}

func (r *router) googleLogin(c *gin.Context) {
	googleConfig := config.SetUpConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusSeeOther, url)
}

func (r *router) googleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "state is invalid",
		})
		return
	}
	code := c.Query("code")
	googleConfig := config.SetUpConfig()
	token, err := googleConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "cannot get token",
		})
		return
	}
	userInfo, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "cannot get user info",
		})
		return
	}
	mail := userInfo.Header.Get("email")
	fmt.Println("mail:::::", mail)
	c.JSON(http.StatusOK, &AuthenResponse{
		Token: token.AccessToken,
	})
}

func (r *router) login(c *gin.Context) {
	var request AuthenRequest
	err := c.ShouldBindJSON(&request)
	if err != nil || !request.Validate() {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "email or password is invalid",
		})
		return
	}

	user, groups, kahoots, token, err := r.u.Login(&entity.User{Email: request.Email, Password: request.Password})
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, &AuthenResponse{
		Token:         token,
		ID:            user.ID,
		Name:          user.Name,
		Workplace:     user.Workplace,
		Organization:  user.Organization,
		CoverImageURL: user.CoverImageURL,
		Groups:        groups,
		Kahoots:       kahoots,
	})
}

func (r *router) register(c *gin.Context) {
	var request AuthenRequest

	err := c.ShouldBindJSON(&request)
	if err != nil || !request.Validate() {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "email or password is invalid",
		})
		return
	}
	id, token, err := r.u.Register(&entity.User{Email: request.Email, Password: request.Password})
	if err != nil || token == "" {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "email is already used",
		})
		return
	}
	c.JSON(http.StatusOK, &AuthenResponse{
		Token: token,
		ID:    id,
	})
}
