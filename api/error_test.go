package api_test

import (
	"github.com/SHyx0rmZ/go-one/api"
	"testing"
)

func TestOneError_Error(t *testing.T) {
	err := api.OneError{
		Code:    42,
		Message: "test-message",
	}

	if err.Error() != "test-message" {
		t.Errorf("Error() == %q, want %q", err.Error(), "test-message")
	}
}
