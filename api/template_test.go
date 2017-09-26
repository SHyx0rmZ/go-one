package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestTemplate_Allocate(t *testing.T) {
	c := client(t, "one.template.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Allocate("test-session", "test-template")
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Chmod(t *testing.T) {
	c := client(t, "one.template.chmod", []arg{
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
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1, true)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Chown(t *testing.T) {
	c := client(t, "one.template.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Clone(t *testing.T) {
	c := client(t, "one.template.clone", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Clone("test-session", 42, "test-name", true)
	if err != nil || v != 42 {
		t.Errorf("Clone() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Delete(t *testing.T) {
	c := client(t, "one.template.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Delete("test-session", 42, true)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Info(t *testing.T) {
	c := client(t, "one.template.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.Template{Client: c}
	v, err := i.Info("test-session", 42, true)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestTemplate_Instantiate(t *testing.T) {
	c := client(t, "one.template.instantiate", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
		{reflect.Bool, true},
		{reflect.String, "test-extra-template"},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Instantiate("test-session", 42, "test-name", true, "test-extra-template", true)
	if err != nil || v != 42 {
		t.Errorf("Instantiate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Rename(t *testing.T) {
	c := client(t, "one.template.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestTemplate_Update(t *testing.T) {
	c := client(t, "one.template.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Template{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
