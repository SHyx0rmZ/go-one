package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestImage_Allocate(t *testing.T) {
	c := client(t, "one.image.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Allocate("test-session", "test-template", 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Chmod(t *testing.T) {
	c := client(t, "one.image.chmod", []arg{
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

	i := api.Image{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Chown(t *testing.T) {
	c := client(t, "one.image.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Chtype(t *testing.T) {
	c := client(t, "one.image.chtype", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, api.ImageTypeRAMDisk},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Chtype("test-session", 42, api.ImageTypeRAMDisk)
	if err != nil || v != 42 {
		t.Errorf("Chtype() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Clone(t *testing.T) {
	c := client(t, "one.image.clone", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(13),
	})

	i := api.Image{Client: c}
	v, err := i.Clone("test-session", 42, "test-name", 21)
	if err != nil || v != 13 {
		t.Errorf("Clone() == (%v, %q), want (%v, %v)", v, err, 13, nil)
	}
}

func TestImage_Delete(t *testing.T) {
	c := client(t, "one.image.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Enable(t *testing.T) {
	c := client(t, "one.image.enable", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Enable("test-session", 42, true)
	if err != nil || v != 42 {
		t.Errorf("Enable() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Info(t *testing.T) {
	c := client(t, "one.image.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.Image{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestImage_Persistent(t *testing.T) {
	c := client(t, "one.image.persistent", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Persistent("test-session", 42, true)
	if err != nil || v != 42 {
		t.Errorf("Persistent() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_Rename(t *testing.T) {
	c := client(t, "one.image.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestImage_SnapshotDelete(t *testing.T) {
	c := client(t, "one.image.snapshotdelete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.Image{Client: c}
	v, err := i.SnapshotDelete("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("SnapshotDelete() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestImage_SnapshotFlatten(t *testing.T) {
	c := client(t, "one.image.snapshotflatten", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.Image{Client: c}
	v, err := i.SnapshotFlatten("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("SnapshotFlatten() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestImage_SnapshotRevert(t *testing.T) {
	c := client(t, "one.image.snapshotrevert", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.Image{Client: c}
	v, err := i.SnapshotRevert("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("SnapshotRevert() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestImage_Update(t *testing.T) {
	c := client(t, "one.image.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.Image{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
