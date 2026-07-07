package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TransmissionComb struct {
	Choice TransmissionCombAlt
}

type TransmissionCombAlt interface {
	TransmissionCombAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 TransmissionCombN2) TransmissionCombAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TransmissionCombN4) TransmissionCombAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerTransmissionCombExtIEs) TransmissionCombAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *TransmissionComb) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.TransmissionCombAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *TransmissionComb) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(TransmissionCombN2)
	} else if option_idx == 1 {
		x.Choice = new(TransmissionCombN4)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerTransmissionCombExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}

type TransmissionCombN2 struct {
	CombOffsetN2  *int64 // valueLB:0,valueUB:1
	CyclicShiftN2 *int64 // valueLB:0,valueUB:7
}

func (x *TransmissionCombN2) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TransmissionCombN2OptPresentFlag := []bool{}
	// mandatory field
	if x.CombOffsetN2 == nil {
		return errors.Errorf("CombOffsetN2 is missing")
	}
	// mandatory field
	if x.CyclicShiftN2 == nil {
		return errors.Errorf("CyclicShiftN2 is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TransmissionCombN2OptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteInteger(*(x.CombOffsetN2), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteInteger(*(x.CyclicShiftN2), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	return nil
}

func (x *TransmissionCombN2) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TransmissionCombN2OptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TransmissionCombN2OptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1
	x.CombOffsetN2 = new(int64)
	*(x.CombOffsetN2), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 7
	x.CyclicShiftN2 = new(int64)
	*(x.CyclicShiftN2), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	return nil
}

type TransmissionCombN4 struct {
	CombOffsetN4  *int64 // valueLB:0,valueUB:3
	CyclicShiftN4 *int64 // valueLB:0,valueUB:11
}

func (x *TransmissionCombN4) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TransmissionCombN4OptPresentFlag := []bool{}
	// mandatory field
	if x.CombOffsetN4 == nil {
		return errors.Errorf("CombOffsetN4 is missing")
	}
	// mandatory field
	if x.CyclicShiftN4 == nil {
		return errors.Errorf("CyclicShiftN4 is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TransmissionCombN4OptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteInteger(*(x.CombOffsetN4), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 11
	err = pd.WriteInteger(*(x.CyclicShiftN4), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	return nil
}

func (x *TransmissionCombN4) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TransmissionCombN4OptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TransmissionCombN4OptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3
	x.CombOffsetN4 = new(int64)
	*(x.CombOffsetN4), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 11
	x.CyclicShiftN4 = new(int64)
	*(x.CyclicShiftN4), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	return nil
}
