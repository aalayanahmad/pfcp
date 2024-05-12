package pfcpType

import (
	"encoding/binary"
	"fmt"
)

type PacketDelayThresholds struct {
	DL                               bool   //Downlink packet delay should be there
	UL                               bool   //Uplink packet delay should be there
	RP                               bool   //Round trip packet delay should be there
	DownlinkPacketDelayThresholdRID  uint32 //Downlink packet delay threshold
	UpPacketDelayThresholdRID        uint32 //Uplink packet delay threshold
	RoundTripPacketDelayThresholdRID uint32 //Round trip packet delay threshold
}

func (pDt *PacketDelayThresholds) MarshalBinary() (data []byte, err error) {
	octet5 := btou(pDt.RP)<<2 |
		btou(pDt.UL)<<1 |
		btou(pDt.DL)
	data = append([]byte(""), octet5)

	if octet5 == 0 {
		return []byte(""), fmt.Errorf("none of the bits is set!")
	}

	if pDt.DL {
		dlBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(dlBytes, pDt.DownlinkPacketDelayThresholdRID)
		data = append(data, dlBytes...)
	}

	if pDt.UL {
		ulBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(ulBytes, pDt.UpPacketDelayThresholdRID)
		data = append(data, ulBytes...)
	}

	if pDt.RP {
		rpBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(rpBytes, pDt.RoundTripPacketDelayThresholdRID)
		data = append(data, rpBytes...)
	}

	return data, nil
}

func (pDt *PacketDelayThresholds) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var index uint16 = 0
	if length < index+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	tmpUint8 := data[index]
	pDt.DL = utob(tmpUint8 & (1 << 0))
	pDt.UL = utob(tmpUint8 & (1 << 1))
	pDt.RP = utob(tmpUint8 & (1 << 2))
	index = index + 1

	if tmpUint8 == 0 {
		return fmt.Errorf("none of the bits is set!")
	}

	if pDt.DL {
		if length < index+4 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		pDt.DownlinkPacketDelayThresholdRID = binary.BigEndian.Uint32(data[index : index+4])
		index = index + 4
	}

	if pDt.UL {
		if length < index+4 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		pDt.UpPacketDelayThresholdRID = binary.BigEndian.Uint32(data[index : index+4])
		index = index + 4
	}

	if pDt.RP {
		if length < index+4 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		pDt.RoundTripPacketDelayThresholdRID = binary.BigEndian.Uint32(data[index : index+4])
		index = index + 4
	}
	if length != index {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
