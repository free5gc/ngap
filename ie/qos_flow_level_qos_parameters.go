package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowLevelQosParameters struct {
	QosCharacteristics             *QosCharacteristics                                        // valueLB:0,valueUB:2
	AllocationAndRetentionPriority *AllocationAndRetentionPriority                            // valueExt
	GBRQosInformation              *GBRQosInformation                                         // valueExt,optional
	ReflectiveQosAttribute         *ReflectiveQosAttribute                                    // valueExt,valueLB:0,valueUB:0,optional
	AdditionalQosFlowInformation   *AdditionalQosFlowInformation                              // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions                   *ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs // optional
}

func (x *QosFlowLevelQosParameters) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowLevelQosParametersOptPresentFlag := []bool{}
	// mandatory field
	if x.QosCharacteristics == nil {
		return errors.Errorf("QosCharacteristics is missing")
	}
	// mandatory field
	if x.AllocationAndRetentionPriority == nil {
		return errors.Errorf("AllocationAndRetentionPriority is missing")
	}
	// optional field
	if x.GBRQosInformation != nil {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, true)
	} else {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, false)
	}
	// optional field
	if x.ReflectiveQosAttribute != nil {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, true)
	} else {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, false)
	}
	// optional field
	if x.AdditionalQosFlowInformation != nil {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, true)
	} else {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, true)
	} else {
		QosFlowLevelQosParametersOptPresentFlag = append(QosFlowLevelQosParametersOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowLevelQosParametersOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosCharacteristics.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosCharacteristics marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AllocationAndRetentionPriority.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AllocationAndRetentionPriority marshal failed")
	}

	// optional field
	if x.GBRQosInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.GBRQosInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "GBRQosInformation marshal failed")
		}
	}

	// optional field
	if x.ReflectiveQosAttribute != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ReflectiveQosAttribute.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ReflectiveQosAttribute marshal failed")
		}
	}

	// optional field
	if x.AdditionalQosFlowInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalQosFlowInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalQosFlowInformation marshal failed")
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

func (x *QosFlowLevelQosParameters) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowLevelQosParametersOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&QosFlowLevelQosParametersOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosCharacteristics = new(QosCharacteristics)
	err = x.QosCharacteristics.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosCharacteristics error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AllocationAndRetentionPriority = new(AllocationAndRetentionPriority)
	err = x.AllocationAndRetentionPriority.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AllocationAndRetentionPriority error")
	}

	// optional field (optPresentFlag index: 0)
	if QosFlowLevelQosParametersOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.GBRQosInformation = new(GBRQosInformation)
		err = x.GBRQosInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode GBRQosInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowLevelQosParametersOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ReflectiveQosAttribute = new(ReflectiveQosAttribute)
		err = x.ReflectiveQosAttribute.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ReflectiveQosAttribute error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if QosFlowLevelQosParametersOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalQosFlowInformation = new(AdditionalQosFlowInformation)
		err = x.AdditionalQosFlowInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalQosFlowInformation error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if QosFlowLevelQosParametersOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
