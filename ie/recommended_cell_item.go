package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type RecommendedCellItem struct {
	NGRANCGI         *NGRANCGI                                            // valueLB:0,valueUB:2
	TimeStayedInCell *int64                                               // valueLB:0,valueUB:4095,optional
	IEExtensions     *ProtocolExtensionContainerRecommendedCellItemExtIEs // optional
}

func (x *RecommendedCellItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RecommendedCellItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCGI == nil {
		return errors.Errorf("NGRANCGI is missing")
	}
	// optional field
	if x.TimeStayedInCell != nil {
		RecommendedCellItemOptPresentFlag = append(RecommendedCellItemOptPresentFlag, true)
	} else {
		RecommendedCellItemOptPresentFlag = append(RecommendedCellItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RecommendedCellItemOptPresentFlag = append(RecommendedCellItemOptPresentFlag, true)
	} else {
		RecommendedCellItemOptPresentFlag = append(RecommendedCellItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RecommendedCellItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCGI marshal failed")
	}

	// optional field
	if x.TimeStayedInCell != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 4095
		err = pd.WriteInteger(*(x.TimeStayedInCell), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *RecommendedCellItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RecommendedCellItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&RecommendedCellItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCGI = new(NGRANCGI)
	err = x.NGRANCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCGI error")
	}

	// optional field (optPresentFlag index: 0)
	if RecommendedCellItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 4095
		x.TimeStayedInCell = new(int64)
		*(x.TimeStayedInCell), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if RecommendedCellItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRecommendedCellItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
