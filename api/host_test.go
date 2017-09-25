package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestHost_Allocate(t *testing.T) {
	c := client(t, "one.host.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-hostname"},
		{reflect.String, "test-information-manager"},
		{reflect.String, "test-virtual-manager"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Host{Client: c}
	v, err := i.Allocate("test-session", "test-hostname", "test-information-manager", "test-virtual-manager", 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestHost_Delete(t *testing.T) {
	c := client(t, "one.host.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Host{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestHost_Info(t *testing.T) {
	c := client(t, "one.host.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.Host{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestHost_Monitoring(t *testing.T) {
	c := client(t, "one.host.monitoring", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-monitor"),
	})

	i := api.Host{Client: c}
	v, err := i.Monitoring("test-session", 42)
	if err != nil || v != "test-monitor" {
		t.Errorf("Monitoring() == (%q, %q), want (%q, %v)", v, err, "test-monitor", nil)
	}
}

func TestHost_Rename(t *testing.T) {
	c := client(t, "one.host.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Host{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestHost_Status(t *testing.T) {
	c := client(t, "one.host.status", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, api.HostStatusOffline},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Host{Client: c}
	v, err := i.Status("test-session", 42, api.HostStatusOffline)
	if err != nil || v != 42 {
		t.Errorf("Status() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestHost_Update(t *testing.T) {
	c := client(t, "one.host.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Host{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
