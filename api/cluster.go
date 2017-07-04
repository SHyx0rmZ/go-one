package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type cluster struct {
	client xmlrpc.Client
}

func (c *cluster) allocate(session string, name string) (int, error) {
	v, err := c.client.Call("one.cluster.allocate", session, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) delete(session string, clusterID int) (int, error) {
	v, err := c.client.Call("one.cluster.delete", session, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) update(session string, clusterID int, template string, updateType int) (int, error) {
	v, err := c.client.Call("one.cluster.update", session, clusterID, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) addHost(session string, clusterID int, hostID int) (int, error) {
	v, err := c.client.Call("one.cluster.addhost", session, clusterID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) delHost(session string, clusterID int, hostID int) (int, error) {
	v, err := c.client.Call("one.cluster.delhost", session, clusterID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) addDataStore(session string, clusterID int, dataStoreID int) (int, error) {
	v, err := c.client.Call("one.cluster.adddatastore", session, clusterID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) delDataStore(session string, clusterID int, dataStoreID int) (int, error) {
	v, err := c.client.Call("one.cluster.deldatastore", session, clusterID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) addVNet(session string, clusterID int, virtualNetworkID int) (int, error) {
	v, err := c.client.Call("one.cluster.addvnet", session, clusterID, virtualNetworkID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) delVNet(session string, clusterID int, virtualNetworkID int) (int, error) {
	v, err := c.client.Call("one.cluster.delvnet", session, clusterID, virtualNetworkID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) rename(session string, clusterID int, name string) (int, error) {
	v, err := c.client.Call("one.cluster.rename", session, clusterID, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (c *cluster) info(session string, clusterID int) (string, error) {
	v, err := c.client.Call("one.cluster.info", session, clusterID)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
