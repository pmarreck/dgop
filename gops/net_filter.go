package gops

import (
	"strings"

	gnet "github.com/shirou/gopsutil/v4/net"
)

func indexInterfacesByName(ifaces []gnet.InterfaceStat) map[string]gnet.InterfaceStat {
	index := make(map[string]gnet.InterfaceStat, len(ifaces))
	for _, iface := range ifaces {
		index[iface.Name] = iface
	}
	return index
}

func hasInterfaceFlag(flags []string, needle string) bool {
	for _, flag := range flags {
		if strings.EqualFold(flag, needle) {
			return true
		}
	}
	return false
}

func isUsableNetworkInterface(name string, ifaceMap map[string]gnet.InterfaceStat) bool {
	if !matchesNetworkInterface(name) {
		return false
	}

	if len(ifaceMap) == 0 {
		return true
	}

	iface, ok := ifaceMap[name]
	if !ok {
		return true
	}

	if hasInterfaceFlag(iface.Flags, "loopback") {
		return false
	}
	if !hasInterfaceFlag(iface.Flags, "up") {
		return false
	}

	return true
}
