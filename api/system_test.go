package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestSystem_Config(t *testing.T) {
	c := client(t, "one.system.config", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-config"),
	})

	i := api.System{Client: c}
	v, err := i.Config("test-session")
	if err != nil || v != "test-config" {
		t.Errorf("Config() == (%q, %q), want (%q, %v)", v, err, "test-config", nil)
	}
}

func TestSystem_Version(t *testing.T) {
	c := client(t, "one.system.version", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-version"),
	})

	i := api.System{Client: c}
	v, err := i.Version("test-session")
	if err != nil || v != "test-version" {
		t.Errorf("Version() == (%q, %q), want (%q, %v)", v, err, "test-version", nil)
	}
}
