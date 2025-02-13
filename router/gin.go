package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pallat/todoapi/todo"
)

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *MyContext) JSON(statusCode int, v interface{}) {
	c.Context.JSON(statusCode, v)
}

func (c *MyContext) TransactionID() string {
	return c.Context.Request.Header.Get("TransactionID")
}

func (c *MyContext) Audience() string {
	if aud, ok := c.Context.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func NewGinHandler(handler func(todo.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewMyContext(ctx))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(todo.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}
