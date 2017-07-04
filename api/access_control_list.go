package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type accessControlList struct {
	client xmlrpc.Client
}

func (a *accessControlList) addRule(session string, user string, resource string, rights string) (int, error) {
	v, err := a.client.Call("one.acl.addrule", session, user, resource, rights)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (a *accessControlList) delRule(session string, id int) (int, error) {
	v, err := a.client.Call("one.acl.delrule", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (a *accessControlList) info(session string, id int) (string, error) {
	v, err := a.client.Call("one.acl.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
