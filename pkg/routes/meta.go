package routes

// HTTPMethod HTTP 请求方法。
type HTTPMethod string

// HTTP 方法常量。
const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)
