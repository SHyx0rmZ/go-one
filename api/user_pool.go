package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type userPool struct {
	client xmlrpc.Client
}

func (p *userPool) info(session string) (string, error) {
	v, err := p.client.Call("one.userpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
