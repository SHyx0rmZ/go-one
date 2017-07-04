package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type securityGroup struct {
	client xmlrpc.Client
}

func (s *securityGroup) allocate(session string, template string) (int, error) {
	v, err := s.client.Call("one.secgroup.allocate", session, template)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) clone(session string, id int, name string, dataStoreID int) (int, error) {
	v, err := s.client.Call("one.secgroup.clone", session, id, name, dataStoreID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) delete(session string, id int) (int, error) {
	v, err := s.client.Call("one.secgroup.delete", session, id)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) update(session string, id int, template string, updateType int) (int, error) {
	v, err := s.client.Call("one.secgroup.update", session, id, template, updateType)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) commit(session string, id int, outdatedAndErroredOnly bool) (int, error) {
	v, err := s.client.Call("one.secgroup.commit", session, id, outdatedAndErroredOnly)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) chmod(session string, id int, userUse int, userManage int, userAdmin int, groupUse int, groupManage int, groupAdmin int, otherUse int, otherManage int, otherAdmin int) (int, error) {
	v, err := s.client.Call("one.secgroup.chmod", session, id, userUse, userManage, userAdmin, groupUse, groupManage, groupAdmin, otherUse, otherManage, otherAdmin)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) chown(session string, id int, userID int, groupID int) (int, error) {
	v, err := s.client.Call("one.secgroup.chown", session, id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) rename(session string, id int, name string) (int, error) {
	v, err := s.client.Call("one.secgroup.rename", session, id, name)
	if err != nil {
		return 0, err
	}

	return errorOrInt(v)
}

func (s *securityGroup) info(session string, id int) (string, error) {
	v, err := s.client.Call("one.secgroup.info", session, id)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}
