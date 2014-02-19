/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package txaddr

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/reg/xaddr"
)

/* use to create XAddress for TX_ADDR

   TX_ADDR
   Transmit address. Used for a PTX device only.
   (LSByte is written first)
   Set RX_ADDR_P0 equal to this address to handle
   automatic acknowledge if this is a PTX device with
   Enhanced ShockBurstTM enabled. */
func New(flags xaddr.A) *xaddr.Full {
	return xaddr.NewFull(addr.TX_ADDR, flags)
}
