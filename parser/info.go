package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Info struct {
	Info map[string]interface{}
}

func (info *Info) ExecAirport() (string, error) {
	return execAirport("-I")
}

func (info *Info) Parse(s string) error { //nolint:cyclop
	m := map[string]interface{}{
		"channel": []int{},
	}

	for _, line := range strings.Split(s, "\n") {
		kv := strings.Split(line, ":")
		if len(kv) < 2 {
			continue
		}

		k := strings.TrimSpace(kv[0])

		v := kv[1]
		if len(v) == 0 {
			/*
				m[k] = nil
				continue
			*/
			v = ""
		} else if v[0] == ' ' {
			v = v[1:]
		}

		switch k {
		case "agrCtlRSSI", "argExtRSSI", "agrCtlNoise", "agrExtNoise",
			"lastTxRate", "maxRate", "lastAssocStatus", "MCS", "guardInterval", "NSS":
			iv, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("on strconv.Atoi(%q): %w", v, err)
			}

			m[k] = iv
		case "channel":
			for _, vv := range strings.Split(v, ",") {
				c, err := strconv.Atoi(vv)
				if err != nil {
					return fmt.Errorf("on strconv.Atoi(%q): %w", vv, err)
				}

				m[k] = append(m[k].([]int), c) //nolint:forcetypeassert
			}
		case "802.11 auth":
			m["IEEE80211Auth"] = v
		case "link auth":
			m["linkAuth"] = v
		case "op mode":
			m["opMode"] = v

		default:
			m[k] = v
		}
	}

	info.Info = m

	return nil
}
