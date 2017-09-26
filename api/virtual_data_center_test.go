package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualDataCenter_AddCluster(t *testing.T) {
	c := client(t, "one.vdc.addcluster", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.AddCluster("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("AddCluster() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_AddDataStore(t *testing.T) {
	c := client(t, "one.vdc.adddatastore", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.AddDataStore("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("AddDataStore() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_AddGroup(t *testing.T) {
	c := client(t, "one.vdc.addgroup", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.AddGroup("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddGroup() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_AddHost(t *testing.T) {
	c := client(t, "one.vdc.addhost", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.AddHost("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("AddHost() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_AddVNet(t *testing.T) {
	c := client(t, "one.vdc.addvnet", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.AddVNet("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("AddVNet() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_Allocate(t *testing.T) {
	c := client(t, "one.vdc.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.Allocate("test-session", "test-template", 21)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_DelCluster(t *testing.T) {
	c := client(t, "one.vdc.delcluster", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.DelCluster("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("DelCluster() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_DelDataStore(t *testing.T) {
	c := client(t, "one.vdc.deldatastore", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.DelDataStore("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("DelDataStore() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_Delete(t *testing.T) {
	c := client(t, "one.vdc.delete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_DelGroup(t *testing.T) {
	c := client(t, "one.vdc.delgroup", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.DelGroup("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelGroup() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_DelHost(t *testing.T) {
	c := client(t, "one.vdc.delhost", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.DelHost("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("DelHost() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_DelVNet(t *testing.T) {
	c := client(t, "one.vdc.delvnet", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.DelVNet("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("DelVNet() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_Info(t *testing.T) {
	c := client(t, "one.vdc.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%v, %q), want (%v, %v)", v, err, "test-info", nil)
	}
}

func TestVirtualDataCenter_Rename(t *testing.T) {
	c := client(t, "one.vdc.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualDataCenter_Update(t *testing.T) {
	c := client(t, "one.vdc.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualDataCenter{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
