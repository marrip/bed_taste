package main

func addRemainGenReg(regioninfos []RegionInfo, genreg map[string]RegionInfo) []RegionInfo {
	for _, regioninfo := range regioninfos {
		delete(genreg, regioninfo.ID)
	}
	for _, regioninfo := range genreg {
		regioninfos = append(regioninfos, regioninfo)
	}
	return regioninfos
}
