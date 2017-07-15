package api

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"errors"
)

var ErrInvalidOneResponse = errors.New("Got an invalid response from the XML-RPC API")

type OneError struct {
	Message string
	Code    int
}

func (e OneError) Error() string {
	return e.Message
}

func errorOrBool(v xmlrpc.Value) (bool, error) {
	if v.Kind() != xmlrpc.Array || len(v.Values()) < 1 || v.Values()[0].Kind() != xmlrpc.Bool {
		return false, ErrInvalidOneResponse
	}

	vs := v.Values()

	if vs[0].Bool() != true {
		if len(vs) < 3 || vs[1].Kind() != xmlrpc.String || vs[2].Kind() != xmlrpc.Int {
			return false, ErrInvalidOneResponse
		}

		return false, &OneError{vs[1].String(), vs[2].Int()}
	}

	if len(vs) < 2 || vs[1].Kind() != xmlrpc.Bool {
		return false, ErrInvalidOneResponse
	}

	return vs[1].Bool(), nil
}

func errorOrInt(v xmlrpc.Value) (int, error) {
	if v.Kind() != xmlrpc.Array || len(v.Values()) < 1 || v.Values()[0].Kind() != xmlrpc.Bool {
		return 0, ErrInvalidOneResponse
	}

	vs := v.Values()

	if vs[0].Bool() != true {
		if len(vs) < 3 || vs[1].Kind() != xmlrpc.String || vs[2].Kind() != xmlrpc.Int {
			return 0, ErrInvalidOneResponse
		}

		return 0, &OneError{vs[1].String(), vs[2].Int()}
	}

	if len(vs) < 2 || vs[1].Kind() != xmlrpc.Int {
		return 0, ErrInvalidOneResponse
	}

	return vs[1].Int(), nil
}

func errorOrString(v xmlrpc.Value) (string, error) {
	if v.Kind() != xmlrpc.Array || len(v.Values()) < 1 || v.Values()[0].Kind() != xmlrpc.Bool {
		return "", ErrInvalidOneResponse
	}
	
	vs := v.Values()
	
	if vs[0].Bool() != true {
		if len(vs) < 3 || vs[1].Kind() != xmlrpc.String || vs[2].Kind() != xmlrpc.Int {
			return "", ErrInvalidOneResponse
		}

		return "", &OneError{vs[1].String(), vs[2].Int()}
	}

	if len(vs) < 2 || vs[1].Kind() != xmlrpc.String {
		return "", ErrInvalidOneResponse
	}

	return vs[1].String(), nil
}
