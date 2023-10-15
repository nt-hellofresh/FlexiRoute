package flexiroute

type RouterFacade interface {
	Namespace(name string) RouterFacade
	Use(fn Middleware)
	Get(pattern string, handler ApiHandler)
	Put(pattern string, handler ApiHandler)
	Patch(pattern string, handler ApiHandler)
	Post(pattern string, handler ApiHandler)
	Delete(pattern string, handler ApiHandler)
	LoadTemplates(directory string)
	ServeHTTP(addr string) error
}
