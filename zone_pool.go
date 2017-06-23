package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type zonePool struct {
	client xmlrpc.Client
}

func (p *zonePool) info(session string) (string, error) {
	v, err := p.client.Call("one.zonepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
