package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Zone struct {
	Client xmlrpc.Client
}

func (z *Zone) Allocate(session string, template string) (int, error) {
	v, err := z.Client.Call("one.zone.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *Zone) Delete(session string, id int) (int, error) {
	v, err := z.Client.Call("one.zone.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *Zone) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := z.Client.Call("one.zone.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *Zone) Rename(session string, id int, name string) (int, error) {
	v, err := z.Client.Call("one.zone.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *Zone) Info(session string, id int) (string, error) {
	v, err := z.Client.Call("one.zone.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
