package pfcpType

import (
	"fmt"
)

type SRRID struct {
	SrrIdValue uint8
}

func (srr *SRRID) MarshalBinary() (data []byte, err error) {
	data = append([]byte(""), srr.SrrIdValue)
	return data, nil
}

func (srr *SRRID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var index uint16 = 0
	if length < index+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	srr.SrrIdValue = data[index]
	index = index + 1
	if length != index {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	return nil
}
