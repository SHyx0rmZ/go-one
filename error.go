package one

import "github.com/SHyx0rmZ/go-xmlrpc"

type OneError struct {
	message string
	code    int
}

func (e OneError) Error() string {
	return e.message
}

func errorOrInt(v xmlrpc.Value) (int, error) {
	if v.AsArray()[0].AsBool() != true {
		return 0, &OneError{v.AsArray()[1].AsString(), v.AsArray()[2].AsInt()}
	}

	return v.AsArray()[1].AsInt(), nil
}

func errorOrString(v xmlrpc.Value) (string, error) {
	if v.AsArray()[0].AsBool() != true {
		return "", &OneError{v.AsArray()[1].AsString(), v.AsArray()[2].AsInt()}
	}

	return v.AsArray()[1].AsString(), nil
}
