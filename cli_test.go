package main

import(
	"testing"

	"github.com/go-test/deep"
)

func TestCheckFlag(t *testing.T) {
	var cases = map[string]struct{
		input string
		output bool
	}{
		"Flag is set":
		{
			"/path/to/cna/probes.list",
			false,
		},
		"Flag is not set":
		{
			"",
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := checkFlag(c.input)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
