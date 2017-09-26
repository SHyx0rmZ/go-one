package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestUser_AddGroup(t *testing.T) {
	c := client(t, "one.user.addgroup", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.AddGroup("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddGroup() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Allocate(t *testing.T) {
	c := client(t, "one.user.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-username"},
		{reflect.String, "test-password"},
		{reflect.String, "test-authentication-driver"},
		{reflect.Slice, []int{21, 21}},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Allocate("test-session", "test-username", "test-password", "test-authentication-driver", 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Chauth(t *testing.T) {
	c := client(t, "one.user.chauth", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-authentication-driver"},
		{reflect.String, ""},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Chauth("test-session", 42, "test-authentication-driver", "")
	if err != nil || v != 42 {
		t.Errorf("Chauth() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Chgrp(t *testing.T) {
	c := client(t, "one.user.chgrp", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Chgrp("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("Chgrp() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Delete(t *testing.T) {
	c := client(t, "one.user.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_DelGroup(t *testing.T) {
	c := client(t, "one.user.delgroup", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.DelGroup("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelGroup() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Info(t *testing.T) {
	c := client(t, "one.user.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.User{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestUser_Login(t *testing.T) {
	c := client(t, "one.user.login", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-username"},
		{reflect.String, ""},
		{reflect.Int, 300},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-token"),
	})

	i := api.User{Client: c}
	v, err := i.Login("test-session", "test-username", "", 300, 21)
	if err != nil || v != "test-token" {
		t.Errorf("Login() == (%q, %q), want (%q, %v)", v, err, "test-token", nil)
	}
}

func TestUser_Password(t *testing.T) {
	c := client(t, "one.user.password", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-password"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Password("test-session", 42, "test-password")
	if err != nil || v != 42 {
		t.Errorf("Password() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Quota(t *testing.T) {
	c := client(t, "one.user.quota", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Quota("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("Quota() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestUser_Update(t *testing.T) {
	c := client(t, "one.user.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.User{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
