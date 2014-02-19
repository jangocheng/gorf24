/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package feature

import (
	"testing"
	
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNewFeature_RegisterAddress_IsFEATURE(t *testing.T) {
	f := New(0)
	expected := addr.FEATURE

	actual := f.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestSetDynamicAck_Enabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("00000000"))
	expected := util.B("00000001")

	f.SetDynamicAck(true)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestSetDynamicAck_Disabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("11111111"))
	expected := util.B("11111110")

	f.SetDynamicAck(false)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestIsDynamicAckEnabled_BitZero_ReturnsFalse(t *testing.T) {
	f := New(util.B("11111110"))
	expected := false

	actual := f.IsDynamicAckEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}

func TestIsDynamicAckEnabled_BitOne_ReturnsTrue(t *testing.T) {
	f := New(util.B("00000001"))
	expected := true

	actual := f.IsDynamicAckEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}

func TestSetPayloadWithAck_Enabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("00000010"))
	expected := util.B("00000010")

	f.SetPayloadWithAck(true)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestSetPayloadWithAck_Disabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("11111111"))
	expected := util.B("11111101")

	f.SetPayloadWithAck(false)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestIsPayloadWithAckEnabled_BitZero_ReturnsFalse(t *testing.T) {
	f := New(util.B("11111101"))
	expected := false

	actual := f.IsPayloadWithAckEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}

func TestIsPayloadWithAckEnabled_BitOne_ReturnsTrue(t *testing.T) {
	f := New(util.B("00000010"))
	expected := true

	actual := f.IsPayloadWithAckEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}

func TestSetDynamicPayloadLength_Enabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("00000000"))
	expected := util.B("00000100")

	f.SetDynamicPayloadLength(true)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestSetDynamicPayloadLength_Disabled_FlipsRightBit(t *testing.T) {
	f := New(util.B("00000100"))
	expected := util.B("00000000")

	f.SetDynamicPayloadLength(false)

	actual := f.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with feature '%v'", expected, actual, f)
	}
}

func TestIsDynamicPayloadLengthEnabled_BitZero_ReturnsFalse(t *testing.T) {
	f := New(util.B("11111011"))
	expected := false

	actual := f.IsDynamicPayloadLengthEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}

func TestIsDynamicPayloadLengthEnabled_BitOne_ReturnsTrue(t *testing.T) {
	f := New(util.B("00000100"))
	expected := true

	actual := f.IsDynamicPayloadLengthEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with feature '%v'", expected, actual, f)
	}
}




















