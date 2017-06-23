package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type marketplace struct {
	client xmlrpc.Client
}

func (m *marketplace) allocate(session string, template string) (int, error) {
	v, err := m.client.Call("one.marketplace.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) deleteMarketplace(session string, id int) (int, error) {
	v, err := m.client.Call("one.marketplace.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.client.Call("one.marketplace.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.client.Call("one.marketplace.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.client.Call("one.marketplace.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) rename(session string, id int, name string) (int, error) {
	v, err := m.client.Call("one.marketplace.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplace) info(session string, id int) (string, error) {
	v, err := m.client.Call("one.marketplace.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
