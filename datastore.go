package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type datastore struct {
	client xmlrpc.Client
}

func (d *datastore) allocate(session string, template string, clusterID int) (int, error) {
	v, err := d.client.Call("one.datastore.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) delete(session string, id int) (int, error) {
	v, err := d.client.Call("one.datastore.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.client.Call("one.datastore.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := d.client.Call("one.datastore.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := d.client.Call("one.datastore.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) rename(session string, id int, name string) (int, error) {
	v, err := d.client.Call("one.datastore.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) enable(session string, id int, enable bool) (int, error) {
	v, err := d.client.Call("one.datastore.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *datastore) info(session string, id int) (string, error) {
	v, err := d.client.Call("one.datastore.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
