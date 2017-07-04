package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Datastore struct {
	Client xmlrpc.Client
}

func (d *Datastore) Allocate(session string, template string, clusterID int) (int, error) {
	v, err := d.Client.Call("one.datastore.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Delete(session string, id int) (int, error) {
	v, err := d.Client.Call("one.datastore.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.Client.Call("one.datastore.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := d.Client.Call("one.datastore.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := d.Client.Call("one.datastore.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Rename(session string, id int, name string) (int, error) {
	v, err := d.Client.Call("one.datastore.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Enable(session string, id int, enable bool) (int, error) {
	v, err := d.Client.Call("one.datastore.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Datastore) Info(session string, id int) (string, error) {
	v, err := d.Client.Call("one.datastore.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
