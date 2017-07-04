package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type datastorePool struct {
	client xmlrpc.Client
}

func (p *datastorePool) info(session string) (string, error) {
	v, err := p.client.Call("one.datastorepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
