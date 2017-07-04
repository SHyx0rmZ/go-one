package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	AuthenticationDriverCore = "core"
)

type User struct {
	Client xmlrpc.Client
}

func (u *User) Allocate(session string, username string, password string, authenticationDriver string, groupIDs ...int) (int, error) {
	v, err := u.Client.Call("one.user.allocate", session, username, password, authenticationDriver, groupIDs)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Delete(session string, id int) (int, error) {
	v, err := u.Client.Call("one.user.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Password(session string, id int, password string) (int, error) {
	v, err := u.Client.Call("one.user.password", session, id, password)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Login(session string, username string, token string, ttlInSeconds int, effectiveGroupID int) (string, error) {
	v, err := u.Client.Call("one.user.login", session, username, token, ttlInSeconds, effectiveGroupID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (u *User) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := u.Client.Call("one.user.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Chauth(session string, id int, authenticationDriver string, optionalNewPassword string) (int, error) {
	v, err := u.Client.Call("one.user.chauth", session, id, authenticationDriver, optionalNewPassword)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Quota(session string, id int, template string) (int, error) {
	v, err := u.Client.Call("one.user.quota", session, id, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Chgrp(session string, id int, groupID int) (int, error) {
	v, err := u.Client.Call("one.user.chgrp", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) AddGroup(session string, id int, groupID int) (int, error) {
	v, err := u.Client.Call("one.user.addgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) DelGroup(session string, id int, groupID int) (int, error) {
	v, err := u.Client.Call("one.user.delgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (u *User) Info(session string, id int) (string, error) {
	v, err := u.Client.Call("one.user.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
