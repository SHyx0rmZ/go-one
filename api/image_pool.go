package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type ImagePool struct {
	Client xmlrpc.Client
}

func (p *ImagePool) Info(session string, filter int, start int, end int) (string, error) {
	v, err := p.Client.Call("one.imagepool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
