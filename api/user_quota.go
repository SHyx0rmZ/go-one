package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type UserQuota struct {
	Client xmlrpc.Client
}

func (u *UserQuota) Info(session string) (string, error) {
	v, err := u.Client.Call("one.userquota.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (u *UserQuota) Update(session string, template string) (string, error) {
	v, err := u.Client.Call("one.userquota.update", session, template)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
