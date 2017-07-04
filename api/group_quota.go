package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type groupQuota struct {
	client xmlrpc.Client
}

func (g *groupQuota) info(session string) (string, error) {
	v, err := g.client.Call("one.groupquota.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (g *groupQuota) update(session string, template string) (string, error) {
	v, err := g.client.Call("one.groupquota.update", session, template)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
