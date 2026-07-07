package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FromNGRANtoEUTRAN struct {
	SourceNGRANnodeID *IntersystemSONNGRANnodeID                         // valueExt
	TargeteNBID       *IntersystemSONeNBID                               // valueExt
	IEExtensions      *ProtocolExtensionContainerFromNGRANtoEUTRANExtIEs // optional
}

func (x *FromNGRANtoEUTRAN) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FromNGRANtoEUTRANOptPresentFlag := []bool{}
	// mandatory field
	if x.SourceNGRANnodeID == nil {
		return errors.Errorf("SourceNGRANnodeID is missing")
	}
	// mandatory field
	if x.TargeteNBID == nil {
		return errors.Errorf("TargeteNBID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		FromNGRANtoEUTRANOptPresentFlag = append(FromNGRANtoEUTRANOptPresentFlag, true)
	} else {
		FromNGRANtoEUTRANOptPresentFlag = append(FromNGRANtoEUTRANOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FromNGRANtoEUTRANOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SourceNGRANnodeID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SourceNGRANnodeID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TargeteNBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargeteNBID marshal failed")
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

func (x *FromNGRANtoEUTRAN) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FromNGRANtoEUTRANOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FromNGRANtoEUTRANOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourceNGRANnodeID = new(IntersystemSONNGRANnodeID)
	err = x.SourceNGRANnodeID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourceNGRANnodeID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargeteNBID = new(IntersystemSONeNBID)
	err = x.TargeteNBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargeteNBID error")
	}

	// optional field (optPresentFlag index: 0)
	if FromNGRANtoEUTRANOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFromNGRANtoEUTRANExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
