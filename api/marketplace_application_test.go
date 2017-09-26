package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestMarketplaceApplication_Allocate(t *testing.T) {
	c := client(t, "one.marketapp.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Allocate("test-session", "test-template", 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Chmod(t *testing.T) {
	c := client(t, "one.marketapp.chmod", []arg{
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

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Chown(t *testing.T) {
	c := client(t, "one.marketapp.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Delete(t *testing.T) {
	c := client(t, "one.marketapp.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Enable(t *testing.T) {
	c := client(t, "one.marketapp.enable", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Enable("test-session", 42, true)
	if err != nil || v != 42 {
		t.Errorf("Enable() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Info(t *testing.T) {
	c := client(t, "one.marketapp.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestMarketplaceApplication_Rename(t *testing.T) {
	c := client(t, "one.marketapp.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestMarketplaceApplication_Update(t *testing.T) {
	c := client(t, "one.marketapp.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.MarketplaceApplication{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
