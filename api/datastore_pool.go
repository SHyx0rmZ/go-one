package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type DatastorePool struct {
	Client xmlrpc.Client
}

func (p *DatastorePool) Info(session string) (string, error) {
	v, err := p.Client.Call("one.datastorepool.info", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
