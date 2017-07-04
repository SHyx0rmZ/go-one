package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type MarketplacePool struct {
	Client xmlrpc.Client
}

func (p *MarketplacePool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.marketplacepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
