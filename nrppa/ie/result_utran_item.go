package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultUTRANItemPhysCellIDUTRAN struct {
	Choice ResultUTRANItemPhysCellIDUTRANAlt
}

type ResultUTRANItemPhysCellIDUTRANAlt interface {
	ResultUTRANItemPhysCellIDUTRANAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 PhysCellIDUTRAFDD) ResultUTRANItemPhysCellIDUTRANAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 PhysCellIDUTRATDD) ResultUTRANItemPhysCellIDUTRANAltIndex() int64 {
	return int64(1)
}

// Choice Type Read/Write Functions

func (x *ResultUTRANItemPhysCellIDUTRAN) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 1
	var option_idx int64 = x.Choice.ResultUTRANItemPhysCellIDUTRANAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *ResultUTRANItemPhysCellIDUTRAN) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 1
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(PhysCellIDUTRAFDD)
	} else if option_idx == 1 {
		x.Choice = new(PhysCellIDUTRATDD)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}

type ResultUTRANItem struct {
	UARFCN       *UARFCN
	UTRARSCP     *UTRARSCP                                        // optional
	UTRAEcN0     *UTRAEcN0                                        // optional
	IEExtensions *ProtocolExtensionContainerResultUTRANItemExtIEs // optional
}

func (x *ResultUTRANItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultUTRANItemOptPresentFlag := []bool{}
	// mandatory field
	if x.UARFCN == nil {
		return errors.Errorf("UARFCN is missing")
	}
	// optional field
	if x.UTRARSCP != nil {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, true)
	} else {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, false)
	}
	// optional field
	if x.UTRAEcN0 != nil {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, true)
	} else {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, true)
	} else {
		ResultUTRANItemOptPresentFlag = append(ResultUTRANItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultUTRANItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UARFCN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UARFCN marshal failed")
	}

	// optional field
	if x.UTRARSCP != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UTRARSCP.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UTRARSCP marshal failed")
		}
	}

	// optional field
	if x.UTRAEcN0 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UTRAEcN0.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UTRAEcN0 marshal failed")
		}
	}

	// optional field
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *ResultUTRANItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultUTRANItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ResultUTRANItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UARFCN = new(UARFCN)
	err = x.UARFCN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UARFCN error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultUTRANItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UTRARSCP = new(UTRARSCP)
		err = x.UTRARSCP.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UTRARSCP error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultUTRANItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.UTRAEcN0 = new(UTRAEcN0)
		err = x.UTRAEcN0.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UTRAEcN0 error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultUTRANItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultUTRANItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
