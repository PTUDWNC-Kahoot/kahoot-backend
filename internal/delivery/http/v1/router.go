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
	p         usecase.Presentation
	g         usecase.Group
	u         usecase.User
}

const (
	BEARER_SCHEMA            = "Bearer"
	DefaultPresentationCover = "https://i.pinimg.com/564x/7e/ff/2d/7eff2dbf4765d9ce581181f5c7002a72.jpg"
)

func NewRouter(handler *gin.RouterGroup, s service.JWTHelper, u usecase.Presentation, g usecase.Group, p usecase.User) {
	newRouter(handler, s, u, g, p)
}

func newRouter(handler *gin.RouterGroup, s service.JWTHelper, u usecase.Presentation, g usecase.Group, p usecase.User) {
	r := &router{
		jwtHelper: s,
		p:         u,
		g:         g,
		u:         p,
	}
	user := handler.Group("/user")
	user.Use(r.verifyToken())
	{
		user.GET("/me", response.GinWrap(r.getProfile))
		user.POST("/update", response.GinWrap(r.updateProfile))
		user.DELETE("/delete", response.GinWrap(r.deleteProfile))
	}

	group := handler.Group("/groups")
	group.Use(r.verifyToken())
	{
		group.GET("", r.getGroups)
		group.GET("/:id", r.getByID)
		group.POST("", r.createGroup)
		group.POST("/join-group/:group-code", r.joinGroupByLink)
		group.PUT("/:id", r.groupOwnerMiddleWare, r.editGroup)
		group.DELETE("/:id", r.groupOwnerMiddleWare, r.deleteGroup)
		group.POST("/:id/invite", r.groupOwnerMiddleWare, r.invite)
		group.PUT("/:id/assign-role", r.groupOwnerMiddleWare, r.assignRole)
		ps := group.Group("/:id/presentations")
		{
			ps.POST("", response.GinWrap(r.createGroupPresentation))
			ps.GET("", response.GinWrap(r.getGroupPresentations))
		}
	}

	ps := handler.Group("/presentations")
	{
		ps.POST("", response.GinWrap(r.createMyPresentation))
		ps.GET("", response.GinWrap(r.getMyPresentations))
		ps.GET("/:id", response.GinWrap(r.getPresentation))
		ps.POST("/:id/collaborators", response.GinWrap(r.addCollaborator)) //by email
		ps.GET("/:id/collaborators", response.GinWrap(r.getCollaborators))
		ps.DELETE("/:id/collaborators/:user-id", response.GinWrap(r.removeCollaborator)) //by user_ID
		ps.PUT("/:id", r.collabMiddleWare, response.GinWrap(r.editPresentation))
		ps.DELETE("/:id", r.presentationMiddleWare, response.GinWrap(r.deletePresentation))
		ps.POST("/:id/slides", r.presentationMiddleWare, response.GinWrap(r.createSlide))
		ps.PUT("/:id/slides/:slide-id", r.presentationMiddleWare, response.GinWrap(r.editSlide))
		ps.DELETE("/:id/slides/:slide-id", r.presentationMiddleWare, response.GinWrap(r.deleteSlide))
	}
}

func (r *router) verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error_message": "Unauthorized",
			})
			return
		}
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
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error_message": "Unauthorized",
		})
		return nil
	}
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	claims, err := r.jwtHelper.ValidateJWT(tokenString)
	if err != nil || claims == nil || claims.Email == "" {
		return nil
	}
	user, err := r.u.GetSite(claims.Email)
	if err != nil {
		return nil
	}
	return user
}

func (r *router) getGroups(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
		return
	}
	groups, err := r.g.GetGroups(user.ID)
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
	user := r.getRequestingUser(c)
	fmt.Println(user)
	if user == nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
		return
	}

	if err := c.ShouldBindJSON(&group); err != nil || group.Name == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "request is invalid",
		})
		return
	}

	id, err := r.g.Create(group, user)
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

func (r *router) editGroup(c *gin.Context) {
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
	if user == nil {
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
	if user == nil {
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
	if requestingUser == nil {
		return response.Unauthorized()
	}
	fmt.Println("requestingUser:", requestingUser)

	user, err := r.u.GetProfile(requestingUser.ID)
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

	if err := r.u.UpdateProfile(&user); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) deleteProfile(c *gin.Context) *response.Response {
	user := r.getRequestingUser(c)
	if user == nil {
		return response.Unauthorized()
	}

	if err := r.u.DeleteProfile(user.ID); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) createGroupPresentation(c *gin.Context) *response.Response {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.StatusBadRequest()
	}

	presentation := &entity.Presentation{}

	if err := c.ShouldBindJSON(&presentation); err != nil {
		return response.Failure(err)
	}

	presentation.GroupID = uint32(groupId)
	if presentation.CoverImageURL == "" {
		presentation.CoverImageURL = DefaultPresentationCover
	}

	err = r.p.CreatePresentation(presentation)
	if err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(presentation.ID)
}

func (r *router) getGroupPresentations(c *gin.Context) *response.Response {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.StatusBadRequest()
	}

	presentationList, err := r.p.GroupCollection(uint32(groupId))
	if err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(presentationList)
}

func (r *router) getPresentation(c *gin.Context) *response.Response {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.StatusBadRequest()
	}

	presentation, err := r.p.GetPresentation(uint32(id))
	if err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(presentation)
}

func (r *router) editPresentation(c *gin.Context) *response.Response {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.StatusBadRequest()
	}

	presentation := entity.Presentation{}
	presentation.ID = uint32(id)
	if err := c.ShouldBindJSON(&presentation); err != nil {
		return response.StatusBadRequest()
	}

	if err := r.p.UpdatePresentation(&presentation); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) deletePresentation(c *gin.Context) *response.Response {
	user := r.getRequestingUser(c)
	if user == nil {
		return response.Unauthorized()
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.StatusBadRequest()
	}

	if err := r.p.DeletePresentation(uint32(id), user.ID); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) createMyPresentation(c *gin.Context) *response.Response {
	user := r.getRequestingUser(c)
	if user == nil {
		return response.Unauthorized()
	}

	presentation := &entity.Presentation{}

	if err := c.ShouldBindJSON(&presentation); err != nil {
		return response.Failure(err)
	}

	presentation.Owner = user.ID
	if presentation.CoverImageURL == "" {
		presentation.CoverImageURL = DefaultPresentationCover
	}

	if err := r.p.CreatePresentation(presentation); err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(presentation.ID)
}

func (r *router) getMyPresentations(c *gin.Context) *response.Response {
	user := r.getRequestingUser(c)
	if user == nil {
		return response.Unauthorized()
	}

	presentationList, err := r.p.MyCollection(uint32(user.ID))
	if err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(presentationList)
}

func (r *router) createSlide(c *gin.Context) *response.Response {
	presentationId, err := strconv.Atoi(c.Param("id"))
	if err != nil || presentationId == 0 {
		return response.StatusBadRequest()
	}

	slide := &entity.Slide{}
	if err := c.ShouldBindJSON(slide); err != nil {
		return response.StatusBadRequest()
	}

	slide.PresentationID = uint32(presentationId)

	if err := r.p.CreateSlide(slide); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) editSlide(c *gin.Context) *response.Response {
	id, err := strconv.Atoi(c.Param("slide-id"))
	if err != nil || id == 0 {
		return response.StatusBadRequest()
	}

	slide := &entity.Slide{}
	if err := c.ShouldBindJSON(slide); err != nil {
		return response.StatusBadRequest()
	}

	slide.ID = uint32(id)

	if err := r.p.UpdateSlide(slide); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}
func (r *router) deleteSlide(c *gin.Context) *response.Response {
	id, err := strconv.Atoi(c.Param("slide-id"))
	if err != nil || id == 0 {
		return response.StatusBadRequest()
	}

	if err := r.p.DeleteSlide(uint32(id)); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) addCollaborator(c *gin.Context) *response.Response {
	presentationId, err := strconv.Atoi(c.Param("id"))
	if err != nil || presentationId == 0 {
		return response.StatusBadRequest()
	}
	e := EmailList{}
	if err := c.ShouldBindJSON(&e); err != nil {
		return response.StatusBadRequest()
	}

	if err := r.p.AddCollaborator(e.Emails, uint32(presentationId)); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}

func (r *router) getCollaborators(c *gin.Context) *response.Response {
	presentationId, err := strconv.Atoi(c.Param("id"))
	if err != nil || presentationId == 0 {
		return response.StatusBadRequest()
	}

	collaborators, err := r.p.GetCollaborators(uint32(presentationId))
	if err != nil {
		return response.Failure(err)
	}

	return response.SuccessWithData(collaborators)
}

func (r *router) removeCollaborator(c *gin.Context) *response.Response {
	presentationId, err := strconv.Atoi(c.Param("id"))
	if err != nil || presentationId == 0 {
		return response.StatusBadRequest()
	}
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil || userId == 0 {
		return response.StatusBadRequest()
	}

	if err := r.p.RemoveCollaborator(uint32(userId), uint32(presentationId)); err != nil {
		return response.Failure(err)
	}

	return response.Success()
}
