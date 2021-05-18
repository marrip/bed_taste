package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestReadTsv(t *testing.T) {
	var cases = map[string]struct {
		input   string
		output  [][]string
		wantErr bool
	}{
		"Reading existing tsv file": {
			"test_data/exons.tsv",
			[][]string{
				{
					"ENSG00000147883",
					"ENST00000380142",
					"CDKN2B",
					"1",
					"9",
					"22008675",
					"22009272",
				},
				{
					"ENSG00000147883",
					"ENST00000380142",
					"CDKN2B",
					"2",
					"9",
					"22004748",
					"22006247",
				},
			},
			false,
		},
		"Failing to read non-existent tsv file": {
			"test_data/eksons.tsv",
			nil,
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := readTsv(c.input)
			checkError(t, err, c.wantErr)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestCheckRow(t *testing.T) {
	var cases = map[string]struct {
		input   []string
		wantErr bool
	}{
		"Row complies to standard": {
			[]string{
				"ENSG00000147883",
				"ENST00000380142",
				"CDKN2B",
				"1",
				"9",
				"22008675",
				"22009272",
			},
			false,
		},
		"Row too short": {
			[]string{
				"ENSG00000147883",
				"ENST00000380142",
				"CDKN2B",
				"9",
				"22008675",
				"22009272",
			},
			true,
		},
		"Row contains invalid id": {
			[]string{
				"ENSG00000147883",
				"ENSV00000380142",
				"CDKN2B",
				"1",
				"9",
				"22008675",
				"22009272",
			},
			true,
		},
		"Row contains empty fields at required positions": {
			[]string{
				"ENSG00000147883",
				"ENSV00000380142",
				"CDKN2B",
				"1",
				"9",
				"",
				"",
			},
			true,
		},
		"Row does not contain numbers at required positions": {
			[]string{
				"ENSG00000147883",
				"ENSV00000380142",
				"CDKN2B",
				"one",
				"9",
				"22008675",
				"22009272",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := checkRow(c.input)
			checkError(t, err, c.wantErr)
		})
	}
}

func TestCheckID(t *testing.T) {
	var cases = map[string]struct {
		input  string
		output bool
	}{
		"Is a valid id": {
			"ENSG00000147883",
			true,
		},
		"Is not a valid id": {
			"MYID",
			false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := checkID(c.input)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestCheckEmpty(t *testing.T) {
	var cases = map[string]struct {
		input  []string
		output bool
	}{
		"No empty field": {
			[]string{
				"this",
				"is",
				"valid",
				"data",
			},
			false,
		},
		"Empty fields": {
			[]string{
				"",
				"data",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := checkEmpty(c.input)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	var cases = map[string]struct {
		input   []string
		wantErr bool
	}{
		"Convertable integers": {
			[]string{
				"1",
				"2",
			},
			false,
		},
		"Non-convertable strings": {
			[]string{
				"one",
				"two",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := isInt(c.input)
			checkError(t, err, c.wantErr)
		})
	}
}
