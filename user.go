package one

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	authenticationDriverCore = "core"
)

type user struct {
	client xmlrpc.Client
}

func (u *user) allocate(session string, username string, password string, authenticationDriver string, groupIDs ...int) (int, error) {
	v, err := u.client.Call("one.user.allocate", session, username, password, authenticationDriver, groupIDs)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) delete(session string, id int) (int, error) {
	v, err := u.client.Call("one.user.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) password(session string, id int, password string) (int, error) {
	v, err := u.client.Call("one.user.password", session, id, password)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) login(session string, username string, token string, ttlInSeconds int, effectiveGroupID int) (string, error) {
	v, err := u.client.Call("one.user.login", session, username, token, ttlInSeconds, effectiveGroupID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (u *user) update(session string, id int, template string, updateType int) (int, error) {
	v, err := u.client.Call("one.user.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) chauth(session string, id int, authenticationDriver string, optionalNewPassword string) (int, error) {
	v, err := u.client.Call("one.user.chauth", session, id, authenticationDriver, optionalNewPassword)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) quota(session string, id int, template string) (int, error) {
	v, err := u.client.Call("one.user.quota", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) chgrp(session string, id int, groupID int) (int, error) {
	v, err := u.client.Call("one.user.chgrp", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) addgroup(session string, id int, groupID int) (int, error) {
	v, err := u.client.Call("one.user.addgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) delgroup(session string, id int, groupID int) (int, error) {
	v, err := u.client.Call("one.user.delgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *user) info(session string, id int) (string, error) {
	v, err := u.client.Call("one.user.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
