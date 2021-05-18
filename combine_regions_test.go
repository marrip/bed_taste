package main

import (
	"testing"

	"github.com/go-test/deep"
)

func checkError(t *testing.T, err error, exp bool) {
	if (err != nil) != exp {
		t.Errorf("Expectation and result are different. Error is\n%v", err)
	}
}

func TestPrepAndCombine(t *testing.T) {
	var cases = map[string]struct {
		input   Region
		probe   Region
		pad     int64
		output  Region
		wantErr bool
	}{
		"Successful combination of probe and region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				110,
				130,
			},
			5,
			Region{
				"1",
				100,
				150,
			},
			false,
		},
		"Chromosomes of region and probe do not match": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"2",
				110,
				130,
			},
			5,
			Region{
				"1",
				100,
				150,
			},
			true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := c.input.prepAndCombine(c.probe, c.pad)
			checkError(t, err, c.wantErr)
			if diff := deep.Equal(c.input, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestChromosomeIdent(t *testing.T) {
	var cases = map[string]struct {
		input  Region
		chr    string
		output bool
	}{
		"Chromosomes are identical": {
			Region{
				"1",
				100,
				150,
			},
			"1",
			true,
		},
		"Chromosomes are not identical": {
			Region{
				"1",
				100,
				150,
			},
			"2",
			false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.input.chromosomeIndent(c.chr)
			if diff := deep.Equal(result, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestAddPadding(t *testing.T) {
	var cases = map[string]struct {
		input   Region
		padding int64
		output  Region
	}{
		"successfully apply padding": {
			Region{
				"1",
				100,
				150,
			},
			50,
			Region{
				"1",
				50,
				200,
			},
		},
		"start falls under 0": {
			Region{
				"1",
				100,
				150,
			},
			150,
			Region{
				"1",
				0,
				300,
			},
		},
		"end excedes Chromosome limit": {
			Region{
				"1",
				248956022,
				248956422,
			},
			150,
			Region{
				"1",
				248955872,
				248956422,
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			c.input.addPadding(c.padding)
			if diff := deep.Equal(c.input, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestCombineRegions(t *testing.T) {
	var cases = map[string]struct {
		input  Region
		probe  Region
		output Region
	}{
		"Probe is completely covered by region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				110,
				130,
			},
			Region{
				"1",
				100,
				150,
			},
		},
		"Probe is completely outside and upstream of region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				50,
				80,
			},
			Region{
				"1",
				50,
				80,
			},
		},
		"Probe is completely outside and downstream of region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				250,
				280,
			},
			Region{
				"1",
				250,
				280,
			},
		},
		"Probe hangs upstream over region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				90,
				110,
			},
			Region{
				"1",
				90,
				150,
			},
		},
		"Probe hangs downstream over region": {
			Region{
				"1",
				100,
				150,
			},
			Region{
				"1",
				140,
				160,
			},
			Region{
				"1",
				100,
				160,
			},
		},
		"Probe hangs on both sides over region": {
			Region{
				"1",
				145,
				150,
			},
			Region{
				"1",
				140,
				160,
			},
			Region{
				"1",
				140,
				160,
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			c.input.combineRegions(c.probe)
			if diff := deep.Equal(c.input, c.output); diff != nil {
				t.Error(diff)
			}
		})
	}
}
