package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type GroupPool struct {
	Client xmlrpc.Client
}

func (p *GroupPool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.grouppool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
