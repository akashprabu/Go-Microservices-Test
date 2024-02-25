package test

import "testing"

var TestData = []struct {
	Name     string
	xValue   float32
	yValue   float32
	isError  bool
	Expected float32
}{
	{"positive-scenario", 100.0, 10.0, false, 10.0},
	{"negative-scenario", 100.0, 0, true, 0},
}

func TestDividePositive(t *testing.T) {
	_, err := divide(100.0, 10.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestDivideNegative(t *testing.T) {
	_, err := divide(100.0, 0)
	if err == nil {
		t.Error("Excepted an error but we did not get")
	}
}

func TestDivision(t *testing.T) {
	for _, tt := range TestData {
		value, err := divide(tt.xValue, tt.yValue)
		if tt.isError {
			if err == nil {
				t.Error("Expected a error, but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect a error but got one", err)
			}
		}
		if value != tt.Expected {
			t.Errorf("Expected %f but got %f", tt.Expected, value)
		}
	}
}
