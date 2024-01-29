package todo

import "github.com/gin-gonic/gin"

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

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewMyContext(ctx))
	}
}
