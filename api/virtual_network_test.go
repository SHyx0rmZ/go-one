package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualNetwork_AddAR(t *testing.T) {
	c := client(t, "one.vn.add_ar", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.AddAR("test-session", 42, "test-template")
	if err != nil || v != 21 {
		t.Errorf("AddAR() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualNetwork_Allocate(t *testing.T) {
	c := client(t, "one.vn.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Allocate("test-session", "test-template", 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Chmod(t *testing.T) {
	c := client(t, "one.vn.chmod", []arg{
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

	i := api.VirtualNetwork{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Chown(t *testing.T) {
	c := client(t, "one.vn.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Delete(t *testing.T) {
	c := client(t, "one.vn.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_FreeAR(t *testing.T) {
	c := client(t, "one.vn.free_ar", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.FreeAR("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("FreeAR() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualNetwork_Hold(t *testing.T) {
	c := client(t, "one.vn.hold", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Hold("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("Hold() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Info(t *testing.T) {
	c := client(t, "one.vn.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%v, %q), want (%v, %v)", v, err, "test-info", nil)
	}
}

func TestVirtualNetwork_Release(t *testing.T) {
	c := client(t, "one.vn.release", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Release("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("Release() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Rename(t *testing.T) {
	c := client(t, "one.vn.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Reserve(t *testing.T) {
	c := client(t, "one.vn.reserve", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Reserve("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("Reserve() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_RmAR(t *testing.T) {
	c := client(t, "one.vn.rm_ar", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.RmAR("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("RmAR() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_Update(t *testing.T) {
	c := client(t, "one.vn.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualNetwork_UpdateAR(t *testing.T) {
	c := client(t, "one.vn.update_ar", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 21},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualNetwork{Client: c}
	v, err := i.UpdateAR("test-session", 21, "test-template")
	if err != nil || v != 21 {
		t.Errorf("UpdateAR() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}
