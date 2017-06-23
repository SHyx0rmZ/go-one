package one

import	"github.com/SHyx0rmZ/go-xmlrpc"

const (
	filterUsersPrimaryGroup = -4
	filterUserOnly          = -3
	filterAll               = -2
	filterUsersAndGroups    = -1
)

type templatePool struct {
	client xmlrpc.Client
}

func (t *templatePool) info(session string, filter int, start int, end int) (string, error) {
	v, err := t.client.Call("one.templatepool.info", session, filter, start, end)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
