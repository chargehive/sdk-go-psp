package psp

import (
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
