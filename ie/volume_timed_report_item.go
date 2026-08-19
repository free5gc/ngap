package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type VolumeTimedReportItem struct {
	StartTimeStamp *aper.OctetString                                      // sizeLB:4,sizeUB:4
	EndTimeStamp   *aper.OctetString                                      // sizeLB:4,sizeUB:4
	UsageCountUL   *int64                                                 // valueLB:0,valueUB:18446744073709551615
	UsageCountDL   *int64                                                 // valueLB:0,valueUB:18446744073709551615
	IEExtensions   *ProtocolExtensionContainerVolumeTimedReportItemExtIEs // optional
}

func (x *VolumeTimedReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	VolumeTimedReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.StartTimeStamp == nil {
		return errors.Errorf("StartTimeStamp is missing")
	}
	// mandatory field
	if x.EndTimeStamp == nil {
		return errors.Errorf("EndTimeStamp is missing")
	}
	// mandatory field
	if x.UsageCountUL == nil {
		return errors.Errorf("UsageCountUL is missing")
	}
	// mandatory field
	if x.UsageCountDL == nil {
		return errors.Errorf("UsageCountDL is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		VolumeTimedReportItemOptPresentFlag = append(VolumeTimedReportItemOptPresentFlag, true)
	} else {
		VolumeTimedReportItemOptPresentFlag = append(VolumeTimedReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(VolumeTimedReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write OctetString (Pointer)
	*sLb, *sUb = 4, 4
	err = pd.WriteOctetString(*(x.StartTimeStamp), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
	}

	// Write OctetString (Pointer)
	*sLb, *sUb = 4, 4
	err = pd.WriteOctetString(*(x.EndTimeStamp), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
	}

	// Write Integer
	var uint64Lb, uint64Ub *uint64
	*uint64Lb, *uint64Ub = 0, 18446744073709551615
	err = pd.WriteUint64Integer(uint64(*(x.UsageCountUL)), false, uint64Lb, uint64Ub)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer
	*uint64Lb, *uint64Ub = 0, 18446744073709551615
	err = pd.WriteUint64Integer(uint64(*(x.UsageCountDL)), false, uint64Lb, uint64Ub)
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

func (x *VolumeTimedReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	VolumeTimedReportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&VolumeTimedReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read OctetString (Pointer)
	*sLb, *sUb = 4, 4
	x.StartTimeStamp = new(aper.OctetString)
	*(x.StartTimeStamp), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// mandatory field
	// Read OctetString (Pointer)
	*sLb, *sUb = 4, 4
	x.EndTimeStamp = new(aper.OctetString)
	*(x.EndTimeStamp), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// mandatory field
	// Read Integer
	var uint64Lb, uint64Ub *uint64
	var uint64Val uint64
	*uint64Lb, *uint64Ub = 0, 18446744073709551615
	uint64Val, err = pd.ReadUint64Integer(false, uint64Lb, uint64Ub)
	if err != nil {
		return errors.Wrap(err, "integer unmarshal failed")
	}
	*(x.UsageCountUL) = int64(uint64Val)

	// mandatory field
	// Read Integer
	*uint64Lb, *uint64Ub = 0, 18446744073709551615
	uint64Val, err = pd.ReadUint64Integer(false, uint64Lb, uint64Ub)
	if err != nil {
		return errors.Wrap(err, "integer unmarshal failed")
	}
	*(x.UsageCountDL) = int64(uint64Val)

	// optional field (optPresentFlag index: 0)
	if VolumeTimedReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerVolumeTimedReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
