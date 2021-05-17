package main

import(
	"fmt"
)

func (reg *Region) addPadding(pad int64) {
	start := reg.Start - pad
	end := reg.End + pad
	if start >= 0 {
		reg.Start = start
	} else {
		fmt.Printf("[INFO] adding padding to chr%s:%v-%v results in a negative value - start is set to 0\n", reg.Chr, reg.Start, reg.End)
		reg.Start = 0
	}
	if end <= hg38[reg.Chr] {
		reg.End = end
	} else  {
		fmt.Printf("[INFO] adding padding to chr%s:%v-%v results in value exceding the chromosome limit - end is set to %v\n", reg.Chr, reg.Start, reg.End, hg38[reg.Chr])
		reg.End = hg38[reg.Chr]
	}
}

func combineRegions(exon Region, probe Region) (result Region) {
	result.Chr = exon.Chr
	if exon.Start > probe.End || probe.Start > exon.End {
		result = probe
		fmt.Println(result)
		return
	}
	if exon.Start > probe.Start {
		result.Start = probe.Start
	} else {
		result.Start = exon.Start
	}
	if probe.End > exon.End {
		result.End = probe.End
	} else {
		result.End = exon.End
	}
	fmt.Println(result)
	return
}
