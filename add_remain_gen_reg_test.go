package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestAddRemainGenReg(t *testing.T) {
	var cases = map[string]struct {
		input  map[string]RegionInfo
		probes []RegionInfo
		output []RegionInfo
	}{
		"Remaining region is added successfully": {
			map[string]RegionInfo{
				"Gene1_1": {
					ID:     "Gene1_1",
					Gene:   "Gene1",
					ExonID: "1",
				},
				"Gene2_1": {
					ID:     "Gene2_1",
					Gene:   "Gene2",
					ExonID: "1",
				},
			},
			[]RegionInfo{
				{
					ID:     "Gene1_1",
					Gene:   "Gene1",
					ExonID: "1",
				},
			},
			[]RegionInfo{
				{
					ID:     "Gene1_1",
					Gene:   "Gene1",
					ExonID: "1",
				},
				{
					ID:     "Gene2_1",
					Gene:   "Gene2",
					ExonID: "1",
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := addRemainGenReg(c.probes, c.input)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
