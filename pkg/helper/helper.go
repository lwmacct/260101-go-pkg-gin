package helper

import (
	"github.com/gin-gonic/gin"

	"github.com/lwmacct/260101-go-pkg-gin/pkg/response"
)

// GetUserIDMust 从 Context 获取当前用户 ID，若不存在则直接返回 401 响应并中止请求处理。
func GetUserIDMust(c *gin.Context, err error) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c)
		c.Abort()
		return 0
	}

	id, ok := userID.(uint)
	if !ok {
		response.Unauthorized(c, err.Error())
		c.Abort()
		return 0
	}

	return id
}

// GetOrgIDMust 从 Context 获取当前组织 ID，若不存在则直接返回 401 响应并中止请求处理。
func GetOrgIDMust(c *gin.Context, err error) uint {
	orgID, exists := c.Get("org_id")
	if !exists {
		response.Unauthorized(c)
		c.Abort()
		return 0
	}

	id, ok := orgID.(uint)
	if !ok {
		response.Unauthorized(c, err.Error())
		c.Abort()
		return 0
	}

	return id
}

// GetTeamIDMust 从 Context 获取当前团队 ID，若不存在则直接返回 401 响应并中止请求处理。
func GetTeamIDMust(c *gin.Context, err error) uint {
	teamID, exists := c.Get("team_id")
	if !exists {
		response.Unauthorized(c)
		c.Abort()
		return 0
	}

	id, ok := teamID.(uint)
	if !ok {
		response.Unauthorized(c, err.Error())
		c.Abort()
		return 0
	}

	return id
}
