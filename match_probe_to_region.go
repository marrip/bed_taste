package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func (s *Session) matchProbeToRegion(regions map[string]RegionInfo, probes []RegionInfo) (processed []RegionInfo, err error) {
	for _, probe := range probes {
		region := regions[probe.ID]
		err = region.Region.prepAndCombine(probe.Region, s.Padding)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("probe id is %s", probe.ID))
			return
		}
		processed = append(processed, region)
	}
	return
}

func (reg *Region) prepAndCombine(probe Region, pad int64) (err error) {
	if !reg.chromosomeIndent(probe.Chr) {
		err = errors.New(fmt.Sprintf("%s:%v-%v and %s:%v-%v are located on different chromosomes", reg.Chr, reg.Start, reg.End, probe.Chr, probe.Start, probe.End))
		return
	}
	probe.addPadding(pad)
	reg.combineRegions(probe)
	return
}

func (reg *Region) chromosomeIndent(chr string) bool {
	if reg.Chr == chr {
		return true
	}
	return false
}

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
	} else {
		fmt.Printf("[INFO] adding padding to chr%s:%v-%v results in value exceding the chromosome limit - end is set to %v\n", reg.Chr, reg.Start, reg.End, hg38[reg.Chr])
		reg.End = hg38[reg.Chr]
	}
}

func (reg *Region) combineRegions(probe Region) {
	if reg.Start > probe.End || probe.Start > reg.End {
		reg.Start = probe.Start
		reg.End = probe.End
		return
	}
	if reg.Start > probe.Start {
		reg.Start = probe.Start
	}
	if probe.End > reg.End {
		reg.End = probe.End
	}
	return
}
