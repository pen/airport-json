package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Plist struct {
	Info map[string]interface{}
}

func (plist *Plist) ExecAirport() (string, error) {
	return execAirport("-I", "-x")
}

func (plist *Plist) Parse(s string) error {
	m := map[string]interface{}{
		"RSSI_CTL_LIST": []int{},
		"RSSI_EXT_LIST": []int{},
	}

	k := ""

	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "<key>") {
			k = strings.TrimPrefix(strings.TrimSuffix(line, "</key>"), "<key>")

			continue
		}

		if k == "" {
			continue
		}

		if !strings.HasPrefix(line, "<integer>") {
			continue
		}

		s := strings.TrimPrefix(
			strings.TrimSuffix(line, "</integer>"),
			"<integer>",
		)

		v, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("on strconv.Atoi(%q): %w", s, err)
		}

		switch k {
		case "RSSI_CTL_LIST":
			m[k] = append(m[k].([]int), v) //nolint:forcetypeassert
		case "RSSI_EXT_LIST":
			m[k] = append(m[k].([]int), v) //nolint:forcetypeassert
		default:
			m[k] = v
		}
	}

	plist.Info = m

	return nil
}
