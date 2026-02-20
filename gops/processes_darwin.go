//go:build darwin

package gops

import "fmt"

func getPssDirty(_ int32) (uint64, error) {
	return 0, fmt.Errorf("pss dirty is not supported on darwin")
}
