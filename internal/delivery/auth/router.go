package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"examples/identity/config"
	"examples/identity/internal/entity"
	mailService "examples/identity/internal/service/mail"

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
		auth.POST("/emailVerification", r.emailVerification)
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
	body, err := ioutil.ReadAll(userInfo.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "cannot read user info",
		})
		return
	}
	data := GoogleResponse{}
	json.Unmarshal([]byte(string(body)), &data)
	fmt.Println(data.Email)

	isEmailExisted := r.u.CheckEmailExisted(data.Email)
	if !isEmailExisted {
		r.u.Register(&entity.User{Email: data.Email, Password: "google"})
	}
	user, groups, _, token1, err := r.u.Login(&entity.User{Email: data.Email, Password: "google"})
	user.Groups = groups
	c.JSON(http.StatusOK, &AuthenResponse{
		Token:         token1,
		ID:            user.ID,
		Name:          user.Name,
		Workplace:     user.Workplace,
		Organization:  user.Organization,
		CoverImageURL: user.CoverImageURL,
		Groups:        groups,
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
	verifyCode := mailService.GenerateVerifyCode()
	orderId, err := r.u.CreateRegisterOrder(&entity.RegisterOrder{Email: request.Email, VerifyCode: verifyCode})
	if orderId == 0 || err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email is already used",
		})
		return
	}
	err = mailService.SendEmail(verifyCode, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Cannot send email",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "mail sent, waiting for email verification",
		"status":  "pending",
	})
}

func (r *router) emailVerification(c *gin.Context) {
	var request RegisterWithVerification

	err := c.ShouldBindJSON(&request)
	if err != nil || !request.Validate() {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "Request is invalid",
		})
		return
	}
	isVerified := r.u.VerifyEmail(request.Email, request.VerifyCode)
	if !isVerified {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Verify code is invalid",
		})
		return
	}
	err = r.u.Register(&entity.User{Email: request.Email, Password: request.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email is already used",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "register successfully",
		"status":  "ok",
	})
}
