package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type documentPool struct {
	client xmlrpc.Client
}

func (p *documentPool) info(session string, filter int, start int, end int, documentType int) (string, error) {
	v, err := p.client.Call("one.documentpool.info", session, filter, start, end, documentType)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
