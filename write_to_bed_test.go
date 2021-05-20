package main

import (
	"log"
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestWrite2Bed(t *testing.T) {
	var cases = map[string]struct {
		input   []RegionInfo
		session Session
		wantErr bool
	}{
		"Success": {
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG01",
					TransID: "ENST01",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   200,
					},
				},
			},
			Session{
				Output: "out.bed",
			},
			false,
		},
		"Wrong path": {
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG01",
					TransID: "ENST01",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   200,
					},
				},
			},
			Session{
				Output: "non-existent/out.bed",
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := c.session.write2Bed(c.input)
			checkError(t, err, c.wantErr)
			if err == nil {
				err = os.Remove("out.bed")
				if err != nil {
					log.Fatalf("%v", err)
				}
			}
		})
	}
}

func TestPrepBedData(t *testing.T) {
	var cases = map[string]struct {
		input  []RegionInfo
		output [][]string
	}{
		"Success": {
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG01",
					TransID: "ENST01",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   200,
					},
				},
			},
			[][]string{
				{
					"1",
					"100",
					"200",
					"Gene1_1,ENSG01/ENST01",
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := prepBedData(c.input)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
