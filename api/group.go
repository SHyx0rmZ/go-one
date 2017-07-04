package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Group struct {
	Client xmlrpc.Client
}

func (g *Group) Allocate(session string, name string) (int, error) {
	v, err := g.Client.Call("one.group.allocate", session, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *Group) Delete(session string, id int) (int, error) {
	v, err := g.Client.Call("one.group.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *Group) Info(session string, id int) (string, error) {
	v, err := g.Client.Call("one.group.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (g *Group) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := g.Client.Call("one.group.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *Group) AddAdmin(session string, id int, userID int) (int, error) {
	v, err := g.Client.Call("one.group.addadmin", session, id, userID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *Group) DelAdmin(session string, id int, userID int) (int, error) {
	v, err := g.Client.Call("one.group.deladmin", session, id, userID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (g *Group) Quota(session string, id int, template string) (int, error) {
	v, err := g.Client.Call("one.group.quota", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}
