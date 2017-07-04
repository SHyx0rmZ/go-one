package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type GroupQuota struct {
	Client xmlrpc.Client
}

func (g *GroupQuota) Info(session string) (string, error) {
	v, err := g.Client.Call("one.groupquota.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (g *GroupQuota) Update(session string, template string) (string, error) {
	v, err := g.Client.Call("one.groupquota.update", session, template)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
