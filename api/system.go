package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type system struct {
	client xmlrpc.Client
}

func (s *system) version(session string) (string, error) {
	v, err := s.client.Call("one.system.version", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (s *system) config(session string) (string, error) {
	v, err := s.client.Call("one.system.config", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
