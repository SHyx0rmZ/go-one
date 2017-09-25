package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualRouterPool_Info(t *testing.T) {
	c := client(t, "one.vrouterpool.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, api.FilterAll},
		{reflect.Int, -1},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualRouterPool{Client: c}
	v, err := i.Info("test-session", api.FilterAll, -1, -1)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}
