package main

import (
	"flag"
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestGetFlags(t *testing.T) {
	var cases = map[string]struct {
		input   []string
		output  Session
		wantErr bool
	}{
		"Missing required flags": {
			[]string{""},
			Session{
				Hg:         "38",
				Padding:    250,
				PathProbes: "",
				PathGenReg: "",
				Output:     "out.bed",
			},
			true,
		},
		"Successfully set values from flags": {
			[]string{"", "-probe", "test_data/exons.tsv", "-genreg", "test_data/exons.tsv"},
			Session{
				Hg:         "38",
				Padding:    250,
				PathProbes: "test_data/exons.tsv",
				PathGenReg: "test_data/exons.tsv",
				Output:     "out.bed",
			},
			false,
		},
		"Human genome version unavailable": {
			[]string{"", "-probe", "test_data/exons.tsv", "-genreg", "test_data/exons.tsv", "-hg", "10"},
			Session{
				Hg:         "10",
				Padding:    250,
				PathProbes: "test_data/exons.tsv",
				PathGenReg: "test_data/exons.tsv",
				Output:     "out.bed",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var s Session
			os.Args = c.input
			flag.CommandLine = flag.NewFlagSet("Reset", flag.ExitOnError)
			err := s.getFlags()
			checkError(t, err, c.wantErr)
			if diff := deep.Equal(s, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestCheckFlag(t *testing.T) {
	var cases = map[string]struct {
		input  string
		output bool
	}{
		"Flag is set": {
			"/path/to/cna/probes.list",
			false,
		},
		"Flag is not set": {
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

func TestCheckHg(t *testing.T) {
	var cases = map[string]struct {
		session Session
		wantErr bool
	}{
		"Human genome version available": {
			Session{
				Hg: "38",
			},
			false,
		},
		"Human genome version unavailable": {
			Session{
				Hg: "10",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := c.session.checkHg()
			checkError(t, err, c.wantErr)
		})
	}
}
