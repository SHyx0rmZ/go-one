package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type virtualDataCenter struct {
	client xmlrpc.Client
}

func (d *virtualDataCenter) allocate(session string, template string, clusterID int) (int, error) {
	v, err := d.client.Call("one.vdc.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delete(session string, id int) (int, error) {
	v, err := d.client.Call("one.vdc.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.client.Call("one.vdc.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) rename(session string, id int, name string) (int, error) {
	v, err := d.client.Call("one.vdc.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) info(session string, id int) (string, error) {
	v, err := d.client.Call("one.vdc.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (d *virtualDataCenter) addGroup(session string, id int, groupID int) (int, error) {
	v, err := d.client.Call("one.vdc.addgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delGroup(session string, id int, groupID int) (int, error) {
	v, err := d.client.Call("one.vdc.delgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) addCluster(session string, id int, zoneID int, clusterID int) (int, error) {
	v, err := d.client.Call("one.vdc.addcluster", session, id, zoneID, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delCluster(session string, id int, zoneID int, clusterID int) (int, error) {
	v, err := d.client.Call("one.vdc.delcluster", session, id, zoneID, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) addHost(session string, id int, zoneID int, hostID int) (int, error) {
	v, err := d.client.Call("one.vdc.addhost", session, id, zoneID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delHost(session string, id int, zoneID int, hostID int) (int, error) {
	v, err := d.client.Call("one.vdc.delhost", session, id, zoneID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) addDataStore(session string, id int, zoneID int, dataStoreID int) (int, error) {
	v, err := d.client.Call("one.vdc.adddatastore", session, id, zoneID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delDataStore(session string, id int, zoneID int, dataStoreID int) (int, error) {
	v, err := d.client.Call("one.vdc.deldatastore", session, id, zoneID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) addVNet(session string, id int, zoneID int, vnetID int) (int, error) {
	v, err := d.client.Call("one.vdc.addvnet", session, id, zoneID, vnetID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *virtualDataCenter) delVNet(session string, id int, zoneID int, vnetID int) (int, error) {
	v, err := d.client.Call("one.vdc.delvnet", session, id, zoneID, vnetID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}
