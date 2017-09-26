package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestSecurityGroup_Allocate(t *testing.T) {
	c := client(t, "one.secgroup.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Allocate("test-session", "test-template")
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Chmod(t *testing.T) {
	c := client(t, "one.secgroup.chmod", []arg{
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

	i := api.SecurityGroup{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Chown(t *testing.T) {
	c := client(t, "one.secgroup.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Clone(t *testing.T) {
	c := client(t, "one.secgroup.clone", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Clone("test-session", 42, "test-name", 21)
	if err != nil || v != 42 {
		t.Errorf("Clone() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Commit(t *testing.T) {
	c := client(t, "one.secgroup.commit", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Commit("test-session", 42, true)
	if err != nil || v != 42 {
		t.Errorf("Commit() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Delete(t *testing.T) {
	c := client(t, "one.secgroup.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Info(t *testing.T) {
	c := client(t, "one.secgroup.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%v, %q), want (%v, %v)", v, err, "test-info", nil)
	}
}

func TestSecurityGroup_Rename(t *testing.T) {
	c := client(t, "one.secgroup.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestSecurityGroup_Update(t *testing.T) {
	c := client(t, "one.secgroup.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.SecurityGroup{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
