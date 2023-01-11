package v1

import (
	"examples/kahootee/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *router) groupOwnerMiddleWare(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	groupId, err := strconv.Atoi(c.Param("id"))
	if groupId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}
	group, err := r.g.Get(uint32(groupId))
	for _, groupUser := range group.Users {
		if groupUser.UserID == user.ID && groupUser.Role == entity.Owner {
			c.Next()
			return
		}
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) groupMiddleWare(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	groupId, err := strconv.Atoi(c.Param("id"))
	if groupId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}
	group, err := r.g.Get(uint32(groupId))
	for _, groupUser := range group.Users {
		if groupUser.UserID == user.ID && (groupUser.Role == entity.Owner || groupUser.Role == entity.CoOwner) {
			c.Next()
			return
		}
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) presentationMiddleWare(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	presentationId, err := strconv.Atoi(c.Param("id"))
	if presentationId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}

	presentation, err := r.p.GetPresentation(uint32(presentationId))
	if presentation.Owner == user.ID {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) collabMiddleWare(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	presentationId, err := strconv.Atoi(c.Param("id"))
	if presentationId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}

	presentation, err := r.p.GetPresentation(uint32(presentationId))
	for _, collab := range presentation.Collaborators {
		if collab.UserID == user.ID {
			c.Next()
			return
		}
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) presentMiddleWare(c *gin.Context) {
	token := c.Query("token")
	user := r.getUserByToken(token)
	if user == nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	presentationId, err := strconv.Atoi(c.Param("id"))
	if presentationId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}

	presentation, err := r.p.GetPresentation(uint32(presentationId))
	if presentation.Owner == user.ID {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) groupPresentMiddleWare(c *gin.Context) {
	token := c.Param("token")
	user := r.getUserByToken(token)
	if user == nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	presentationId, err := strconv.Atoi(c.Param("id"))
	if presentationId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}

	presentation, err := r.p.GetPresentation(uint32(presentationId))
	if presentation.Owner == user.ID {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}

func (r *router) groupPresentJoinMiddleWare(c *gin.Context) {
	token := c.Param("token")
	user := r.getUserByToken(token)
	if user == nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
		})
		return
	}
	presentationId, err := strconv.Atoi(c.Param("id"))
	if presentationId == 0 || err != nil {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "Do not have permission",
			"error":   err.Error(),
		})
		return
	}

	presentation, err := r.p.GetPresentation(uint32(presentationId))
	if presentation.Owner == user.ID {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(403, map[string]string{
		"message": "Do not have permission",
	})
	return
}
