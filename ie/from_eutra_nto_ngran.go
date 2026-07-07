package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FromEUTRANtoNGRAN struct {
	SourceeNBID       *IntersystemSONeNBID                               // valueExt
	TargetNGRANnodeID *IntersystemSONNGRANnodeID                         // valueExt
	IEExtensions      *ProtocolExtensionContainerFromEUTRANtoNGRANExtIEs // optional
}

func (x *FromEUTRANtoNGRAN) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FromEUTRANtoNGRANOptPresentFlag := []bool{}
	// mandatory field
	if x.SourceeNBID == nil {
		return errors.Errorf("SourceeNBID is missing")
	}
	// mandatory field
	if x.TargetNGRANnodeID == nil {
		return errors.Errorf("TargetNGRANnodeID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		FromEUTRANtoNGRANOptPresentFlag = append(FromEUTRANtoNGRANOptPresentFlag, true)
	} else {
		FromEUTRANtoNGRANOptPresentFlag = append(FromEUTRANtoNGRANOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FromEUTRANtoNGRANOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SourceeNBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SourceeNBID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TargetNGRANnodeID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetNGRANnodeID marshal failed")
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

func (x *FromEUTRANtoNGRAN) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FromEUTRANtoNGRANOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FromEUTRANtoNGRANOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourceeNBID = new(IntersystemSONeNBID)
	err = x.SourceeNBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourceeNBID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetNGRANnodeID = new(IntersystemSONNGRANnodeID)
	err = x.TargetNGRANnodeID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetNGRANnodeID error")
	}

	// optional field (optPresentFlag index: 0)
	if FromEUTRANtoNGRANOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFromEUTRANtoNGRANExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
