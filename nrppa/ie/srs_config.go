package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSConfig struct {
	SRSResourceList       *SRSResourceList                           // optional
	PosSRSResourceList    *PosSRSResourceList                        // optional
	SRSResourceSetList    *SRSResourceSetList                        // optional
	PosSRSResourceSetList *PosSRSResourceSetList                     // optional
	IEExtensions          *ProtocolExtensionContainerSRSConfigExtIEs // optional
}

func (x *SRSConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSConfigOptPresentFlag := []bool{}
	// optional field
	if x.SRSResourceList != nil {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, true)
	} else {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, false)
	}
	// optional field
	if x.PosSRSResourceList != nil {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, true)
	} else {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, false)
	}
	// optional field
	if x.SRSResourceSetList != nil {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, true)
	} else {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, false)
	}
	// optional field
	if x.PosSRSResourceSetList != nil {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, true)
	} else {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, true)
	} else {
		SRSConfigOptPresentFlag = append(SRSConfigOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSConfigOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.SRSResourceList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSResourceList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SRSResourceList marshal failed")
		}
	}

	// optional field
	if x.PosSRSResourceList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PosSRSResourceList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PosSRSResourceList marshal failed")
		}
	}

	// optional field
	if x.SRSResourceSetList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSResourceSetList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SRSResourceSetList marshal failed")
		}
	}

	// optional field
	if x.PosSRSResourceSetList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PosSRSResourceSetList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PosSRSResourceSetList marshal failed")
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

func (x *SRSConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSConfigOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&SRSConfigOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if SRSConfigOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SRSResourceList = new(SRSResourceList)
		err = x.SRSResourceList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SRSResourceList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SRSConfigOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PosSRSResourceList = new(PosSRSResourceList)
		err = x.PosSRSResourceList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PosSRSResourceList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if SRSConfigOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SRSResourceSetList = new(SRSResourceSetList)
		err = x.SRSResourceSetList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SRSResourceSetList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if SRSConfigOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.PosSRSResourceSetList = new(PosSRSResourceSetList)
		err = x.PosSRSResourceSetList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PosSRSResourceSetList error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if SRSConfigOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSConfigExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
