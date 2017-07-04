package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type System struct {
	Client xmlrpc.Client
}

func (s *System) Version(session string) (string, error) {
	v, err := s.Client.Call("one.system.version", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (s *System) Config(session string) (string, error) {
	v, err := s.Client.Call("one.system.config", session)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
