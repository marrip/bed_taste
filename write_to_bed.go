package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func (s *Session) write2Bed(data []RegionInfo) (err error) {
	file, err := os.Create(s.Output)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Could not create %s", s.Output))
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	err = writer.WriteAll(prepBedData(data))
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Could not write to file %s", s.Output))
	}
	return
}

func prepBedData(data []RegionInfo) (prepped [][]string) {
	for _, region := range data {
		prepped = append(prepped,
			[]string{
				region.Region.Chr,
				fmt.Sprintf("%v", region.Region.Start),
				fmt.Sprintf("%v", region.Region.End),
				fmt.Sprintf("%s_%s,%s/%s", region.Gene, region.ExonID, region.GeneID, region.TransID),
			})
	}
	return
}
