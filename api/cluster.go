package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type Cluster struct {
	Client xmlrpc.Client
}

func (c *Cluster) Allocate(session string, name string) (int, error) {
	v, err := c.Client.Call("one.cluster.allocate", session, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) Delete(session string, clusterID int) (int, error) {
	v, err := c.Client.Call("one.cluster.delete", session, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) Update(session string, clusterID int, template string, updateType int) (int, error) {
	v, err := c.Client.Call("one.cluster.update", session, clusterID, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) AddHost(session string, clusterID int, hostID int) (int, error) {
	v, err := c.Client.Call("one.cluster.addhost", session, clusterID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) DelHost(session string, clusterID int, hostID int) (int, error) {
	v, err := c.Client.Call("one.cluster.delhost", session, clusterID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) AddDataStore(session string, clusterID int, dataStoreID int) (int, error) {
	v, err := c.Client.Call("one.cluster.adddatastore", session, clusterID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) DelDataStore(session string, clusterID int, dataStoreID int) (int, error) {
	v, err := c.Client.Call("one.cluster.deldatastore", session, clusterID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) AddVNet(session string, clusterID int, virtualNetworkID int) (int, error) {
	v, err := c.Client.Call("one.cluster.addvnet", session, clusterID, virtualNetworkID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) DelVNet(session string, clusterID int, virtualNetworkID int) (int, error) {
	v, err := c.Client.Call("one.cluster.delvnet", session, clusterID, virtualNetworkID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) Rename(session string, clusterID int, name string) (int, error) {
	v, err := c.Client.Call("one.cluster.rename", session, clusterID, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *Cluster) Info(session string, clusterID int) (string, error) {
	v, err := c.Client.Call("one.cluster.info", session, clusterID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
