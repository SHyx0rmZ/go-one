package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type DocumentPool struct {
	Client xmlrpc.Client
}

func (p *DocumentPool) Info(session string, filter int, start int, end int, documentType int) (string, error) {
	v, err := p.Client.Call("one.documentpool.info", session, filter, start, end, documentType)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
