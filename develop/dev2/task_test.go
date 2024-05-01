package main

import (
	"testing"
)

func TestUnpack_1(t *testing.T) {
	s := "a4bc2d5e"
	expected := "aaaabccddddde"

	res, err := Unpack(s)

	if res != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, res)
	}

	if err != nil {
		t.Errorf("Incorrect result. Error must be nil, got %s", err)
	}

}

func TestUnpack_2(t *testing.T) {
	s := "abcd"
	expected := "abcd"

	res, err := Unpack(s)

	if res != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, res)
	}

	if err != nil {
		t.Errorf("Incorrect result. Error must be nil, got %s", err)
	}
}

func TestUnpack_3(t *testing.T) {
	s := "45"
	expected := ""

	res, err := Unpack(s)

	if res != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, res)
	}

	if err == nil {
		t.Errorf("Incorrect result. Error must be nil, got")
	}
}

func TestUnpack_4(t *testing.T) {
	s := ""
	expected := ""

	res, err := Unpack(s)

	if res != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, res)
	}

	if err != nil {
		t.Errorf("Incorrect result. Error must be nil, got %s", err)
	}
}
