package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualRouter struct {
	Client xmlrpc.Client
}

func (r *VirtualRouter) Allocate(session string, template string) (int, error) {
	v, err := r.Client.Call("one.vrouter.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Delete(session string, id int) (int, error) {
	v, err := r.Client.Call("one.vrouter.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Instantiate(session string, id int, virtualMachineCount int, templateID int, onHold bool, extraTemplate string) (int, error) {
	v, err := r.Client.Call("one.vrouter.instantiate", session, id, virtualMachineCount, templateID, onHold, extraTemplate)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) AttachNIC(session string, id int, nicAttribute string) (int, error) {
	v, err := r.Client.Call("one.vrouter.attachnic", session, id, nicAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) DetachNIC(session string, id int, nicID int) (int, error) {
	v, err := r.Client.Call("one.vrouter.detachnic", session, id, nicID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := r.Client.Call("one.vrouter.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := r.Client.Call("one.vrouter.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := r.Client.Call("one.vrouter.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Rename(session string, id int, name string) (int, error) {
	v, err := r.Client.Call("one.vrouter.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *VirtualRouter) Info(session string, id int) (string, error) {
	v, err := r.Client.Call("one.vrouter.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
