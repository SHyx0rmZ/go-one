package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type zone struct {
	client xmlrpc.Client
}

func (z *zone) allocate(session string, template string) (int, error) {
	v, err := z.client.Call("one.zone.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *zone) delete(session string, id int) (int, error) {
	v, err := z.client.Call("one.zone.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *zone) update(session string, id int, template string, updateType int) (int, error) {
	v, err := z.client.Call("one.zone.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *zone) rename(session string, id int, name string) (int, error) {
	v, err := z.client.Call("one.zone.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (z *zone) info(session string, id int) (string, error) {
	v, err := z.client.Call("one.zone.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
