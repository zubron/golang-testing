package operations

import (
	"testing"
)

func TestSomething(t *testing.T) {
	t.Log("Starting the test")
	if true != false {
		t.Fatal("Cannot continue!")
	}
	t.Log("This log will not be printed")
}

func TestShouldBeSkipped(t *testing.T) {
	if true != false {
		t.Errorf("Something has gone very wrong")
		t.Skip("We'll figure this out later")
	}
}
