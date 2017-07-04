package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	UpdateTypeReplace = 0
	UpdateTypeMerge   = 1

	PermissionDontChange = -1
	PermissionDeny       = 0
	PermissionAllow      = 1
)

type Template struct {
	Client xmlrpc.Client
}

func (t *Template) Allocate(session string, template string) (int, error) {
	v, err := t.Client.Call("one.template.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Clone(session string, id int, name string, withImage bool) (int, error) {
	v, err := t.Client.Call("one.template.clone", session, id, name, withImage)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Delete(session string, id int, withImage bool) (int, error) {
	v, err := t.Client.Call("one.template.delete", session, id, withImage)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Instantiate(session string, id int, name string, onHold bool, extraTemplate string, persistentCopy bool) (int, error) {
	v, err := t.Client.Call("one.template.instantiate", session, id, name, onHold, extraTemplate, persistentCopy)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := t.Client.Call("one.template.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int, withImage bool) (int, error) {
	v, err := t.Client.Call("one.template.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin, withImage)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := t.Client.Call("one.template.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Rename(session string, id int, name string) (int, error) {
	v, err := t.Client.Call("one.template.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (t *Template) Info(session string, id int, withDetails bool) (string, error) {
	v, err := t.Client.Call("one.template.info", session, id, withDetails)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
