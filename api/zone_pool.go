package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type ZonePool struct {
	Client xmlrpc.Client
}

func (p *ZonePool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.zonepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
