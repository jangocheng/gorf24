/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package feature

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* FEATURE
   Feature Register
   bits 7:3 reserved */
type F struct {
	reg.R
}

func New(flags byte) *F {
	return &F{reg.New(addr.FEATURE, flags)}
}

/* EN_DYN_ACK (bit 0)
   Enables the W_TX_PAYLOAD_NOACK command
   xxxxxxx0 -> disabled
   xxxxxxx1 -> enabled */
func (f *F) IsDynamicAckEnabled() bool {
	return f.Byte()&0x01 == 0x01
}
func (f *F) SetDynamicAck(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x01)
	} else {
		f.R.Set(f.Byte() & 0xFE)
	}
}

/* EN_ACK_PAY (bit 1)
   Enables Payload with ACK

   If ACK packet payload is activated, ACK packets have
   dynamic payload lengths and the Dynamic PayloadLength
   feature should be enabled for pipe 0 on the PTX and PRX.
   This is to ensure that they receive the ACK packets
   with payloads. If the ACK payload is more than 15 byte
   in 2Mbps mode the ARD must be 500μS or more, and if the
   ACK payload is more than 5 byte in 1Mbps mode the ARD
   must be 500μS or more. In 250kbps mode (even when the
   payload is not in ACK) the ARD must be 500μS or more.

   xxxxxx0x -> disabled
   xxxxxx1x -> enabled */
func (f *F) IsPayloadWithAckEnabled() bool {
	return f.Byte()&0x02 == 0x02
}
func (f *F) SetPayloadWithAck(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x02)
	} else {
		f.R.Set(f.Byte() & 0xFD)
	}
}

/* EN_DPL (bit 2)
   Enables Dynamic Payload Length

   xxxxx0xx -> disabled
   xxxxx1xx -> enabled */
func (f *F) IsDynamicPayloadLengthEnabled() bool {
	return f.Byte()&0x04 == 0x04
}
func (f *F) SetDynamicPayloadLength(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x04)
	} else {
		f.R.Set(f.Byte() & 0x7B)
	}
}
