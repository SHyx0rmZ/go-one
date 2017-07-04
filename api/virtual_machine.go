package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	ActionHold          = "hold"
	ActionPoweroff      = "poweroff"
	ActionPoweroffHard  = "poweroff-hard"
	ActionReboot        = "reboot"
	ActionRebootHard    = "reboot-hard"
	ActionRelease       = "release"
	ActionReschedule    = "resched"
	ActionResume        = "resume"
	ActionStop          = "stop"
	ActionSuspend       = "suspend"
	ActionTerminate     = "terminate"
	ActionTerminateHard = "terminate-hard"
	ActionUndeploy      = "undeploy"
	ActionUndeployHard  = "undeploy-hard"
	ActionUnreschedule  = "unresched"

	RecoverOperationFailure           = 0
	RecoverOperationSuccess           = 1
	RecoverOperationRetry             = 2
	RecoverOperationDelete            = 3
	RecoverOperationDeleteAndRecreate = 4
)

type VirtualMachine struct {
	Client xmlrpc.Client
}

func (m *VirtualMachine) Allocate(session string, template string, onHold bool) (int, error) {
	v, err := m.Client.Call("one.vm.allocate", session, template, onHold)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Deploy(session string, id int, hostID int, dontOvercommit bool, optionalDataStoreID int) (int, error) {
	v, err := m.Client.Call("one.vm.deploy", session, id, hostID, dontOvercommit, optionalDataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Action(session string, action string, id int) (int, error) {
	v, err := m.Client.Call("one.vm.action", session, action, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Migrate(session string, id int, hostID int, migrateLive bool, dontOvercommit bool, dataStoreID int) (int, error) {
	v, err := m.Client.Call("one.vm.migrate", session, id, hostID, migrateLive, dontOvercommit, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) DiskSaveAs(session string, id int, diskID int, name string, imageType string, optionalSnapshotID int) (int, error) {
	v, err := m.Client.Call("one.vm.disksaveas", session, id, diskID, name, imageType, optionalSnapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) DiskSnapshotCreate(session string, id int, diskID int, description string) (int, error) {
	v, err := m.Client.Call("one.vm.disksnapshotcreate", session, id, diskID, description)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) DiskSnapshotDelete(session string, id int, diskID int, snapshotID int) (int, error) {
	v, err := m.Client.Call("one.vm.disksnapshotdelete", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) DiskSnapshotRevert(session string, id int, diskID int, snapshotID int) (int, error) {
	v, err := m.Client.Call("one.vm.disksnapshotrevert", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Attach(session string, id int, diskAttribute string) (int, error) {
	v, err := m.Client.Call("one.vm.attach", session, id, diskAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Detach(session string, id int, diskID int) (int, error) {
	v, err := m.Client.Call("one.vm.detach", session, id, diskID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) AttachNIC(session string, id int, nicAttribute string) (int, error) {
	v, err := m.Client.Call("one.vm.attachnic", session, id, nicAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) DetachNIC(session string, id int, nicID int) (int, error) {
	v, err := m.Client.Call("one.vm.detachnic", session, id, nicID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.Client.Call("one.vm.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.Client.Call("one.vm.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Rename(session string, id int, name string) (int, error) {
	v, err := m.Client.Call("one.vm.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) SnapshotCreate(session string, id int, name string) (int, error) {
	v, err := m.Client.Call("one.vm.snapshotcreate", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) SnapshotRevert(session string, id int, snapshotID int) (int, error) {
	v, err := m.Client.Call("one.vm.snapshotrevert", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) SnapshotDelete(session string, id int, snapshotID int) (int, error) {
	v, err := m.Client.Call("one.vm.snapshotdelete", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Resize(session string, id int, template string, dontOvercommit bool) (int, error) {
	v, err := m.Client.Call("one.vm.resize", session, id, template, dontOvercommit)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.Client.Call("one.vm.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) UpdateConf(session string, id int, template string) (int, error) {
	v, err := m.Client.Call("one.vm.updateconf", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Recover(session string, id int, recoverOperation int) (int, error) {
	v, err := m.Client.Call("one.vm.recover", session, id, recoverOperation)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *VirtualMachine) Info(session string, id int) (string, error) {
	v, err := m.Client.Call("one.vm.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (m *VirtualMachine) Monitoring(session string, id int) (string, error) {
	v, err := m.Client.Call("one.vm.monitoring", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
