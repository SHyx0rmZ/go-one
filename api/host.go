package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	HostStatusEnabled  = 0
	HostStatusDisabled = 1
	HostStatusOffline  = 2
)

type Host struct {
	Client xmlrpc.Client
}

func (h *Host) Allocate(session string, hostname string, informationManagerName string, virtualMachineManagerName string, clusterID int) (int, error) {
	v, err := h.Client.Call("one.host.allocate", session, hostname, informationManagerName, virtualMachineManagerName, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *Host) Delete(session string, hostID int) (int, error) {
	v, err := h.Client.Call("one.host.delete", session, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *Host) Status(session string, hostID int, hostStatus int) (int, error) {
	v, err := h.Client.Call("one.host.status", session, hostID, hostStatus)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *Host) Update(session string, hostID int, template string, updateType int) (int, error) {
	v, err := h.Client.Call("one.host.update", session, hostID, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *Host) Rename(session string, hostID int, name string) (int, error) {
	v, err := h.Client.Call("one.host.rename", session, hostID, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (h *Host) Info(session string, hostID int) (string, error) {
	v, err := h.Client.Call("one.host.info", session, hostID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (h *Host) Monitoring(session string, hostID int) (string, error) {
	v, err := h.Client.Call("one.host.monitoring", session, hostID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
