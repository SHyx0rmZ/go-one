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

type virtualMachine struct {
	client xmlrpc.Client
}

func (m *virtualMachine) allocate(session string, template string, onHold bool) (int, error) {
	v, err := m.client.Call("one.vm.allocate", session, template, onHold)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) deploy(session string, id int, hostID int, dontOvercommit bool, optionalDataStoreID int) (int, error) {
	v, err := m.client.Call("one.vm.deploy", session, id, hostID, dontOvercommit, optionalDataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) action(session string, action string, id int) (int, error) {
	v, err := m.client.Call("one.vm.action", session, action, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) migrate(session string, id int, hostID int, migrateLive bool, dontOvercommit bool, dataStoreID int) (int, error) {
	v, err := m.client.Call("one.vm.migrate", session, id, hostID, migrateLive, dontOvercommit, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) diskSaveAs(session string, id int, diskID int, name string, imageType string, optionalSnapshotID int) (int, error) {
	v, err := m.client.Call("one.vm.disksaveas", session, id, diskID, name, imageType, optionalSnapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) diskSnapshotCreate(session string, id int, diskID int, description string) (int, error) {
	v, err := m.client.Call("one.vm.disksnapshotcreate", session, id, diskID, description)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) diskSnapshotDelete(session string, id int, diskID int, snapshotID int) (int, error) {
	v, err := m.client.Call("one.vm.disksnapshotdelete", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) diskSnapshotRevert(session string, id int, diskID int, snapshotID int) (int, error) {
	v, err := m.client.Call("one.vm.disksnapshotrevert", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) attach(session string, id int, diskAttribute string) (int, error) {
	v, err := m.client.Call("one.vm.attach", session, id, diskAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) detach(session string, id int, diskID int) (int, error) {
	v, err := m.client.Call("one.vm.detach", session, id, diskID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) attachNIC(session string, id int, nicAttribute string) (int, error) {
	v, err := m.client.Call("one.vm.attachnic", session, id, nicAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) detachNIC(session string, id int, nicID int) (int, error) {
	v, err := m.client.Call("one.vm.detachnic", session, id, nicID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.client.Call("one.vm.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.client.Call("one.vm.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) rename(session string, id int, name string) (int, error) {
	v, err := m.client.Call("one.vm.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) snapshotCreate(session string, id int, name string) (int, error) {
	v, err := m.client.Call("one.vm.snapshotcreate", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) snapshotRevert(session string, id int, snapshotID int) (int, error) {
	v, err := m.client.Call("one.vm.snapshotrevert", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) snapshotDelete(session string, id int, snapshotID int) (int, error) {
	v, err := m.client.Call("one.vm.snapshotdelete", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) resize(session string, id int, template string, dontOvercommit bool) (int, error) {
	v, err := m.client.Call("one.vm.resize", session, id, template, dontOvercommit)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.client.Call("one.vm.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) updateConf(session string, id int, template string) (int, error) {
	v, err := m.client.Call("one.vm.updateconf", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) recover(session string, id int, recoverOperation int) (int, error) {
	v, err := m.client.Call("one.vm.recover", session, id, recoverOperation)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *virtualMachine) info(session string, id int) (string, error) {
	v, err := m.client.Call("one.vm.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (m *virtualMachine) monitoring(session string, id int) (string, error) {
	v, err := m.client.Call("one.vm.monitoring", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
