package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type AccessControlList struct {
	Client xmlrpc.Client
}

func (a *AccessControlList) AddRule(session string, user string, resource string, rights string) (int, error) {
	v, err := a.Client.Call("one.acl.addrule", session, user, resource, rights)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (a *AccessControlList) DelRule(session string, id int) (int, error) {
	v, err := a.Client.Call("one.acl.delrule", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (a *AccessControlList) Info(session string, id int) (string, error) {
	v, err := a.Client.Call("one.acl.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
