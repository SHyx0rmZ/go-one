package one

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

type vm struct {
	client xmlrpc.Client
}

func (v *vm) allocate(session string, template string, onHold bool) (int, error) {
	value, err := v.client.Call("one.vm.allocate", session, template, onHold)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) deploy(session string, id int, hostID int, dontOvercommit bool, optionalDataStoreID int) (int, error) {
	value, err := v.client.Call("one.vm.deploy", session, id, hostID, dontOvercommit, optionalDataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) action(session string, action string, id int) (int, error) {
	value, err := v.client.Call("one.vm.action", session, action, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) migrate(session string, id int, hostID int, migrateLive bool, dontOvercommit bool, dataStoreID int) (int, error) {
	value, err := v.client.Call("one.vm.migrate", session, id, hostID, migrateLive, dontOvercommit, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) diskSaveAs(session string, id int, diskID int, name string, imageType string, optionalSnapshotID int) (int, error) {
	value, err := v.client.Call("one.vm.disksaveas", session, id, diskID, name, imageType, optionalSnapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) diskSnapshotCreate(session string, id int, diskID int, description string) (int, error) {
	value, err := v.client.Call("one.vm.disksnapshotcreate", session, id, diskID, description)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) diskSnapshotDelete(session string, id int, diskID int, snapshotID int) (int, error) {
	value, err := v.client.Call("one.vm.disksnapshotdelete", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) diskSnapshotRevert(session string, id int, diskID int, snapshotID int) (int, error) {
	value, err := v.client.Call("one.vm.disksnapshotrevert", session, id, diskID, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) attach(session string, id int, diskAttribute string) (int, error) {
	value, err := v.client.Call("one.vm.attach", session, id, diskAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) detach(session string, id int, diskID int) (int, error) {
	value, err := v.client.Call("one.vm.detach", session, id, diskID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) attachNIC(session string, id int, nicAttribute string) (int, error) {
	value, err := v.client.Call("one.vm.attachnic", session, id, nicAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) detachNIC(session string, id int, nicID int) (int, error) {
	value, err := v.client.Call("one.vm.detachnic", session, id, nicID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	value, err := v.client.Call("one.vm.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) chown(session string, id int, userID int, groupID int) (int, error) {
	value, err := v.client.Call("one.vm.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) rename(session string, id int, name string) (int, error) {
	value, err := v.client.Call("one.vm.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) snapshotCreate(session string, id int, name string) (int, error) {
	value, err := v.client.Call("one.vm.snapshotcreate", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) snapshotRevert(session string, id int, snapshotID int) (int, error) {
	value, err := v.client.Call("one.vm.snapshotrevert", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) snapshotDelete(session string, id int, snapshotID int) (int, error) {
	value, err := v.client.Call("one.vm.snapshotdelete", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) resize(session string, id int, template string, dontOvercommit bool) (int, error) {
	value, err := v.client.Call("one.vm.resize", session, id, template, dontOvercommit)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) update(session string, id int, template string, updateType int) (int, error) {
	value, err := v.client.Call("one.vm.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) updateConf(session string, id int, template string) (int, error) {
	value, err := v.client.Call("one.vm.updateconf", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) recoverVM(session string, id int, recoverOperation int) (int, error) {
	value, err := v.client.Call("one.vm.recover", session, id, recoverOperation)
	if err != nil {
		return 0, err
	}

	return errorOrInt(value)
}

func (v *vm) info(session string, id int) (string, error) {
	value, err := v.client.Call("one.vm.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(value)
}

func (v *vm) monitoring(session string, id int) (string, error) {
	value, err := v.client.Call("one.vm.monitoring", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(value)
}
