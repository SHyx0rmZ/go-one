package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestHostPool_Info(t *testing.T) {
	c := client(t, "one.hostpool.info", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.HostPool{Client: c}
	v, err := i.Info("test-session")
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestHostPool_Monitoring(t *testing.T) {
	c := client(t, "one.hostpool.monitoring", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-monitor"),
	})

	i := api.HostPool{Client: c}
	v, err := i.Monitoring("test-session")
	if err != nil || v != "test-monitor" {
		t.Errorf("Monitoring() == (%q, %q), want (%q, %v)", v, err, "test-monitor", nil)
	}
}
