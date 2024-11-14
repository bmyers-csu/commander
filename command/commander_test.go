package command

import (
	"testing"
	"time"
)

func TestGetSystemInfo(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}
}

func TestPingSuccess(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.Ping("google.com")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !info.Successful {
		t.Error("Expected ping to be successful")
	}

	if info.Time <= 0*time.Millisecond {
		t.Error("Expected ping time to be greater than 0")
	}
}

func TestPingFailures(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.Ping("doesnotexist.iewfjreiw")

	if err == nil {
		t.Fatal("Expected an error to occur")
	}

	if info.Successful {
		t.Error("Expected ping to not be successful")
	}
}
