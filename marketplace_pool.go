package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type marketplacePool struct {
	client xmlrpc.Client
}

func (p *marketplacePool) info(session string) (string, error) {
	v, err := p.client.Call("one.marketplacepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
