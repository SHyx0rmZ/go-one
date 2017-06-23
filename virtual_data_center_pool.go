package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type virtualDataCenterPool struct {
	client xmlrpc.Client
}

func (p *virtualDataCenterPool) info(session string) (string, error) {
	v, err := p.client.Call("one.vdcpool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
