package pfcpType

import (
	"fmt"
)

type RequestedQosMonitoring struct {
	DLPD   bool //Downlink Packet Delay
	ULPD   bool //Uplink Packet Delay
	RPPD   bool //Round Trip Packet Delay
	GTPUPM bool //GTP-U Path Monitoring
	DLCI   bool //Downlink Congestion Information
	ULCI   bool //Uplink Congestion Information
	DLDR   bool //Downlink Data Rate
	ULDR   bool //Uplink Data Rate
}

func (rQm *RequestedQosMonitoring) MarshalBinary() (data []byte, err error) {
	octet5 := btou(rQm.ULDR)<<7 |
		btou(rQm.DLDR)<<6 |
		btou(rQm.ULCI)<<5 |
		btou(rQm.DLCI)<<4 |
		btou(rQm.GTPUPM)<<3 |
		btou(rQm.RPPD)<<2 |
		btou(rQm.ULPD)<<1 |
		btou(rQm.DLPD)
	data = append([]byte(""), octet5)

	if octet5 == 0 {
		return []byte(""), fmt.Errorf("none of the bits is set!")
	}

	return data, nil
}

func (rQm *RequestedQosMonitoring) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var index uint16 = 0
	if length < index+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	tmpUint8 := data[index]
	rQm.DLPD = utob(tmpUint8 & (1 << 0))
	rQm.ULPD = utob(tmpUint8 & (1 << 1))
	rQm.RPPD = utob(tmpUint8 & (1 << 2))
	rQm.GTPUPM = utob(tmpUint8 & (1 << 3))
	rQm.DLCI = utob(tmpUint8 & (1 << 4))
	rQm.ULCI = utob(tmpUint8 & (1 << 5))
	rQm.DLDR = utob(tmpUint8 & (1 << 6))
	rQm.ULDR = utob(tmpUint8 & (1 << 7))
	index = index + 1

	if tmpUint8 == 0 {
		return fmt.Errorf("none of the bits is set!")
	}

	if length != index {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
