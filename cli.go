package main

import(
	"flag"
	"fmt"
)

func getFlags() {
	// Read flags
	var cna, hg string
	var padding int64
	flag.StringVar(&cna, "cna", "", "path to file containing MLPA probes")
	flag.StringVar(&hg, "hg", "38", "version of human genome")
	flag.Int64Var(&padding, "padding", 250, "padding that should be applied to MLPA probe regions")

	// Parse flags
	flag.Parse()

	// Check if required flags exist
	if checkFlag(cna) {
		flag.PrintDefaults()
		fmt.Println("missing required flags")
	}
	return
}

func checkFlag(flag string) bool {
	if flag == "" {
		return true
	}
	return false
}
