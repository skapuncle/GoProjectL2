package main

import "testing"

func TestUnpack(t *testing.T) {

	if _, err := unpack(`a4bc2d5e`); err != nil {
		t.Error("Wrong result")
	}
	if _, err := unpack(`abcd`); err != nil {
		t.Error("Wrong result")
	}
	if _, err := unpack(``); err != nil {
		t.Error("Wrong result")
	}
}
func TestUnpack2(t *testing.T) {

	if _, err := unpack(`45`); err == nil {
		t.Error("Wrong result")
	}

}
