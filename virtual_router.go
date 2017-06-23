package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type virtualRouter struct {
	client xmlrpc.Client
}

func (r *virtualRouter) allocate(session string, template string) (int, error) {
	v, err := r.client.Call("one.vrouter.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) delete(session string, id int) (int, error) {
	v, err := r.client.Call("one.vrouter.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) instantiate(session string, id int, virtualMachineCount int, templateID int, onHold bool, extraTemplate string) (int, error) {
	v, err := r.client.Call("one.vrouter.instantiate", session, id, virtualMachineCount, templateID, onHold, extraTemplate)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) attachNIC(session string, id int, nicAttribute string) (int, error) {
	v, err := r.client.Call("one.vrouter.attachnic", session, id, nicAttribute)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) detachNIC(session string, id int, nicID int) (int, error) {
	v, err := r.client.Call("one.vrouter.detachnic", session, id, nicID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) update(session string, id int, template string, updateType int) (int, error) {
	v, err := r.client.Call("one.vrouter.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := r.client.Call("one.vrouter.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := r.client.Call("one.vrouter.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) rename(session string, id int, name string) (int, error) {
	v, err := r.client.Call("one.vrouter.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (r *virtualRouter) info(session string, id int) (string, error) {
	v, err := r.client.Call("one.vrouter.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
