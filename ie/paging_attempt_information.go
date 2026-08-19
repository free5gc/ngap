package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PagingAttemptInformation struct {
	PagingAttemptCount             *PagingAttemptCount
	IntendedNumberOfPagingAttempts *IntendedNumberOfPagingAttempts
	NextPagingAreaScope            *NextPagingAreaScope                                      // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions                   *ProtocolExtensionContainerPagingAttemptInformationExtIEs // optional
}

func (x *PagingAttemptInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PagingAttemptInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.PagingAttemptCount == nil {
		return errors.Errorf("PagingAttemptCount is missing")
	}
	// mandatory field
	if x.IntendedNumberOfPagingAttempts == nil {
		return errors.Errorf("IntendedNumberOfPagingAttempts is missing")
	}
	// optional field
	if x.NextPagingAreaScope != nil {
		PagingAttemptInformationOptPresentFlag = append(PagingAttemptInformationOptPresentFlag, true)
	} else {
		PagingAttemptInformationOptPresentFlag = append(PagingAttemptInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PagingAttemptInformationOptPresentFlag = append(PagingAttemptInformationOptPresentFlag, true)
	} else {
		PagingAttemptInformationOptPresentFlag = append(PagingAttemptInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PagingAttemptInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PagingAttemptCount.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PagingAttemptCount marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IntendedNumberOfPagingAttempts.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntendedNumberOfPagingAttempts marshal failed")
	}

	// optional field
	if x.NextPagingAreaScope != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NextPagingAreaScope.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NextPagingAreaScope marshal failed")
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

func (x *PagingAttemptInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PagingAttemptInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PagingAttemptInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PagingAttemptCount = new(PagingAttemptCount)
	err = x.PagingAttemptCount.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PagingAttemptCount error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntendedNumberOfPagingAttempts = new(IntendedNumberOfPagingAttempts)
	err = x.IntendedNumberOfPagingAttempts.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntendedNumberOfPagingAttempts error")
	}

	// optional field (optPresentFlag index: 0)
	if PagingAttemptInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NextPagingAreaScope = new(NextPagingAreaScope)
		err = x.NextPagingAreaScope.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NextPagingAreaScope error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PagingAttemptInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPagingAttemptInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
