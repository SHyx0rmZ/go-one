package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type imagePool struct {
	client xmlrpc.Client
}

func (p *imagePool) info(session string, filter int, start int, end int) (string, error) {
	v, err := p.client.Call("one.imagepool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
