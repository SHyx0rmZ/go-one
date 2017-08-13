package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

type arg struct {
	Kind  reflect.Kind
	Value interface{}
}

func client(t *testing.T, method string, args []arg, vals []xmlrpc.Value) *mock.Client {
	c := mock.NewClient(t)
	c.ExpectMethodName(method)
	c.ExpectArgumentCount(len(args))
	for i, a := range args {
		c.ExpectArgument(i, a.Kind, a.Value)
	}
	c.WithValue(mock.NewValue().WithValues(vals...))
	return c
}
