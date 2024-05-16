package psp

import (
	"strconv"
	"testing"
)

func TestEmailValidity(t *testing.T) {

	tests := map[Email]bool{
		"b@b.com":  true,
		"b+b.com":  false,
		"b@b#.com": false,
		"b@b.c":    false,
	}

	for k, v := range tests {

		if k.Valid() != v {
			t.Error("Email validity failed for", k, "expected", v, "got", k.Valid())
		}
	}
}

func TestColorDepth(t *testing.T) {
	tests := []struct {
		input    ColorDepth
		expected int32
	}{
		{ColorDepth(1), 1},
		{ColorDepth(3), 1},
		{ColorDepth(4), 4},
		{ColorDepth(7), 4},
		{ColorDepth(8), 8},
		{ColorDepth(14), 8},
		{ColorDepth(15), 15},
		{ColorDepth(16), 16},
		{ColorDepth(23), 16},
		{ColorDepth(32), 32},
		{ColorDepth(47), 32},
		{ColorDepth(48), 48},
		{ColorDepth(100), 48},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(int(test.input)), func(t *testing.T) {
			if test.input.GetNormalized() != test.expected {
				t.Errorf("expected %d, got %d", test.expected, test.input.GetNormalized())
			}
		})
	}
}
