package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type group struct {
	client xmlrpc.Client
}

func (g *group) allocate(session string, name string) (int, error) {
	v, err := g.client.Call("one.group.allocate", session, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *group) delete(session string, id int) (int, error) {
	v, err := g.client.Call("one.group.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *group) info(session string, id int) (string, error) {
	v, err := g.client.Call("one.group.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (g *group) update(session string, id int, template string, updateType int) (int, error) {
	v, err := g.client.Call("one.group.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *group) addAdmin(session string, id int, userID int) (int, error) {
	v, err := g.client.Call("one.group.addadmin", session, id, userID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *group) delAdmin(session string, id int, userID int) (int, error) {
	v, err := g.client.Call("one.group.deladmin", session, id, userID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *group) quota(session string, id int, template string) (int, error) {
	v, err := g.client.Call("one.group.quota", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}
