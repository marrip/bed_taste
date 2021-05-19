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

func TestAddDataToRegionInfo(t *testing.T) {
	var cases = map[string]struct {
		input   []string
		output  RegionInfo
		wantErr bool
	}{
		"Successfully add data to region info": {
			[]string{
				"ENSG00000147883",
				"ENST00000380142",
				"CDKN2B",
				"2",
				"9",
				"22004748",
				"22006247",
			},
			RegionInfo{
				ID:      "CDKN2B_2",
				Gene:    "CDKN2B",
				GeneID:  "ENSG00000147883",
				TransID: "ENST00000380142",
				ExonID:  "2",
				Region: Region{
					Chr:   "9",
					Start: 22004748,
					End:   22006247,
				},
			},
			false,
		},
		"Row is too short": {
			[]string{
				"ENSG00000147883",
				"CDKN2B",
				"2",
				"9",
				"22004748",
				"22006247",
			},
			RegionInfo{
				ID:      "",
				Gene:    "",
				GeneID:  "",
				TransID: "",
				ExonID:  "",
				Region: Region{
					Chr:   "",
					Start: 0,
					End:   0,
				},
			},
			false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := addDataToRegionInfo(c.input)
			checkError(t, err, c.wantErr)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestAddDataToRegion(t *testing.T) {
	var cases = map[string]struct {
		input   []string
		output  Region
		wantErr bool
	}{
		"Successfully add data to region": {
			[]string{
				"9",
				"22008675",
				"22009272",
			},
			Region{
				Chr:   "9",
				Start: 22008675,
				End:   22009272,
			},
			false,
		},
		"Failing at adding data to region": {
			[]string{
				"9",
				"one",
				"two",
			},
			Region{
				Chr:   "9",
				Start: 0,
				End:   0,
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := addDataToRegion(c.input)
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
				"ENST00000380142",
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
				"ENST00000380142",
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
