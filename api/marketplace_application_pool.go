package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type marketplaceApplicationPool struct {
	client xmlrpc.Client
}

func (p *marketplaceApplicationPool) info(session string, filter int, start int, end int) (string, error) {
	v, err := p.client.Call("one.marketapppool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
