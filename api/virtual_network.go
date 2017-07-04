package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type virtualNetwork struct {
	client xmlrpc.Client
}

func (n *virtualNetwork) allocate(session string, template string, clusterID int) (int, error) {
	v, err := n.client.Call("one.vn.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) delete(session string, id int) (int, error) {
	v, err := n.client.Call("one.vn.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) addAR(session string, id int, template string) (int, error) {
	v, err := n.client.Call("one.vn.add_ar", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) rmAR(session string, id int, addressRangeID int) (int, error) {
	v, err := n.client.Call("one.vn.rm_ar", session, id, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) updateAR(session string, addressRangeID int) (int, error) {
	v, err := n.client.Call("one.vn.update_ar", session, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) reserve(session string, id int, template string) (int, error) {
	v, err := n.client.Call("one.vn.reserve", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) freeAR(session string, id int, addressRangeID int) (int, error) {
	v, err := n.client.Call("one.vn.free_ar", session, id, addressRangeID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) hold(session string, id int, template string) (int, error) {
	v, err := n.client.Call("one.vn.hold", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) release(session string, id int, template string) (int, error) {
	v, err := n.client.Call("one.vn.release", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) update(session string, id int, template string, updateType int) (int, error) {
	v, err := n.client.Call("one.vn.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := n.client.Call("one.vn.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := n.client.Call("one.vn.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) rename(session string, id int, name string) (int, error) {
	v, err := n.client.Call("one.vn.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (n *virtualNetwork) info(session string, id int) (string, error) {
	v, err := n.client.Call("one.vn.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
