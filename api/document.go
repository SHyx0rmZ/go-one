package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Document struct {
	Client xmlrpc.Client
}

func (d *Document) Allocate(session string, template string, documentType int) (int, error) {
	v, err := d.Client.Call("one.document.allocate", session, template, documentType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Clone(session string, id int, name string) (int, error) {
	v, err := d.Client.Call("one.document.clone", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Delete(session string, id int) (int, error) {
	v, err := d.Client.Call("one.document.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.Client.Call("one.document.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := d.Client.Call("one.document.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := d.Client.Call("one.document.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Rename(session string, id int, name string) (int, error) {
	v, err := d.Client.Call("one.document.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Lock(session string, id int, application string) (bool, error) {
	v, err := d.Client.Call("one.document.lock", session, id, application)
	if err != nil {
		return false, err
	}

	return errorOrBool(v)
}

func (d *Document) Unlock(session string, id int, application string) (int, error) {
	v, err := d.Client.Call("one.document.unlock", session, id, application)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *Document) Info(session string, id int) (string, error) {
	v, err := d.Client.Call("one.document.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
