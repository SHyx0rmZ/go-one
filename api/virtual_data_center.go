package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type VirtualDataCenter struct {
	Client xmlrpc.Client
}

func (d *VirtualDataCenter) Allocate(session string, template string, clusterID int) (int, error) {
	v, err := d.Client.Call("one.vdc.allocate", session, template, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) Delete(session string, id int) (int, error) {
	v, err := d.Client.Call("one.vdc.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := d.Client.Call("one.vdc.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) Rename(session string, id int, name string) (int, error) {
	v, err := d.Client.Call("one.vdc.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) Info(session string, id int) (string, error) {
	v, err := d.Client.Call("one.vdc.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (d *VirtualDataCenter) AddGroup(session string, id int, groupID int) (int, error) {
	v, err := d.Client.Call("one.vdc.addgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) DelGroup(session string, id int, groupID int) (int, error) {
	v, err := d.Client.Call("one.vdc.delgroup", session, id, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) AddCluster(session string, id int, zoneID int, clusterID int) (int, error) {
	v, err := d.Client.Call("one.vdc.addcluster", session, id, zoneID, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) DelCluster(session string, id int, zoneID int, clusterID int) (int, error) {
	v, err := d.Client.Call("one.vdc.delcluster", session, id, zoneID, clusterID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) AddHost(session string, id int, zoneID int, hostID int) (int, error) {
	v, err := d.Client.Call("one.vdc.addhost", session, id, zoneID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) DelHost(session string, id int, zoneID int, hostID int) (int, error) {
	v, err := d.Client.Call("one.vdc.delhost", session, id, zoneID, hostID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) AddDataStore(session string, id int, zoneID int, dataStoreID int) (int, error) {
	v, err := d.Client.Call("one.vdc.adddatastore", session, id, zoneID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) DelDataStore(session string, id int, zoneID int, dataStoreID int) (int, error) {
	v, err := d.Client.Call("one.vdc.deldatastore", session, id, zoneID, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) AddVNet(session string, id int, zoneID int, vnetID int) (int, error) {
	v, err := d.Client.Call("one.vdc.addvnet", session, id, zoneID, vnetID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (d *VirtualDataCenter) DelVNet(session string, id int, zoneID int, vnetID int) (int, error) {
	v, err := d.Client.Call("one.vdc.delvnet", session, id, zoneID, vnetID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}
