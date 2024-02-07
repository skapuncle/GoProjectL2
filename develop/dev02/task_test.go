package main

import "testing"

func TestUnpack(t *testing.T) {

	if _, error := unpack(`a4bc2d5e`); error != nil {
		t.Error("Wrong result")
	}
	if _, error := unpack(`abcd`); error != nil {
		t.Error("Wrong result")
	}
	if _, error := unpack(``); error != nil {
		t.Error("Wrong result")
	}
}
func TestUnpack2(t *testing.T) {

	if _, error := unpack(`45`); error == nil {
		t.Error("Wrong result")
	}

}
