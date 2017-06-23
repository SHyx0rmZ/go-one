package one

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	imageTypeOperatingSystem = "OS"
	imageTypeCompactDisc     = "CDROM"
	imageTypeDataBlock       = "DATABLOCK"
	imageTypeKernel          = "KERNEL"
	imageTypeRAMDisk         = "RAMDISK"
	imageTypeContext         = "CONTEXT"
)

type image struct {
	client xmlrpc.Client
}

func (i *image) allocate(session string, template string, dataStoreID int) (int, error) {
	v, err := i.client.Call("one.image.allocate", session, template, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) clone(session string, id int, name string, dataStoreID int) (int, error) {
	v, err := i.client.Call("one.image.clone", session, id, name, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) delete(session string, id int) (int, error) {
	v, err := i.client.Call("one.image.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) enable(session string, id int, enable bool) (int, error) {
	v, err := i.client.Call("one.image.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) persistent(session string, id int, persistent bool) (int, error) {
	v, err := i.client.Call("one.image.persistent", session, id, persistent)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) chtype(session string, id int, imageType string) (int, error) {
	v, err := i.client.Call("one.image.chtype", session, id, imageType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) update(session string, id int, template string, updateType int) (int, error) {
	v, err := i.client.Call("one.image.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := i.client.Call("one.image.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := i.client.Call("one.image.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) rename(session string, id int, name string) (int, error) {
	v, err := i.client.Call("one.image.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) snapshotDelete(session string, id int, snapshotID int) (int, error) {
	v, err := i.client.Call("one.image.snapshotdelete", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) snapshotRevert(session string, id int, snapshotID int) (int, error) {
	v, err := i.client.Call("one.image.snapshotrevert", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) snapshotFlatten(session string, id int, snapshotID int) (int, error) {
	v, err := i.client.Call("one.image.snapshotflatten", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *image) info(session string, id int) (string, error) {
	v, err := i.client.Call("one.image.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
