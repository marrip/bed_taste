package main

import(
	"testing"

	"github.com/go-test/deep"
)

func checkError(t *testing.T, err error, exp bool){
	if (err != nil) != exp {
		t.Errorf("Expectation and result are different") 
	}
}

func TestAddPadding(t *testing.T) {
	var cases = map[string]struct{
		input Region
		padding int64
		output Region
	}{
		"successfully apply padding":
		{
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
		"start falls under 0":
		{
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
			"end excedes Chromosome limit":
		{
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
	var cases = map[string]struct{
		exon Region
		probe Region
	}{
		"successfully combine regions":
		{
			Region{
				"1",
				1,
				50,
			},
			Region{
				"1",
				99,
				151,
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			combineRegions(c.exon, c.probe)
			//if diff := deep.Equal(c.input, c.output); diff != nil {
			//	t.Error(diff)
			//}
		})
	}
}
