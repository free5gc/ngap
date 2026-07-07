package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type GNBRxTxTimeDiffMeas struct {
	Choice GNBRxTxTimeDiffMeasAlt
}

type GNBRxTxTimeDiffMeasAlt interface {
	GNBRxTxTimeDiffMeasAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K0ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt0 *K0ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(0)
}

func (x *K0ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 1970049
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K0ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 1970049
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K1ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt1 *K1ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(1)
}

func (x *K1ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 985025
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K1ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 985025
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K2ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt2 *K2ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(2)
}

func (x *K2ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 492513
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K2ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 492513
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K3ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt3 *K3ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(3)
}

func (x *K3ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 246257
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K3ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 246257
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K4ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt4 *K4ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(4)
}

func (x *K4ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 123129
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K4ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 123129
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K5ForGNBRxTxTimeDiffMeas struct {
	Value int64
}

func (alt5 *K5ForGNBRxTxTimeDiffMeas) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(5)
}

func (x *K5ForGNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 61565
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K5ForGNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 61565
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 ProtocolIESingleContainerGNBRxTxTimeDiffMeasExtIEs) GNBRxTxTimeDiffMeasAltIndex() int64 {
	return int64(6)
}

// Choice Type Read/Write Functions

func (x *GNBRxTxTimeDiffMeas) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64 = x.Choice.GNBRxTxTimeDiffMeasAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *GNBRxTxTimeDiffMeas) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 1 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 2 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 3 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 4 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 5 {
		x.Choice = new(K5ForGNBRxTxTimeDiffMeas)
	} else if option_idx == 6 {
		x.Choice = new(ProtocolIESingleContainerGNBRxTxTimeDiffMeasExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
