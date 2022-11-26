package v1

import (
	"examples/identity/internal/entity"
	service "examples/identity/internal/service/jwthelper"
	"examples/identity/internal/usecase"
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
}

const BEARER_SCHEMA = "Bearer"

func NewRouter(s service.JWTHelper, u usecase.KahootUsecase, g usecase.GroupUsecase) Router {
	return &router{
		jwtHelper: s,
		u:         u,
		g:         g,
	}
}

func (r *router) Register(g *gin.Engine) {
	g.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	g.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "welcome to kahoot",
		})
	})
	user := g.Group("/user")
	user.Use(r.verifyToken())
	{
		user.GET("/profile")
		user.GET("/update")
		user.GET("/delete")
	}

	kahoot := g.Group("/kahoots")
	kahoot.Use(r.verifyToken())
	{
		// kahoot.GET("", getKahoots)
	}

	group := g.Group("/groups")
	group.Use(r.verifyToken())
	{
		group.GET("", r.getGroups)
		group.GET("/:id", r.getByID)
		group.POST("", r.createGroup)
		group.PUT("/:id", r.updateGroup)
		group.DELETE("/:id", r.deleteGroup)

	}
	g.GET("/join-group/:group-code", r.joinGroupByLink)
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
			fmt.Println(err)
			return
		}
		return
	}
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
	c.ShouldBindJSON(&group)
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
	//check token and get user email to join group
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	claims, err := r.jwtHelper.ValidateJWT(tokenString)
	if err != nil || claims.Email == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error_message": err.Error(),
		})
		return
	}
	//group jobs
	groupCode := c.Param("group-code")
	if groupCode == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "invalid group",
		})
		return
	}
	group, err := r.g.JoinGroupByLink(claims.Email, groupCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "unable to join group",
		})
		return
	}
	c.JSON(http.StatusOK, group)
}
