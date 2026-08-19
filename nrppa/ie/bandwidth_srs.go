package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	BandwidthSRSFR1PresentMHz5   aper.Enumerated = 0
	BandwidthSRSFR1PresentMHz10  aper.Enumerated = 1
	BandwidthSRSFR1PresentMHz20  aper.Enumerated = 2
	BandwidthSRSFR1PresentMHz40  aper.Enumerated = 3
	BandwidthSRSFR1PresentMHz50  aper.Enumerated = 4
	BandwidthSRSFR1PresentMHz80  aper.Enumerated = 5
	BandwidthSRSFR1PresentMHz100 aper.Enumerated = 6
	BandwidthSRSFR1PresentMHz15  aper.Enumerated = 7
	BandwidthSRSFR1PresentMHz25  aper.Enumerated = 8
	BandwidthSRSFR1PresentMHz30  aper.Enumerated = 9
	BandwidthSRSFR1PresentMHz60  aper.Enumerated = 10
	BandwidthSRSFR1PresentMHz35  aper.Enumerated = 11
	BandwidthSRSFR1PresentMHz45  aper.Enumerated = 12
	BandwidthSRSFR1PresentMHz70  aper.Enumerated = 13
	BandwidthSRSFR1PresentMHz90  aper.Enumerated = 14
)

const ( /* Enum Type */
	BandwidthSRSFR2PresentMHz50   aper.Enumerated = 0
	BandwidthSRSFR2PresentMHz100  aper.Enumerated = 1
	BandwidthSRSFR2PresentMHz200  aper.Enumerated = 2
	BandwidthSRSFR2PresentMHz400  aper.Enumerated = 3
	BandwidthSRSFR2PresentMhz800  aper.Enumerated = 4
	BandwidthSRSFR2PresentMHz1600 aper.Enumerated = 5
	BandwidthSRSFR2PresentMHz2000 aper.Enumerated = 6
)

type BandwidthSRS struct {
	Choice BandwidthSRSAlt
}

type BandwidthSRSAlt interface {
	BandwidthSRSAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type FR1ForBandwidthSRS struct {
	Value aper.Enumerated
}

func (alt0 *FR1ForBandwidthSRS) BandwidthSRSAltIndex() int64 {
	return int64(0)
}

func (x *FR1ForBandwidthSRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 6
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *FR1ForBandwidthSRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 6
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type FR2ForBandwidthSRS struct {
	Value aper.Enumerated
}

func (alt1 *FR2ForBandwidthSRS) BandwidthSRSAltIndex() int64 {
	return int64(1)
}

func (x *FR2ForBandwidthSRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *FR2ForBandwidthSRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 3
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerBandwidthSRSExtIEs) BandwidthSRSAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *BandwidthSRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.BandwidthSRSAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *BandwidthSRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(FR2ForBandwidthSRS)
	} else if option_idx == 1 {
		x.Choice = new(FR2ForBandwidthSRS)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerBandwidthSRSExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
