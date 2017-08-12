package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestAccessControlList_AddRule(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.acl.addrule")
	c.ExpectArgumentCount(4)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.String, "test-user")
	c.ExpectArgument(2, reflect.String, "test-resource")
	c.ExpectArgument(3, reflect.String, "test-rights")

	i := api.AccessControlList{
		Client: c,
	}

	v, err := i.AddRule("test-session", "test-user", "test-resource", "test-rights")
	if err != nil || v != 42 {
		t.Errorf("AddRule() == (%v, %q), want (%v, %q)", v, err, 42, nil)
	}
}

func TestAccessControlList_DelRule(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.acl.delrule")
	c.ExpectArgumentCount(2)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)

	i := api.AccessControlList{
		Client: c,
	}

	v, err := i.DelRule("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("DelRule() == (%v, %q), want (%v, %q)", v, err, 42, nil)
	}
}

func TestAccessControlList_Info(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithString("test-info"),
		),
	)

	c.ExpectMethodName("one.acl.info")
	c.ExpectArgumentCount(2)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)

	i := api.AccessControlList{
		Client: c,
	}

	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%v, %q), want (%v, %q)", v, err, "test-info", nil)
	}
}
