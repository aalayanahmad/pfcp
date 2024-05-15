package pfcpType

import (
	"fmt"
)

type ReportingFrequency struct {
	EVET     bool
	PERIO    bool
	RESERVED bool
}

func (r *ReportingFrequency) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(r.RESERVED)<<2 |
		btou(r.PERIO)<<1 |
		btou(r.EVET)
	data = append([]byte(""), tmpUint8)

	return data, nil
}

func (r *ReportingFrequency) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	r.RESERVED = utob(data[idx] & BitMask3)
	r.PERIO = utob(data[idx] & BitMask2)
	r.EVET = utob(data[idx] & BitMask1)
	idx = idx + 1

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
