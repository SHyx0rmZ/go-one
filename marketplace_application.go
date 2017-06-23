package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type marketplaceApplication struct {
	client xmlrpc.Client
}

func (m *marketplaceApplication) allocate(session string, template string, dataStoreID int) (int, error) {
	v, err := m.client.Call("one.marketapp.allocate", session, template, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) deleteMarketplaceApplication(session string, id int) (int, error) {
	v, err := m.client.Call("one.marketapp.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) enable(session string, id int, enable bool) (int, error) {
	v, err := m.client.Call("one.marketapp.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) update(session string, id int, template string, updateType int) (int, error) {
	v, err := m.client.Call("one.marketapp.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := m.client.Call("one.marketapp.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := m.client.Call("one.marketapp.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) rename(session string, id int, name string) (int, error) {
	v, err := m.client.Call("one.marketapp.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (m *marketplaceApplication) info(session string, id int) (string, error) {
	v, err := m.client.Call("one.marketapp.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
