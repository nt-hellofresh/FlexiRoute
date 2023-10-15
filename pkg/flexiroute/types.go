package flexiroute

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

type ApiHandler func(ctx *Context) error
type Middleware func(handler ApiHandler) ApiHandler
type HttpMethod string
