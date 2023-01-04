package adapter

import (
	"github.com/gin-gonic/gin"
	"io"
	service "project_api/services"
)

type GinAdapter interface {
	Stream(c *gin.Context)
	Post(msg string)
}

type ginAdapter struct {
	rm service.Manager
}

func NewGinAdapter(rm service.Manager) GinAdapter {
	return &ginAdapter{rm}
}

func (ga *ginAdapter) Post(msg string) {
	ga.rm.Submit(msg)
}
func (ga *ginAdapter) Stream(c *gin.Context) {

	listener := ga.rm.OpenListener()
	defer ga.rm.CloseListener(listener)

	clientGone := c.Request.Context().Done()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case msg := <-listener:
			c.SSEvent("message", msg)
			return true
		}
	})
}