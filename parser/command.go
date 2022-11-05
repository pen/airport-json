package parser

import (
	"fmt"
	"os/exec"
)

const AirportPath = `/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport`

func execAirport(arg ...string) (string, error) {
	out, err := exec.Command(AirportPath, arg...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(`on exec.Command(%q, %q).CombinedOutput(): %w`, AirportPath, arg, err)
	}

	return string(out), nil
}
