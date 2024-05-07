package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalSRRID(t *testing.T) {
	testData := SRRID{
		SrrIdValue: 35,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{42}, buf)
}

func TestUnmarshalSRRID(t *testing.T) {
	buf := []byte{35}
	var testData SRRID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := SRRID{
		SrrIdValue: 35,
	}
	assert.Equal(t, expectData, testData)
}
