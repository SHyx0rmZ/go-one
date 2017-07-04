package api

import "github.com/SHyx0rmZ/go-xmlrpc"

const (
	StateAny            = -2
	StateAnyExceptDone  = -1
	StateInit           = 0
	StatePending        = 1
	StateHold           = 2
	StateActive         = 3
	StateStopped        = 4
	StateSuspended      = 5
	StateDone           = 6
	StateFailed         = 7
	StatePowerOff       = 8
	StateUndeployed     = 9
	StateCloning        = 10
	StateCloningFailure = 11
)

type VirtualMachinePool struct {
	Client xmlrpc.Client
}

func (p *VirtualMachinePool) Info(session string, filter int, start int, end int, state int) (string, error) {
	v, err := p.Client.Call("one.vmpool.info", session, filter, start, end, state)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *VirtualMachinePool) Monitoring(session string, filter int) (string, error) {
	v, err := p.Client.Call("one.vmpool.monitoring", session, filter)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *VirtualMachinePool) Accounting(session string, filter int, startTime int, endTime int) (string, error) {
	v, err := p.Client.Call("one.vmpool.accounting", session, filter, startTime, endTime)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *VirtualMachinePool) Showback(session string, filter int, firstMonth int, firstYear int, lastMonth int, lastYear int) (string, error) {
	v, err := p.Client.Call("one.vmpool.showback", session, filter, firstMonth, firstYear, lastMonth, lastYear)
	if err != nil {
		return "", err
	}

	return errorOrString(v)
}

func (p *VirtualMachinePool) CalculateShowback(session string, firstMonth int, firstYear int, lastMonth int, lastYear int) error {
	v, err := p.Client.Call("one.vmpool.calculateshowback", session, firstMonth, firstYear, lastMonth, lastYear)
	if err != nil {
		return err
	}

	if v.AsArray()[0].AsBool() != true {
		return &OneError{v.AsArray()[1].AsString(), v.AsArray()[2].AsInt()}
	}

	return nil
}
