package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
)

func TestCluster_AddDataStore(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.adddatastore")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.AddDataStore("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddDataStore() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_AddHost(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.addhost")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.AddHost("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddHost() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_AddVNet(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.addvnet")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.AddVNet("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("AddVNet() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_Allocate(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.allocate")
	c.ExpectArgumentCount(2)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.String, "test-name")

	i := api.Cluster{
		Client: c,
	}

	v, err := i.Allocate("test-session", "test-name")
	if err != nil || v != 42 {
		t.Errorf("Allocate() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_DelDataStore(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.deldatastore")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.DelDataStore("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelDataStore() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_Delete(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.delete")
	c.ExpectArgumentCount(2)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.Delete("test-session", 42)
	if err != nil || v != 42 {
		t.Errorf("Delete() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_DelHost(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.delhost")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.DelHost("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelHost() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_DelVNet(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.delvnet")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.Int, 21)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.DelVNet("test-session", 42, 21)
	if err != nil || v != 42 {
		t.Errorf("DelVNet() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_Info(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithString("test-info"),
		),
	)

	c.ExpectMethodName("one.cluster.info")
	c.ExpectArgumentCount(2)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.Info("test-session", 42)
	if err != nil || v != "test-info" {
		t.Errorf("Info() == (%v, %q), want (%v, %v)", v, err, "test-info", nil)
	}
}

func TestCluster_Rename(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.rename")
	c.ExpectArgumentCount(3)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.String, "test-name")

	i := api.Cluster{
		Client: c,
	}

	v, err := i.Rename("test-session", 42, "test-name")
	if err != nil || v != 42 {
		t.Errorf("Rename() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}

func TestCluster_Update(t *testing.T) {
	c := mock.NewClient(t).WithValue(
		mock.NewValue().WithValues(
			mock.NewValue().WithBool(true),
			mock.NewValue().WithInt(42),
		),
	)

	c.ExpectMethodName("one.cluster.update")
	c.ExpectArgumentCount(4)
	c.ExpectArgument(0, reflect.String, "test-session")
	c.ExpectArgument(1, reflect.Int, 42)
	c.ExpectArgument(2, reflect.String, "test-template")
	c.ExpectArgument(3, reflect.Int, api.UpdateTypeMerge)

	i := api.Cluster{
		Client: c,
	}

	v, err := i.Update("test-session", 42, "test-template", api.UpdateTypeMerge)
	if err != nil || v != 42 {
		t.Errorf("Update() == (%v, %q), want (%v, %v)", v, err, 42, nil)
	}
}
