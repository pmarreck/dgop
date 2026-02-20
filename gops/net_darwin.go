//go:build darwin

package gops

import "strings"

func matchesNetworkInterface(name string) bool {
	prefixes := []string{"en", "utun", "bridge", "awdl", "llw"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}
