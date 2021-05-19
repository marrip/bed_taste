package main

import (
	"flag"
	"fmt"

	"github.com/pkg/errors"
)

func (s *Session) getFlags() error {
	flag.StringVar(&s.PathProbes, "probe", "", "path to file containing MLPA probes (required)")
	flag.StringVar(&s.PathGenReg, "genreg", "", "path to file containing genetic regions (required)")
	flag.StringVar(&s.Hg, "hg", "38", "version of human genome")
	flag.Int64Var(&s.Padding, "padding", 250, "padding that should be applied to MLPA probe regions")
	flag.Parse()

	if checkFlag(s.PathProbes) || checkFlag(s.PathGenReg) {
		flag.PrintDefaults()
		return errors.New("missing required flags")
	}
	if err := s.checkHg(); err != nil {
		return err
	}
	return nil
}

func checkFlag(flag string) bool {
	if flag == "" {
		return true
	}
	return false
}

func (s *Session) checkHg() (err error) {
	if hg[s.Hg] == nil {
		err = errors.New(fmt.Sprintf("Version %s of the human genome is not available", s.Hg))
		return
	}
	return
}
