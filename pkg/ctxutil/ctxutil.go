package ctxutil

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/response"
)

// Context key 常量定义
const (

	// KeyUsername 表示上下文中存储用户名的键
	KeyUsername = "username"

	// KeyUserID 表示上下文中存储用户 ID 的键
	KeyUserID = "user_id"

	// KeyUserRole 表示上下文中存储用户角色的键
	KeyUserRole = "user_role"

	// KeyOrgID 表示上下文中存储组织 ID 的键
	KeyOrgID = "org_id"

	// KeyTeamID 表示上下文中存储团队 ID 的键
	KeyTeamID = "team_id"

	// KeyRequestID 表示上下文中存储请求 ID 的键
	KeyRequestID = "request_id"
)

// MustGet 从 Context 获取指定 key 的值并断言为类型 T，
// 若 key 不存在或类型不匹配则返回 401 响应并中止请求处理。
//
// 类型参数 T 可以是任意类型，例如：
//
//	userID := ctxutil.MustGet[uint](c, ctxutil.KeyUserID)
//	username := ctxutil.MustGet[string](c, ctxutil.KeyUsername)
func MustGet[T any](c *gin.Context, key string) T {
	var zero T

	value, exists := c.Get(key)
	if !exists {
		response.Unauthorized(c, fmt.Sprintf("missing key %q in context", key))
		c.Abort()
		return zero
	}

	result, ok := value.(T)
	if !ok {
		response.Unauthorized(c, fmt.Sprintf("invalid type for key %q", key))
		c.Abort()
		return zero
	}

	return result
}

// MustGetUserID 从 Context 获取当前用户 ID，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetUserID(c *gin.Context) uint {
	return MustGet[uint](c, KeyUserID)
}

// MustGetUsername 从 Context 获取当前用户名，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetUsername(c *gin.Context) string {
	return MustGet[string](c, KeyUsername)
}

// MustGetUserRole 从 Context 获取当前用户角色，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetUserRole(c *gin.Context) string {
	return MustGet[string](c, KeyUserRole)
}

// MustGetOrgID 从 Context 获取当前组织 ID，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetOrgID(c *gin.Context) uint {
	return MustGet[uint](c, KeyOrgID)
}

// MustGetTeamID 从 Context 获取当前团队 ID，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetTeamID(c *gin.Context) uint {
	return MustGet[uint](c, KeyTeamID)
}

// MustGetRequestID 从 Context 获取当前请求 ID，
// 若不存在或类型不匹配则返回 401 响应并中止请求处理。
func MustGetRequestID(c *gin.Context) string {
	return MustGet[string](c, KeyRequestID)
}
