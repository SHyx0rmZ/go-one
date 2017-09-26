package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestVirtualMachine_Action(t *testing.T) {
	c := client(t, "one.vm.action", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-action"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Action("test-session", "test-action", 42)
	if err != nil || v != 42 {
		t.Errorf("Action() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Allocate(t *testing.T) {
	c := client(t, "one.vm.allocate", []arg{
		{reflect.String, "test-session"},
		{reflect.String, "test-template"},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Allocate("test-session", "test-template", true)
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Attach(t *testing.T) {
	c := client(t, "one.vm.attach", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-attribute"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Attach("test-session", 42, "test-attribute")
	if err != nil || v != 42 {
		t.Errorf("Attach() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_AttachNIC(t *testing.T) {
	c := client(t, "one.vm.attachnic", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-attribute"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.AttachNIC("test-session", 42, "test-attribute")
	if err != nil || v != 42 {
		t.Errorf("AttachNIC() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Chmod(t *testing.T) {
	c := client(t, "one.vm.chmod", []arg{
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

	i := api.VirtualMachine{Client: c}
	v, err := i.Chmod("test-session", 42, 1, 1, 1, 0, 0, 0, -1, -1, -1)
	if err != nil || v != 42 {
		t.Errorf("Chmod() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Chown(t *testing.T) {
	c := client(t, "one.vm.chown", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Chown("test-session", 42, 21, 21)
	if err != nil || v != 42 {
		t.Errorf("Chown() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Deploy(t *testing.T) {
	c := client(t, "one.vm.deploy", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Bool, true},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Deploy("test-session", 42, 21, true, 21)
	if err != nil || v != 42 {
		t.Errorf("Deploy() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Detach(t *testing.T) {
	c := client(t, "one.vm.detach", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Detach("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("Detach() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_DetachNIC(t *testing.T) {
	c := client(t, "one.vm.detachnic", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.DetachNIC("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DetachNIC() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_DiskSaveAs(t *testing.T) {
	c := client(t, "one.vm.disksaveas", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.String, "test-name"},
		{reflect.String, api.ImageTypeRAMDisk},
		{reflect.Int, -1},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.DiskSaveAs("test-session", 42, 21, "test-name", api.ImageTypeRAMDisk, -1)
	if err != nil || v != 42 {
		t.Errorf("DiskSaveAs() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_DiskSnapshotCreate(t *testing.T) {
	c := client(t, "one.vm.disksnapshotcreate", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.String, "test-description"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(13),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.DiskSnapshotCreate("test-session", 42, 21, "test-description")
	if err != nil || v != 13 {
		t.Errorf("DiskSnapshotCreate() == (%v, %q), want (%v, %v)", v, err, 13, nil)
	}
}

func TestVirtualMachine_DiskSnapshotDelete(t *testing.T) {
	c := client(t, "one.vm.disksnapshotdelete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.DiskSnapshotDelete("test-session", 42, 21, 21)
	if err != nil || v != 21 {
		t.Errorf("DiskSnapshotDelete() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualMachine_DiskSnapshotRevert(t *testing.T) {
	c := client(t, "one.vm.disksnapshotrevert", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.DiskSnapshotRevert("test-session", 42, 21, 21)
	if err != nil || v != 21 {
		t.Errorf("DiskSnapshotRevert() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualMachine_Info(t *testing.T) {
	c := client(t, "one.vm.info", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %q), want (%q, %v)", v, err, "test-info", nil)
	}
}

func TestVirtualMachine_Migrate(t *testing.T) {
	c := client(t, "one.vm.migrate", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
		{reflect.Bool, true},
		{reflect.Bool, true},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Migrate("test-session", 42, 21, true, true, 21)
	if err != nil || v != 42 {
		t.Errorf("Migrate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Monitoring(t *testing.T) {
	c := client(t, "one.vm.monitoring", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-monitor"),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Monitoring("test-session", 42)
	if err != nil || v != "test-monitor" {
		t.Errorf("Monitoring() == (%q, %q), want (%q, %v)", v, err, "test-monitor", nil)
	}
}

func TestVirtualMachine_Recover(t *testing.T) {
	c := client(t, "one.vm.recover", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, api.RecoverOperationDeleteAndRecreate},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Recover("test-session", 42, api.RecoverOperationDeleteAndRecreate)
	if err != nil || v != 42 {
		t.Errorf("Recover() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Rename(t *testing.T) {
	c := client(t, "one.vm.rename", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_Resize(t *testing.T) {
	c := client(t, "one.vm.resize", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Bool, true},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Resize("test-session", 42, "test-template", true)
	if err != nil || v != 42 {
		t.Errorf("Resize() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_SnapshotCreate(t *testing.T) {
	c := client(t, "one.vm.snapshotcreate", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-name"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.SnapshotCreate("test-session", 42, "test-name")
	if err != nil || v != 21 {
		t.Errorf("SnapshotCreate() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualMachine_SnapshotDelete(t *testing.T) {
	c := client(t, "one.vm.snapshotdelete", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.SnapshotDelete("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("SnapshotDelete() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualMachine_SnapshotRevert(t *testing.T) {
	c := client(t, "one.vm.snapshotrevert", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.Int, 21},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(21),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.SnapshotRevert("test-session", 42, 21)
	if err != nil || v != 21 {
		t.Errorf("SnapshotRevert() == (%v, %q), want (%v, %v)", v, err, 21, nil)
	}
}

func TestVirtualMachine_Update(t *testing.T) {
	c := client(t, "one.vm.update", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
		{reflect.Int, api.UpdateTypeMerge},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestVirtualMachine_UpdateConf(t *testing.T) {
	c := client(t, "one.vm.updateconf", []arg{
		{reflect.String, "test-session"},
		{reflect.Int, 42},
		{reflect.String, "test-template"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithInt(42),
	})

	i := api.VirtualMachine{Client: c}
	v, err := i.UpdateConf("test-session", 42, "test-template")
	if err != nil || v != 42 {
		t.Errorf("UpdateConf() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
