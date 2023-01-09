package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *router) groupMiddleWare(c *gin.Context) {
	user := r.getRequestingUser(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(401, map[string]string{
			"message": "request is invalid",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if id == 0 || err != nil {
		c.AbortWithStatusJSON(400, map[string]string{
			"message": "request is invalid",
			"error":   err.Error(),
		})
		return
	}
	group, err := r.g.Get(uint32(id))
	if group.Owner != user.ID {
		c.AbortWithStatusJSON(403, map[string]string{
			"message": "request is invalid",
		})
		return
	}

	c.Next()
}
