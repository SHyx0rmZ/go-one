package api

import "github.com/SHyx0rmZ/go-xmlrpc"

type OneError struct {
	Message string
	Code    int
}

func (e OneError) Error() string {
	return e.Message
}

func errorOrBool(v xmlrpc.Value) (bool, error) {
	if v.Values()[0].Bool() != true {
		return false, &OneError{v.Values()[1].String(), v.Values()[2].Int()}
	}

	return v.Values()[1].Bool(), nil
}

func errorOrInt(v xmlrpc.Value) (int, error) {
	if v.Values()[0].Bool() != true {
		return 0, &OneError{v.Values()[1].String(), v.Values()[2].Int()}
	}

	return v.Values()[1].Int(), nil
}

func errorOrString(v xmlrpc.Value) (string, error) {
	if v.Values()[0].Bool() != true {
		return "", &OneError{v.Values()[1].String(), v.Values()[2].Int()}
	}

	return v.Values()[1].String(), nil
}
