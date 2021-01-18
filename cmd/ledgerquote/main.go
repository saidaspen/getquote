package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Setup command line argument passing.
	// We use our own FlagSet, as we want to first (positional) argument to be the price file.
	argSet := flag.NewFlagSet("", flag.ExitOnError)
	argSet.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: lq <pricefile>\n")
		argSet.PrintDefaults()
		os.Exit(-1)
	}

	// Flags
	fromIn := argSet.String("from", "", "Date range start (inclusive)")
	toIn := argSet.String("to", "", "Date range end (inclusive)")
	configIn := argSet.String("config", "", "Config file to use")
	overwrite := argSet.Bool("nooverwrite", false, "Don't overwrite existing prices")

	// Check inputs and parse them
	if len(os.Args) < 2 {
		argSet.Usage()
	}
	priceFileIn := os.Args[1]
	err := argSet.Parse(os.Args[2:])
	if err != nil || argSet.NArg() != 0 {
		argSet.Usage()
	}

	fmt.Println("Price file: ", priceFileIn)
	fmt.Println("From: ", *fromIn)
	fmt.Println("To: ", *toIn)
	fmt.Println("Config: ", *configIn)
	fmt.Println("Overwrite: ", *overwrite)
}

func setupFlags(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Println("foo bar")
	}
}
