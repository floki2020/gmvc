package gmvc

type View interface {
	Render(c *Context, name string, value interface{}) error
}
