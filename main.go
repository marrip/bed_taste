package main

import (
	"fmt"
	"log"
)

func main() {
	var s Session
	err := s.getFlags()
	if err != nil {
		log.Fatalf("%v", err)
	}
	data, err := s.data2CombinedRegions()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(data)
}

func (s Session) data2CombinedRegions() (completed []RegionInfo, err error) {
	probes, err := getTsvData(s.PathProbes)
	if err != nil {
		return
	}
	regions, err := getGenRegTsvData(s.PathGenReg)
	if err != nil {
		return
	}
	processed, err := s.matchProbeToRegion(regions, probes)
	if err != nil {
		return
	}
	completed = addRemainGenReg(processed, regions)
	return
}
