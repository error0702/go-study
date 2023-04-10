/*
ref: {https://pkg.go.dev/testing}
*/
package hello

import "testing"

func TestPrint(t *testing.T) {
	eqMsg := "hello print"
	msg := Print()
	if msg != eqMsg {
		t.Errorf("print equal fail, source: %s target: %s", msg, eqMsg)
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Print()
	}
}
