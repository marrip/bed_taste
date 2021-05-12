package main

import(
	"flag"
	"fmt"
)

func get_flags() {
	// Read flags
	var cna string
	flag.StringVar(&cna, "cna", "", "path to file containing MLPA probes")

	// Parse flags
	flag.Parse()

	// Check if required flags exist
	if check_flag(cna) {
		flag.PrintDefaults()
		fmt.Println("missing required flags")
	}
	return
}

func check_flag(flag string) bool {
	if flag == "" {
		return true
	}
	return false
}
