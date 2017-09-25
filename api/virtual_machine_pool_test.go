package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualMachinePool_Accounting(t *testing.T) {
	c := client(t, "one.vmpool.accounting", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, api.FilterAll},
		{reflect.Int, -1},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-accounting"),
	})

	i := api.VirtualMachinePool{Client: c}
	v, err := i.Accounting("test-session", api.FilterAll, -1, -1)
	if err != nil || v != "test-accounting" {
		t.Errorf("Accounting() == (%q, %q), want (%q, %v)", v, err, "test-accounting", nil)
	}
}

func TestVirtualMachinePool_CalculateShowback(t *testing.T) {
	c := client(t, "one.vmpool.calculateshowback", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
	})

	i := api.VirtualMachinePool{Client: c}
	err := i.CalculateShowback("test-session", -1, -1, -1, -1)
	if err != nil {
		t.Errorf("CalculateShowback() == %q, want %v", err, nil)
	}
}

func TestVirtualMachinePool_Info(t *testing.T) {
	c := client(t, "one.vmpool.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, api.FilterAll},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, api.StateAny},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualMachinePool{Client: c}
	v, err := i.Info("test-session", api.FilterAll, -1, -1, api.StateAny)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestVirtualMachinePool_Monitoring(t *testing.T) {
	c := client(t, "one.vmpool.monitoring", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, api.FilterAll},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-monitor"),
	})

	i := api.VirtualMachinePool{Client: c}
	v, err := i.Monitoring("test-session", api.FilterAll)
	if err != nil || v != "test-monitor" {
		t.Errorf("Monitoring() == (%q, %q), want (%q, %v)", v, err, "test-monitor", nil)
	}
}

func TestVirtualMachinePool_Showback(t *testing.T) {
	c := client(t, "one.vmpool.showback", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, api.FilterAll},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-showback"),
	})

	i := api.VirtualMachinePool{Client: c}
	v, err := i.Showback("test-session", api.FilterAll, -1, -1, -1, -1)
	if err != nil || v != "test-showback" {
		t.Errorf("Showback() == (%q, %q), want (%q, %v)", v, err, "test-showback", nil)
	}
}
