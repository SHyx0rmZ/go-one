package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type clusterPool struct {
	client xmlrpc.Client
}

func (p *clusterPool) info(session string) (string, error) {
	v, err := p.client.Call("one.clusterpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
