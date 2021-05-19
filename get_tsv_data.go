package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

//func (s *Session) getTsvData() {
//
//	return
//}

func readTsv(path string) (data [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	data, err = reader.ReadAll()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func addDataToRegionInfo(data []string) (regioninfo RegionInfo, err error) {
	err = checkRow(data)
	if err != nil {
		return
	}
	region, err := addDataToRegion(data[4:])
	if err != nil {
		return
	}
	regioninfo = RegionInfo{
		ID:      fmt.Sprintf("%s_%s", data[2], data[3]),
		Gene:    data[2],
		GeneID:  data[0],
		TransID: data[1],
		ExonID:  data[3],
		Region:  region,
	}
	return
}

func addDataToRegion(data []string) (region Region, err error) {
	region.Chr = data[0]
	region.Start, err = strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	region.End, err = strconv.ParseInt(data[2], 10, 64)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func checkRow(data []string) (err error) {
	if len(data) != 7 {
		err = errors.New(fmt.Sprintf("row %v does not contain seven fields", data))
		return
	}
	if data[0] != "" && !checkID(data[0]) {
		err = errors.New(fmt.Sprintf("%s does not comply to standard", data[0]))
		return
	}
	if data[1] != "" && !checkID(data[1]) {
		err = errors.New(fmt.Sprintf("%s does not comply to standard", data[1]))
		return
	}
	if checkEmpty(data[2:]) {
		err = errors.New(fmt.Sprintf("required fields of row %v appear to be empty", data))
		return
	}
	err = isInt(data[3:])
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("field in %v is not a convertable number", data))
		return
	}
	return
}

func checkID(id string) bool {
	idregex := regexp.MustCompile("ENS[G,T]\\d+")
	if idregex.MatchString(id) {
		return true
	}
	return false
}

func checkEmpty(data []string) bool {
	for _, field := range data {
		if field == "" {
			return true
		}
	}
	return false
}

func isInt(data []string) (err error) {
	for _, field := range data {
		_, err = strconv.Atoi(field)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}
	return
}
