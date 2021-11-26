package testutil

import (
	"testing"

	controllers "github.com/hamg26/academy-go-q42021/interface/controllers"
)

type Context struct {
	FakeError  error
	Params     map[string]string
	StatusCode int
	store      map[string]interface{}
}

func (c *Context) JSON(code int, i interface{}) error {
	c.store["Response"] = i
	c.store["StatusCode"] = code
	return c.FakeError
}

func (c *Context) Bind(i interface{}) error {
	return c.FakeError
}

func (c *Context) Param(p string) string {
	return c.Params[p]
}

func (c *Context) Get(key string) interface{} {
	return c.store[key]
}

func (c *Context) Set(key string, val interface{}) {
	if c.store == nil {
		c.store = make(map[string]interface{})
	}
	c.store[key] = val
}

func NewContextMock(t *testing.T, fakeError error, paramsValues map[string]string) controllers.Context {
	t.Helper()
	store := map[string]interface{}{
		"Response":   nil,
		"StatusCode": nil,
	}
	return &Context{Params: paramsValues, store: store, FakeError: fakeError}
}
