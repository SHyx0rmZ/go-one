package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type hostPool struct {
	client xmlrpc.Client
}

func (p *hostPool) info(session string) (string, error) {
	v, err := p.client.Call("one.hostpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *hostPool) monitoring(session string) (string, error) {
	v, err := p.client.Call("one.hostpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
