package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type SecurityGroup struct {
	Client xmlrpc.Client
}

func (s *SecurityGroup) Allocate(session string, template string) (int, error) {
	v, err := s.Client.Call("one.secgroup.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Clone(session string, id int, name string, dataStoreID int) (int, error) {
	v, err := s.Client.Call("one.secgroup.clone", session, id, name, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Delete(session string, id int) (int, error) {
	v, err := s.Client.Call("one.secgroup.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Update(session string, id int, template string, updateType int) (int, error) {
	v, err := s.Client.Call("one.secgroup.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Commit(session string, id int, outdatedAndErroredOnly bool) (int, error) {
	v, err := s.Client.Call("one.secgroup.commit", session, id, outdatedAndErroredOnly)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := s.Client.Call("one.secgroup.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := s.Client.Call("one.secgroup.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Rename(session string, id int, name string) (int, error) {
	v, err := s.Client.Call("one.secgroup.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *SecurityGroup) Info(session string, id int) (string, error) {
	v, err := s.Client.Call("one.secgroup.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
