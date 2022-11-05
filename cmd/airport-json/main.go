package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/pen/airport-json/parser"
)

var version = "0.0.1"

func main() {
	var optXML, optBoth, optMerge, optVersion bool

	flag.BoolVar(&optXML, "x", false, `get "-Ix" info`)
	flag.BoolVar(&optBoth, "b", false, "get both")
	flag.BoolVar(&optMerge, "m", false, "flip merge")
	flag.BoolVar(&optVersion, "v", false, "show version")
	flag.Parse()

	if optVersion {
		fmt.Printf("%s\n", version) //nolint:forbidigo

		return
	}

	getInfo := true
	getPlist := false
	doMerge := true

	if optXML {
		getInfo = false
		getPlist = true
	}

	if optBoth {
		getInfo = true
		getPlist = true
		doMerge = false
	}

	if optMerge {
		doMerge = !doMerge
	}

	if err := run(getInfo, getPlist, doMerge); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func run(getInfo, getPlist, doMerge bool) error { //nolint:cyclop
	mm := map[string]interface{}{}

	if getInfo {
		info := &parser.Info{}

		out, err := info.ExecAirport()
		if err != nil {
			return fmt.Errorf("on info.ExecAirport(): %w", err)
		}

		if err := info.Parse(out); err != nil {
			return fmt.Errorf("on info.Parse(): %w", err)
		}

		if doMerge {
			for k, v := range info.Info {
				mm[k] = v
			}
		} else {
			mm["info"] = info.Info
		}
	}

	if getPlist {
		plist := &parser.Plist{}

		out, err := plist.ExecAirport()
		if err != nil {
			return fmt.Errorf(`on plist.ExecAirport(): %w`, err)
		}

		if err := plist.Parse(out); err != nil {
			return fmt.Errorf("on ParsePlist(): %w", err)
		}

		if doMerge {
			for k, v := range plist.Info {
				mm[k] = v
			}
		} else {
			mm["plist"] = plist.Info
		}
	}

	b, err := json.Marshal(mm)
	if err != nil {
		return fmt.Errorf("on json.Marshal(): %w", err)
	}

	fmt.Printf("%s\n", string(b)) //nolint:forbidigo

	return nil
}
