package bitflags

import (
	"testing"
)

var s = struct {
	FLa, FLb, FLc, FLd, FLe, FLf, FLg int8 // works with any other integer types
}{}

func TestBuildFlags(t *testing.T) {
	err := BuildFlagsStruct(&s) // need pointer
	if err != nil {
		t.Fatalf("%s", err)
	}
	if (s.FLa | s.FLb | s.FLc | s.FLd | s.FLe | s.FLf | s.FLg) != 127 {
		t.Fatal("inconsistent flag enumeration")
	}
}

func TestGetFlagComponents(t *testing.T) {
	BuildFlagsStruct(&s)
	data := GetFlagComponents(s.FLb | s.FLd | s.FLg) // need to be in the ascending order for this test case
	if len(data) != 3 {
		t.Fatal("Inconsistent number of components returned")
	}
	params := []int8{s.FLb, s.FLd, s.FLg}
	for i := range data {
		if params[i] != data[i].(int8) { // the same datatype as the input
			t.Fatalf("Flag %v is not present in the component slice", params[i])
		}
	}
}

func TestFlagInSum(t *testing.T) {
	truth, left := FlagInSum(int8(4), uint8(5))
	if truth != true || left.(uint8) != 1 {
		t.Fatal("1. Unexpected result in FlagInSum")
	}
	truth, left = FlagInSum(s.FLd|s.FLa, s.FLa|s.FLc|s.FLg)
	if truth != false || left != nil {
		t.Fatal("2. Unexpected result in FlagInSum")
	}
}
