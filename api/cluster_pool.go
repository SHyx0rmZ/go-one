package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type ClusterPool struct {
	Client xmlrpc.Client
}

func (p *ClusterPool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.clusterpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
