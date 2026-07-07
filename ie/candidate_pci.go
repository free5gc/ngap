package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CandidatePCI struct {
	CandidatePCI     *int64                                        // valueExt,valueLB:0,valueUB:1007
	CandidateNRARFCN *int64                                        // valueLB:0,valueUB:3279165
	IEExtensions     *ProtocolExtensionContainerCandidatePCIExtIEs // optional
}

func (x *CandidatePCI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CandidatePCIOptPresentFlag := []bool{}
	// mandatory field
	if x.CandidatePCI == nil {
		return errors.Errorf("CandidatePCI is missing")
	}
	// mandatory field
	if x.CandidateNRARFCN == nil {
		return errors.Errorf("CandidateNRARFCN is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CandidatePCIOptPresentFlag = append(CandidatePCIOptPresentFlag, true)
	} else {
		CandidatePCIOptPresentFlag = append(CandidatePCIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CandidatePCIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1007
	err = pd.WriteInteger(*(x.CandidatePCI), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.CandidateNRARFCN), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *CandidatePCI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CandidatePCIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CandidatePCIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1007
	x.CandidatePCI = new(int64)
	*(x.CandidatePCI), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.CandidateNRARFCN = new(int64)
	*(x.CandidateNRARFCN), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if CandidatePCIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCandidatePCIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
