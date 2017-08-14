package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestGroup_AddAdmin(t *testing.T) {
	c := client(t, "one.group.addadmin", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.AddAdmin("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddAdmin() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestGroup_Allocate(t *testing.T) {
	c := client(t, "one.group.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.Allocate("test-session", "test-name")
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestGroup_DelAdmin(t *testing.T) {
	c := client(t, "one.group.deladmin", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.DelAdmin("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelAdmin() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestGroup_Delete(t *testing.T) {
	c := client(t, "one.group.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestGroup_Info(t *testing.T) {
	c := client(t, "one.group.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.Group{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestGroup_Quota(t *testing.T) {
	c := client(t, "one.group.quota", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.Quota("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("Quota() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestGroup_Update(t *testing.T) {
	c := client(t, "one.group.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Group{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
