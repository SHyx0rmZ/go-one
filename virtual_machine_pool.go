package one

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	stateAny            = -2
	stateAnyExceptDone  = -1
	stateInit           = 0
	statePending        = 1
	stateHold           = 2
	stateActive         = 3
	stateStopped        = 4
	stateSuspended      = 5
	stateDone           = 6
	stateFailed         = 7
	statePowerOff       = 8
	stateUndeployed     = 9
	stateCloning        = 10
	stateCloningFailure = 11
)

type virtualMachinePool struct {
	client xmlrpc.Client
}

func (p *virtualMachinePool) info(session string, filter int, start int, end int, state int) (string, error) {
	v, err := p.client.Call("one.vmpool.info", session, filter, start, end, state)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *virtualMachinePool) monitoring(session string, filter int) (string, error) {
	v, err := p.client.Call("one.vmpool.monitoring", session, filter)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *virtualMachinePool) accounting(session string, filter int, startTime int, endTime int) (string, error) {
	v, err := p.client.Call("one.vmpool.accounting", session, filter, startTime, endTime)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *virtualMachinePool) showback(session string, filter int, firstMonth int, firstYear int, lastMonth int, lastYear int) (string, error) {
	v, err := p.client.Call("one.vmpool.showback", session, filter, firstMonth, firstYear, lastMonth, lastYear)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *virtualMachinePool) calculateShowback(session string, firstMonth int, firstYear int, lastMonth int, lastYear int) error {
	v, err := p.client.Call("one.vmpool.calculateshowback", session, firstMonth, firstYear, lastMonth, lastYear)
	if err != nil {
		return err
	}

	if v.AsArray()[0].AsBool() != true {
		return &OneError{v.AsArray()[1].AsString(), v.AsArray()[2].AsInt()}
	}

	return nil
}
