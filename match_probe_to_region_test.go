package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMatchProbeToRegion(t *testing.T) {
	var cases = map[string]struct {
		session Session
		input   map[string]RegionInfo
		probes  []RegionInfo
		output  []RegionInfo
		wantErr bool
	}{
		"Successfully match probe to region": {
			Session{
				Padding: 50,
			},
			map[string]RegionInfo{
				"Gene1_1": {
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG001",
					TransID: "ENST001",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   300,
					},
				},
			},
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG001",
					TransID: "ENST001",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 200,
						End:   220,
					},
				},
			},
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG001",
					TransID: "ENST001",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   300,
					},
				},
			},
			false,
		},
		"Failing to match probe to region": {
			Session{
				Padding: 50,
			},
			map[string]RegionInfo{
				"Gene1_1": {
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG001",
					TransID: "ENST001",
					ExonID:  "1",
					Region: Region{
						Chr:   "1",
						Start: 100,
						End:   300,
					},
				},
			},
			[]RegionInfo{
				{
					ID:      "Gene1_1",
					Gene:    "Gene1",
					GeneID:  "ENSG001",
					TransID: "ENST001",
					ExonID:  "1",
					Region: Region{
						Chr:   "2",
						Start: 200,
						End:   220,
					},
				},
			},
			nil,
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.session.matchProbeToRegion(c.input, c.probes)
			checkError(t, err, c.wantErr)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
