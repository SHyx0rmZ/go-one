package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type virtualNetworkPool struct {
	client xmlrpc.Client
}

func (p *virtualNetworkPool) info(session string, filter int, start int, end int) (string, error) {
	v, err := p.client.Call("one.vnpool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
