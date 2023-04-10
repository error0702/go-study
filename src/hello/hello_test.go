package hello

import "testing"

func TestPrint(t *testing.T) {
	eqMsg := "hello print"
	msg := Print()
	if msg != eqMsg {
		t.Errorf("print equal fail, source: %s target: %s", msg, eqMsg)
	}
}
