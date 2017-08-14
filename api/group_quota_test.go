package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestGroupQuota_Info(t *testing.T) {
	c := client(t, "one.groupquota.info", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.GroupQuota{Client: c}
	v, err := i.Info("test-session")
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestGroupQuota_Update(t *testing.T) {
	c := client(t, "one.groupquota.update", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-template"),
	})

	i := api.GroupQuota{Client: c}
	v, err := i.Update("test-session", "test-template")
	if err != nil || v != "test-template" {
		t.Errorf("Update() == (%q, %q), want (%q, %v)", v, err, "test-template", nil)
	}
}
