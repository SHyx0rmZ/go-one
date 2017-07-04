package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualRouterPool struct {
	Client xmlrpc.Client
}

func (p *VirtualRouterPool) Info(session string, filter int, start int, end int) (string, error) {
	v, err := p.Client.Call("one.vrouterpool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
