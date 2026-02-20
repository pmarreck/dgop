package tui

import (
	"testing"

	"github.com/AvengeMedia/dgop/models"
	"github.com/stretchr/testify/require"
)

func TestSelectBestNetworkInterfacePrefersPrimaryOverBridge(t *testing.T) {
	m := &ResponsiveTUIModel{}
	interfaces := []*models.NetworkRateInfo{
		{Interface: "bridge102", RxRate: 8000, TxRate: 2000, RxTotal: 1000000, TxTotal: 1000000},
		{Interface: "en0", RxRate: 1000, TxRate: 1000, RxTotal: 2000000, TxTotal: 3000000},
	}

	best := m.selectBestNetworkInterface(interfaces)
	require.NotNil(t, best)
	require.Equal(t, "en0", best.Interface)
}

func TestSelectBestNetworkInterfaceSupportsUtun(t *testing.T) {
	m := &ResponsiveTUIModel{}
	interfaces := []*models.NetworkRateInfo{
		{Interface: "utun4", RxRate: 3000, TxRate: 2000, RxTotal: 100000, TxTotal: 100000},
		{Interface: "bridge0", RxRate: 0, TxRate: 0, RxTotal: 0, TxTotal: 0},
	}

	best := m.selectBestNetworkInterface(interfaces)
	require.NotNil(t, best)
	require.Equal(t, "utun4", best.Interface)
}
