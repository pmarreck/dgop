package gops

import "github.com/AvengeMedia/dgop/models"

func (self *GopsUtil) GetNetworkInfo() ([]*models.NetworkInfo, error) {
	netIO, err := self.netProvider.IOCounters(true)
	if err != nil {
		return nil, err
	}

	ifaces, _ := self.netProvider.Interfaces()
	index := indexInterfacesByName(ifaces)

	res := make([]*models.NetworkInfo, 0)
	for _, n := range netIO {
		if isUsableNetworkInterface(n.Name, index) {
			res = append(res, &models.NetworkInfo{
				Name: n.Name,
				Rx:   n.BytesRecv,
				Tx:   n.BytesSent,
			})
		}
	}
	return res, nil
}
