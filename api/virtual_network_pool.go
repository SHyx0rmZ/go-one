package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualNetworkPool struct {
	Client xmlrpc.Client
}

func (p *VirtualNetworkPool) Info(session string, filter int, start int, end int) (string, error) {
	v, err := p.Client.Call("one.vnpool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
