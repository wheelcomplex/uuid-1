package uuid

import (
	"encoding/binary"
	"fmt"
)

// The UUID represents Universally Unique IDentifier (which is 128 bit long).
type UUID [16]byte

// NIL is defined in RFC 4122 section 4.1.7.
// The nil UUID is special form of UUID that is specified to have all 128 bits set to zero.
var NIL = &UUID{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// Version of the UUID represents a kind of subtype specifier.
func (u *UUID) Version() int {
	return int(binary.BigEndian.Uint16(u[6:8]) >> 12)
}

// String returns the human readable form of the UUID.
func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func (u *UUID) variantRFC4122() {
	u[8] = (u[8] & 0x3f) | 0x80
}
