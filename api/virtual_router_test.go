package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualRouter_Allocate(t *testing.T) {
	c := client(t, "one.vrouter.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Allocate("test-session", "test-template")
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_AttachNIC(t *testing.T) {
	c := client(t, "one.vrouter.attachnic", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-attribute"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.AttachNIC("test-session", 42, "test-attribute")
	if err != nil || v != 42 {
		t.Errorf("AttachNIC() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Chmod(t *testing.T) {
	c := client(t, "one.vrouter.chmod", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 1},
		{reflect.Int, 1},
		{reflect.Int, 1},
		{reflect.Int, 0},
		{reflect.Int, 0},
		{reflect.Int, 0},
		{reflect.Int, -1},
		{reflect.Int, -1},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Chown(t *testing.T) {
	c := client(t, "one.vrouter.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Delete(t *testing.T) {
	c := client(t, "one.vrouter.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_DetachNIC(t *testing.T) {
	c := client(t, "one.vrouter.detachnic", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.DetachNIC("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DetachNIC() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Info(t *testing.T) {
	c := client(t, "one.vrouter.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestVirtualRouter_Instantiate(t *testing.T) {
	c := client(t, "one.vrouter.instantiate", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 3},
		{reflect.String, "test-name-%i"},
		{reflect.Int, 21},
		{reflect.Bool, true},
		{reflect.String, "test-extra-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Instantiate("test-session", 42, 3, "test-name-%i", 21, true, "test-extra-template")
	if err != nil || v != 42 {
		t.Errorf("Instantiate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Rename(t *testing.T) {
	c := client(t, "one.vrouter.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualRouter_Update(t *testing.T) {
	c := client(t, "one.vrouter.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualRouter{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
