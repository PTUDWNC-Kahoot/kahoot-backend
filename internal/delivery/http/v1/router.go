package v1

import (
	"examples/kahootee/internal/entity"
	service "examples/kahootee/internal/service/jwthelper"
	"examples/kahootee/internal/usecase"
	"examples/kahootee/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(g *gin.Engine)
}

type router struct {
	jwtHelper service.JWTHelper
	u         usecase.KahootUsecase
	g         usecase.GroupUsecase
	p         usecase.User
}

const (
	BEARER_SCHEMA = "Bearer"
)

func NewRouter(handler *gin.RouterGroup, s service.JWTHelper, u usecase.KahootUsecase, g usecase.GroupUsecase) {
	newRouter(handler, s, u, g)
}

func newRouter(handler *gin.RouterGroup, s service.JWTHelper, u usecase.KahootUsecase, g usecase.GroupUsecase) {
	r := &router{
		jwtHelper: s,
		u:         u,
		g:         g,
	}
	user := handler.Group("/user")
	user.Use(r.verifyToken())
	{
		user.GET("/me", response.GinWrap(r.getProfile))
		user.POST("/update", response.GinWrap(r.updateProfile))
		user.DELETE("/delete", response.GinWrap(r.deleteProfile))
	}

	kahoot := handler.Group("/kahoots")
	kahoot.Use(r.verifyToken())
	{
		// kahoot.GET("", getKahoots)
	}

	group := handler.Group("/groups")
	group.Use(r.verifyToken())
	{
		group.GET("", r.getGroups)
		group.GET("/:id", r.getByID)
		group.POST("", r.createGroup)
		group.PUT("/:id", r.updateGroup)
		group.DELETE("/:id", r.deleteGroup)
		group.POST("/join-group/:group-code", r.joinGroupByLink)
		group.POST("/:id/invite", r.invite)
		group.PUT("/:id/assign-role", r.assignRole)
	}
}

func (r *router) verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		_, err := r.jwtHelper.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error_message": err.Error(),
			})
			return
		}
		return
	}
}

func (r *router) getRequestingUser(c *gin.Context) *entity.User {
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	claims, err := r.jwtHelper.ValidateJWT(tokenString)
	if err != nil || claims.Email == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error_message": "Do not have permission",
		})
		return nil
	}
	return &entity.User{ID: claims.ID, Email: claims.Email}
}

func (r *router) getGroups(c *gin.Context) {
	groups, err := r.g.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (r *router) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": err.Error(),
		})
	}

	group, err := r.g.Get(uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (r *router) createGroup(c *gin.Context) {
	group := &entity.Group{}

	if err := c.ShouldBindJSON(&group); err != nil || group.Name == "" || group.AdminID == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "request is invalid",
		})
		return
	}

	id, err := r.g.Create(group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "group created successfully",
		"id":      strconv.Itoa(int(id)),
	})
}

func (r *router) updateGroup(c *gin.Context) {
	group := entity.Group{}

	id, err := strconv.Atoi(c.Param("id"))
	if id == 0 || err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "request is invalid",
			"error":   err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "request is invalid",
			"error":   err.Error(),
		})
		return
	}

	group.ID = uint32(id)

	err = r.g.Update(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": " group updated successfully!",
	})
}

func (r *router) deleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if id == 0 || err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "request is invalid",
			"error":   err.Error(),
		})
		return
	}
	err = r.g.Delete(uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": " group deleted successfully!",
		"id":      strconv.Itoa(int(id)),
	})
}

func (r *router) joinGroupByLink(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error_message": "unauthorized",
		})
		return
	}

	groupCode := c.Param("group-code")
	if groupCode == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid group",
		})
		return
	}
	group, err := r.g.JoinGroupByLink(user.Email, groupCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "unable to join group",
		})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (r *router) invite(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid request",
		})
		return
	}

	e := EmailList{}
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid request email_list",
		})
		return
	}

	if err := r.g.Invite(e.Emails, uint32(groupID)); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": "unable to invite",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "invited successfully",
	})
}

func (r *router) assignRole(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid request",
		})
		return
	}

	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error_message": "unauthorized",
		})
		return
	}

	groupUser := entity.GroupUser{}
	if err := c.ShouldBindJSON(&groupUser); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid request",
		})
		return
	}

	groupUser.GroupID = uint32(groupID)

	if err := r.g.AssignRole(&groupUser, user.Email); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "role assigned successfully",
	})
}

func (r *router) getProfile(c *gin.Context) *response.Response {
	requestingUser := r.getRequestingUser(c)
	if requestingUser == nil || requestingUser.ID == 0 {
		return response.Unauthorized()
	}
	fmt.Println("requestingUser:", requestingUser)

	user, err := r.p.GetProfile(requestingUser.ID)
	fmt.Println("err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": err.Error(),
		})
		return response.Failure(err)
	}

	return response.SuccessWithData(user)
}

func (r *router) updateProfile(c *gin.Context) *response.Response {
	user := entity.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		return response.Failure(err)
	}

	if err := r.p.UpdateProfile(&user); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) deleteProfile(c *gin.Context) *response.Response {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		return response.Unauthorized()
	}

	if err := r.p.DeleteProfile(user.ID); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}
