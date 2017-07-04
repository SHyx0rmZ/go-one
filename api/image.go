package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	ImageTypeOperatingSystem = "OS"
	ImageTypeCompactDisc     = "CDROM"
	ImageTypeDataBlock       = "DATABLOCK"
	ImageTypeKernel          = "KERNEL"
	ImageTypeRAMDisk         = "RAMDISK"
	ImageTypeContext         = "CONTEXT"
)

type Image struct {
	Client xmlrpc.Client
}

func (i *Image) Allocate(session string, template string, dataStoreID int) (int, error) {
	v, err := i.Client.Call("one.image.allocate", session, template, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Clone(session string, id int, name string, dataStoreID int) (int, error) {
	v, err := i.Client.Call("one.image.clone", session, id, name, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Delete(session string, id int) (int, error) {
	v, err := i.Client.Call("one.image.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Enable(session string, id int, enable bool) (int, error) {
	v, err := i.Client.Call("one.image.enable", session, id, enable)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Persistent(session string, id int, persistent bool) (int, error) {
	v, err := i.Client.Call("one.image.persistent", session, id, persistent)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Chtype(session string, id int, imageType string) (int, error) {
	v, err := i.Client.Call("one.image.chtype", session, id, imageType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := i.Client.Call("one.image.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := i.Client.Call("one.image.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := i.Client.Call("one.image.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Rename(session string, id int, name string) (int, error) {
	v, err := i.Client.Call("one.image.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) SnapshotDelete(session string, id int, snapshotID int) (int, error) {
	v, err := i.Client.Call("one.image.snapshotdelete", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) SnapshotRevert(session string, id int, snapshotID int) (int, error) {
	v, err := i.Client.Call("one.image.snapshotrevert", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) SnapshotFlatten(session string, id int, snapshotID int) (int, error) {
	v, err := i.Client.Call("one.image.snapshotflatten", session, id, snapshotID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (i *Image) Info(session string, id int) (string, error) {
	v, err := i.Client.Call("one.image.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
