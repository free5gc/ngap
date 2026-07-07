package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSMutingConfigurationEUTRA struct {
	Choice PRSMutingConfigurationEUTRAAlt
}

type PRSMutingConfigurationEUTRAAlt interface {
	PRSMutingConfigurationEUTRAAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type TwoForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt0 *TwoForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(0)
}

func (x *TwoForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 2, 2
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *TwoForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 2, 2
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type FourForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt1 *FourForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(1)
}

func (x *FourForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 4, 4
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *FourForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 4, 4
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type EightForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt2 *EightForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(2)
}

func (x *EightForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 8, 8
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *EightForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 8, 8
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type SixteenForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt3 *SixteenForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(3)
}

func (x *SixteenForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *SixteenForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 16, 16
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ThirtyTwoForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt4 *ThirtyTwoForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(4)
}

func (x *ThirtyTwoForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 32, 32
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *ThirtyTwoForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 32, 32
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type SixtyFourForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt5 *SixtyFourForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(5)
}

func (x *SixtyFourForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 64, 64
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *SixtyFourForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 64, 64
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type OneHundredAndTwentyEightForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt6 *OneHundredAndTwentyEightForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(6)
}

func (x *OneHundredAndTwentyEightForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 128, 128
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *OneHundredAndTwentyEightForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 128, 128
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type TwoHundredAndFiftySixForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt7 *TwoHundredAndFiftySixForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(7)
}

func (x *TwoHundredAndFiftySixForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 256, 256
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *TwoHundredAndFiftySixForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 256, 256
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type FiveHundredAndTwelveForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt8 *FiveHundredAndTwelveForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(8)
}

func (x *FiveHundredAndTwelveForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 512, 512
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *FiveHundredAndTwelveForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 512, 512
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA struct {
	Value aper.BitString
}

func (alt9 *OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(9)
}

func (x *OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 1024, 1024
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 1024, 1024
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt10 ProtocolIESingleContainerPRSMutingConfigurationEUTRAExtensionIE) PRSMutingConfigurationEUTRAAltIndex() int64 {
	return int64(10)
}

// Choice Type Read/Write Functions

func (x *PRSMutingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 10
	var option_idx int64 = x.Choice.PRSMutingConfigurationEUTRAAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *PRSMutingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 10
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 1 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 2 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 3 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 4 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 5 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 6 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 7 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 8 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 9 {
		x.Choice = new(OneThousandAndTwentyFourForPRSMutingConfigurationEUTRA)
	} else if option_idx == 10 {
		x.Choice = new(ProtocolIESingleContainerPRSMutingConfigurationEUTRAExtensionIE)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
