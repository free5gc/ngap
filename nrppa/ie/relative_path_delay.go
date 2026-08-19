package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type RelativePathDelay struct {
	Choice RelativePathDelayAlt
}

type RelativePathDelayAlt interface {
	RelativePathDelayAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K0ForRelativePathDelay struct {
	Value int64
}

func (alt0 *K0ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(0)
}

func (x *K0ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 16351
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K0ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 16351
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K1ForRelativePathDelay struct {
	Value int64
}

func (alt1 *K1ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(1)
}

func (x *K1ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 8176
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K1ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 8176
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K2ForRelativePathDelay struct {
	Value int64
}

func (alt2 *K2ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(2)
}

func (x *K2ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 4088
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K2ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 4088
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K3ForRelativePathDelay struct {
	Value int64
}

func (alt3 *K3ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(3)
}

func (x *K3ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 2044
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K3ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 2044
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K4ForRelativePathDelay struct {
	Value int64
}

func (alt4 *K4ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(4)
}

func (x *K4ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 1022
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K4ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 1022
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type K5ForRelativePathDelay struct {
	Value int64
}

func (alt5 *K5ForRelativePathDelay) RelativePathDelayAltIndex() int64 {
	return int64(5)
}

func (x *K5ForRelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 511
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *K5ForRelativePathDelay) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 511
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 ProtocolIESingleContainerRelativePathDelayExtIEs) RelativePathDelayAltIndex() int64 {
	return int64(6)
}

// Choice Type Read/Write Functions

func (x *RelativePathDelay) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64 = x.Choice.RelativePathDelayAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *RelativePathDelay) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 1 {
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 2 {
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 3 {
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 4 {
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 5 {
		x.Choice = new(K5ForRelativePathDelay)
	} else if option_idx == 6 {
		x.Choice = new(ProtocolIESingleContainerRelativePathDelayExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
