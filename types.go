package main

type Session struct {
	Padding    int64
	PathCna    string
	PathProbes string
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
