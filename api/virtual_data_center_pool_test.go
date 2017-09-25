package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualDataCenterPool_Info(t *testing.T) {
	c := client(t, "one.vdcpool.info", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualDataCenterPool{Client: c}
	v, err := i.Info("test-session")
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}
