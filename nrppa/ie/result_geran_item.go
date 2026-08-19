package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultGERANItem struct {
	BCCH            *BCCH
	PhysCellIDGERAN *PhysCellIDGERAN
	RSSI            *RSSI
	IEExtensions    *ProtocolExtensionContainerResultGERANItemExtIEs // optional
}

func (x *ResultGERANItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultGERANItemOptPresentFlag := []bool{}
	// mandatory field
	if x.BCCH == nil {
		return errors.Errorf("BCCH is missing")
	}
	// mandatory field
	if x.PhysCellIDGERAN == nil {
		return errors.Errorf("PhysCellIDGERAN is missing")
	}
	// mandatory field
	if x.RSSI == nil {
		return errors.Errorf("RSSI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultGERANItemOptPresentFlag = append(ResultGERANItemOptPresentFlag, true)
	} else {
		ResultGERANItemOptPresentFlag = append(ResultGERANItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultGERANItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.BCCH.Write(pd)
	if err != nil {
		return errors.Wrap(err, "BCCH marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PhysCellIDGERAN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PhysCellIDGERAN marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.RSSI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RSSI marshal failed")
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

func (x *ResultGERANItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultGERANItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResultGERANItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.BCCH = new(BCCH)
	err = x.BCCH.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode BCCH error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PhysCellIDGERAN = new(PhysCellIDGERAN)
	err = x.PhysCellIDGERAN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PhysCellIDGERAN error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RSSI = new(RSSI)
	err = x.RSSI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RSSI error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultGERANItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultGERANItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
