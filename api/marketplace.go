package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Marketplace struct {
	Client xmlrpc.Client
}

func (m *Marketplace) Allocate(session string, template string) (int, error) {
	v, err := m.Client.Call("one.marketplace.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Delete(session string, id int) (int, error) {
	v, err := m.Client.Call("one.marketplace.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.Client.Call("one.marketplace.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.Client.Call("one.marketplace.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.Client.Call("one.marketplace.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Rename(session string, id int, name string) (int, error) {
	v, err := m.Client.Call("one.marketplace.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *Marketplace) Info(session string, id int) (string, error) {
	v, err := m.Client.Call("one.marketplace.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
