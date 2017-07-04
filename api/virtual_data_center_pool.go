package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualDataCenterPool struct {
	Client xmlrpc.Client
}

func (p *VirtualDataCenterPool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.vdcpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
