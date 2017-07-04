package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type HostPool struct {
	Client xmlrpc.Client
}

func (p *HostPool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.hostpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *HostPool) Monitoring(session string) (string, error) {
	v, err := p.Client.Call("one.hostpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
