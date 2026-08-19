package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type NGRANCGI struct {
	PLMNIdentity *PLMNIdentity
	NGRANcell    *NGRANCell                                // valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerNGRANCGIExtIEs // optional
}

func (x *NGRANCGI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANCGIOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.NGRANcell == nil {
		return errors.Errorf("NGRANcell is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANCGIOptPresentFlag = append(NGRANCGIOptPresentFlag, true)
	} else {
		NGRANCGIOptPresentFlag = append(NGRANCGIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANCGIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANcell.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANcell marshal failed")
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

func (x *NGRANCGI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANCGIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANCGIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNIdentity = new(PLMNIdentity)
	err = x.PLMNIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNIdentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANcell = new(NGRANCell)
	err = x.NGRANcell.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANcell error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANCGIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANCGIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
