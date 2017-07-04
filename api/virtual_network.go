package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualNetwork struct {
	Client xmlrpc.Client
}

func (n *VirtualNetwork) Allocate(session string, template string, clusterID int) (int, error) {
	v, err := n.Client.Call("one.vn.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Delete(session string, id int) (int, error) {
	v, err := n.Client.Call("one.vn.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) AddAR(session string, id int, template string) (int, error) {
	v, err := n.Client.Call("one.vn.add_ar", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) RmAR(session string, id int, addressRangeID int) (int, error) {
	v, err := n.Client.Call("one.vn.rm_ar", session, id, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) UpdateAR(session string, addressRangeID int) (int, error) {
	v, err := n.Client.Call("one.vn.update_ar", session, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Reserve(session string, id int, template string) (int, error) {
	v, err := n.Client.Call("one.vn.reserve", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) FreeAR(session string, id int, addressRangeID int) (int, error) {
	v, err := n.Client.Call("one.vn.free_ar", session, id, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Hold(session string, id int, template string) (int, error) {
	v, err := n.Client.Call("one.vn.hold", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Release(session string, id int, template string) (int, error) {
	v, err := n.Client.Call("one.vn.release", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := n.Client.Call("one.vn.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := n.Client.Call("one.vn.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := n.Client.Call("one.vn.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Rename(session string, id int, name string) (int, error) {
	v, err := n.Client.Call("one.vn.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *VirtualNetwork) Info(session string, id int) (string, error) {
	v, err := n.Client.Call("one.vn.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
