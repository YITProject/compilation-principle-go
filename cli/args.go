package cli

import (
	"github.com/startracex/argp"
)

var (
	Table   = false
	Default = false
)

func init() {
	ap := argp.New()
	ap.BoolVar(&Table, "-t", "--table")
	ap.BoolVar(&Default, "-d")
}
