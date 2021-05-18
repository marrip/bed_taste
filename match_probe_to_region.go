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
