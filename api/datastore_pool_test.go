package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestDatastorePool_Info(t *testing.T) {
	c := client(t, "one.datastorepool.info", []arg{
		{reflect.String, "test-session"},
	}, []xmlrpc.Value{
		mock.NewValue().WithBool(true),
		mock.NewValue().WithString("test-info"),
	})

	i := api.DatastorePool{Client: c}
	v, err := i.Info("test-session")
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%q, %v), want (%q, %v)", v, err, "test-info", nil)
	}
}
