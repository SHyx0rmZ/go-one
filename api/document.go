package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type document struct {
	client xmlrpc.Client
}

func (d *document) allocate(session string, template string, documentType int) (int, error) {
	v, err := d.client.Call("one.document.allocate", session, template, documentType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) clone(session string, id int, name string) (int, error) {
	v, err := d.client.Call("one.document.clone", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) delete(session string, id int) (int, error) {
	v, err := d.client.Call("one.document.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.client.Call("one.document.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := d.client.Call("one.document.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := d.client.Call("one.document.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) rename(session string, id int, name string) (int, error) {
	v, err := d.client.Call("one.document.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) lock(session string, id int, application string) (bool, error) {
	v, err := d.client.Call("one.document.lock", session, id, application)
	if err != nil {
		return false, err
	}

	return errorOrBool(v)
}

func (d *document) unlock(session string, id int, application string) (int, error) {
	v, err := d.client.Call("one.document.unlock", session, id, application)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *document) info(session string, id int) (string, error) {
	v, err := d.client.Call("one.document.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
