package bitflags

import (
	"testing"
)

func TestBuildFlags(t *testing.T) {
	var s = struct {
		FLa, FLb, FLc, FLd, FLe, FLf, FLg int8 // works with any other integer types
	}{}
	err := BuildFlagsStruct(&s)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if (s.FLa | s.FLb | s.FLc | s.FLd | s.FLe | s.FLf | s.FLg) != 127 {
		t.Fatal("inconsistent flag enumeration")
	}
}
