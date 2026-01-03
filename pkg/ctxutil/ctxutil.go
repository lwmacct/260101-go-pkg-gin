package ctxutil

import (
	"github.com/gin-gonic/gin"
)

const (
	// Permissions 表示上下文中存储权限列表的键
	Permissions = "permissions"

	// Roles 表示上下文中存储角色列表的键
	Roles = "roles"

	// Username 表示上下文中存储用户名的键
	Username = "username"

	// UserID 表示上下文中存储用户 ID 的键
	UserID = "user_id"

	// UserRole 表示上下文中存储用户角色的键
	UserRole = "user_role"

	// OrgID 表示上下文中存储组织 ID 的键
	OrgID = "org_id"

	// TeamID 表示上下文中存储团队 ID 的键
	TeamID = "team_id"

	// RequestID 表示上下文中存储请求 ID 的键
	RequestID = "request_id"
)

// Get 从 Context 安全获取指定 key 的值并断言为类型 T。
// 返回值和成功标志，类似 map 查找的 comma-ok 模式。
func Get[T any](c *gin.Context, key string) (T, bool) {
	var zero T
	value, exists := c.Get(key)
	if !exists {
		return zero, false
	}
	result, ok := value.(T)
	return result, ok
}
