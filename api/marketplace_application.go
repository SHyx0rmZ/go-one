package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type MarketplaceApplication struct {
	Client xmlrpc.Client
}

func (m *MarketplaceApplication) Allocate(session string, template string, dataStoreID int) (int, error) {
	v, err := m.Client.Call("one.marketapp.allocate", session, template, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Delete(session string, id int) (int, error) {
	v, err := m.Client.Call("one.marketapp.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Enable(session string, id int, enable bool) (int, error) {
	v, err := m.Client.Call("one.marketapp.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.Client.Call("one.marketapp.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.Client.Call("one.marketapp.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.Client.Call("one.marketapp.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Rename(session string, id int, name string) (int, error) {
	v, err := m.Client.Call("one.marketapp.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *MarketplaceApplication) Info(session string, id int) (string, error) {
	v, err := m.Client.Call("one.marketapp.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
