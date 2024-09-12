package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ToUpperFirst(t *testing.T) {
	tests := map[string]struct {
		input  string
		expect string
	}{
		"ok: common case": {input: "hello", expect: "Hello"},
		"ok: empty":       {input: "", expect: ""},
		"ng: non str":     {input: "0123", expect: "\x10123"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ToUpperFirst(tt.input)
			require.Equal(t, tt.expect, got)
		})
	}

}
