package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type NRPRSBeamInformationItem struct {
	PRSresourceSetID *PRSResourceSetID
	/* Sequence of = 35, FULL Name = struct NR_PRS_Beam_InformationItem__pRSAngle */
	/* Type Name = PRSAngleItem */
	/* Sequence Of Embed */
	PRSAngle     []PRSAngleItem                                            // valueExt,sizeLB:1,sizeUB:64
	IEExtensions *ProtocolExtensionContainerNRPRSBeamInformationItemExtIEs // optional
}

func (x *NRPRSBeamInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRPRSBeamInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSresourceSetID == nil {
		return errors.Errorf("PRSresourceSetID is missing")
	}
	// mandatory field
	if x.PRSAngle == nil {
		return errors.Errorf("PRSAngle is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NRPRSBeamInformationItemOptPresentFlag = append(NRPRSBeamInformationItemOptPresentFlag, true)
	} else {
		NRPRSBeamInformationItemOptPresentFlag = append(NRPRSBeamInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRPRSBeamInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSresourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSresourceSetID marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 64
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.PRSAngle)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.PRSAngle {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
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

func (x *NRPRSBeamInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRPRSBeamInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NRPRSBeamInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSresourceSetID = new(PRSResourceSetID)
	err = x.PRSresourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSresourceSetID error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsPRSAngle uint64
	numElementsPRSAngle, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.PRSAngle = []PRSAngleItem{}
	for i := 0; i < int(numElementsPRSAngle); i++ {
		var val PRSAngleItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.PRSAngle = append(x.PRSAngle, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if NRPRSBeamInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNRPRSBeamInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
