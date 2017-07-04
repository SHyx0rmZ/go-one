package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	FilterUsersPrimaryGroup = -4
	FilterUserOnly          = -3
	FilterAll               = -2
	FilterUsersAndGroups    = -1
)

type TemplatePool struct {
	Client xmlrpc.Client
}

func (t *TemplatePool) Info(session string, filter int, start int, end int) (string, error) {
	v, err := t.Client.Call("one.templatepool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
