package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type MeasuredResultsValue struct {
	Choice MeasuredResultsValueAlt
}

type MeasuredResultsValueAlt interface {
	MeasuredResultsValueAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ValueAngleOfArrivalEUTRAForMeasuredResultsValue struct {
	Value int64
}

func (alt0 *ValueAngleOfArrivalEUTRAForMeasuredResultsValue) MeasuredResultsValueAltIndex() int64 {
	return int64(0)
}

func (x *ValueAngleOfArrivalEUTRAForMeasuredResultsValue) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 719
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *ValueAngleOfArrivalEUTRAForMeasuredResultsValue) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 719
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ValueTimingAdvanceType1EUTRAForMeasuredResultsValue struct {
	Value int64
}

func (alt1 *ValueTimingAdvanceType1EUTRAForMeasuredResultsValue) MeasuredResultsValueAltIndex() int64 {
	return int64(1)
}

func (x *ValueTimingAdvanceType1EUTRAForMeasuredResultsValue) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 7690
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *ValueTimingAdvanceType1EUTRAForMeasuredResultsValue) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 7690
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ValueTimingAdvanceType2EUTRAForMeasuredResultsValue struct {
	Value int64
}

func (alt2 *ValueTimingAdvanceType2EUTRAForMeasuredResultsValue) MeasuredResultsValueAltIndex() int64 {
	return int64(2)
}

func (x *ValueTimingAdvanceType2EUTRAForMeasuredResultsValue) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 7690
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *ValueTimingAdvanceType2EUTRAForMeasuredResultsValue) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 7690
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 ResultRSRPEUTRA) MeasuredResultsValueAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 ResultRSRQEUTRA) MeasuredResultsValueAltIndex() int64 {
	return int64(4)
}

// Choice type and its Read/Write is defined elsewhere
func (alt5 ProtocolIESingleContainerMeasuredResultsValueExtensionIE) MeasuredResultsValueAltIndex() int64 {
	return int64(5)
}

// Choice Type Read/Write Functions

func (x *MeasuredResultsValue) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 5
	var option_idx int64 = x.Choice.MeasuredResultsValueAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *MeasuredResultsValue) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 5
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(ValueTimingAdvanceType2EUTRAForMeasuredResultsValue)
	} else if option_idx == 1 {
		x.Choice = new(ValueTimingAdvanceType2EUTRAForMeasuredResultsValue)
	} else if option_idx == 2 {
		x.Choice = new(ValueTimingAdvanceType2EUTRAForMeasuredResultsValue)
	} else if option_idx == 3 {
		x.Choice = new(ResultRSRPEUTRA)
	} else if option_idx == 4 {
		x.Choice = new(ResultRSRQEUTRA)
	} else if option_idx == 5 {
		x.Choice = new(ProtocolIESingleContainerMeasuredResultsValueExtensionIE)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
