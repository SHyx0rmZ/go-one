package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type userQuota struct {
	client xmlrpc.Client
}

func (u *userQuota) info(session string) (string, error) {
	v, err := u.client.Call("one.userquota.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (u *userQuota) update(session string, template string) (string, error) {
	v, err := u.client.Call("one.userquota.update", session, template)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
