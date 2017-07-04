package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type groupPool struct {
	client xmlrpc.Client
}

func (p *groupPool) info(session string) (string, error) {
	v, err := p.client.Call("one.grouppool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
