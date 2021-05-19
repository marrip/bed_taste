package main

type Session struct {
	Hg         string
	Padding    int64
	PathProbes string
	PathGenReg string
}

type RegionInfo struct {
	ID      string
	Gene    string
	GeneID  string
	TransID string
	ExonID  string
	Region  Region
}

type Region struct {
	Chr   string
	Start int64
	End   int64
}
