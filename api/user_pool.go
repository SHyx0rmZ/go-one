package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type UserPool struct {
	Client xmlrpc.Client
}

func (p *UserPool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.userpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
