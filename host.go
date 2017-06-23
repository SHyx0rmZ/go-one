package one

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	hostStatusEnabled  = 0
	hostStatusDisabled = 1
	hostStatusOffline  = 2
)

type host struct {
	client xmlrpc.Client
}

func (h *host) allocate(session string, hostname string, informationManagerName string, virtualMachineManagerName string, clusterID int) (int, error) {
	v, err := h.client.Call("one.host.allocate", session, hostname, informationManagerName, virtualMachineManagerName, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *host) delete(session string, hostID int) (int, error) {
	v, err := h.client.Call("one.host.delete", session, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *host) status(session string, hostID int, hostStatus int) (int, error) {
	v, err := h.client.Call("one.host.status", session, hostID, hostStatus)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *host) update(session string, hostID int, template string, updateType int) (int, error) {
	v, err := h.client.Call("one.host.update", session, hostID, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *host) rename(session string, hostID int, name string) (int, error) {
	v, err := h.client.Call("one.host.rename", session, hostID, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *host) info(session string, hostID int) (string, error) {
	v, err := h.client.Call("one.host.info", session, hostID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (h *host) monitoring(session string, hostID int) (string, error) {
	v, err := h.client.Call("one.host.monitoring", session, hostID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
