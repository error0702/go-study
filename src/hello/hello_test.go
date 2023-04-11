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

func TestSplit(t *testing.T) {
	val := 10
	x, y := Split(val)
	if x != 4 && y != 6 {
		t.Errorf("test fail, value: %d, x: %d, y: %d", val, x, y)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split()
	}
}
