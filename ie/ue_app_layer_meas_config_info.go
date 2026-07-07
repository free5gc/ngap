package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	UEAppLayerMeasConfigInfoQoEMeasurementStatusPresentOngoing aper.Enumerated = 0
)

type UEAppLayerMeasConfigInfo struct {
	QoEReference                   *QoEReference
	ServiceType                    *ServiceType    // valueExt,valueLB:0,valueUB:2
	AreaScopeOfQMC                 *AreaScopeOfQMC // valueLB:0,valueUB:4,optional
	MeasCollEntityIPAddress        *TransportLayerAddress
	QoEMeasurementStatus           *aper.Enumerated                                          // valueExt,valueLB:0,valueUB:0,optional
	ContainerForAppLayerMeasConfig *aper.OctetString                                         // sizeLB:1,sizeUB:8000,optional
	MeasConfigAppLayerID           *int64                                                    // valueExt,valueLB:0,valueUB:15,optional
	SliceSupportListQMC            *SliceSupportListQMC                                      // optional
	MDTAlignmentInfo               *MDTAlignmentInfo                                         // valueLB:0,valueUB:1,optional
	AvailableRANVisibleQoEMetrics  *AvailableRANVisibleQoEMetrics                            // valueExt,optional
	IEExtensions                   *ProtocolExtensionContainerUEAppLayerMeasConfigInfoExtIEs // optional
}

func (x *UEAppLayerMeasConfigInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEAppLayerMeasConfigInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.QoEReference == nil {
		return errors.Errorf("QoEReference is missing")
	}
	// mandatory field
	if x.ServiceType == nil {
		return errors.Errorf("ServiceType is missing")
	}
	// optional field
	if x.AreaScopeOfQMC != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// mandatory field
	if x.MeasCollEntityIPAddress == nil {
		return errors.Errorf("MeasCollEntityIPAddress is missing")
	}
	// optional field
	if x.QoEMeasurementStatus != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.ContainerForAppLayerMeasConfig != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.MeasConfigAppLayerID != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.SliceSupportListQMC != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.MDTAlignmentInfo != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.AvailableRANVisibleQoEMetrics != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	} else {
		UEAppLayerMeasConfigInfoOptPresentFlag = append(UEAppLayerMeasConfigInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEAppLayerMeasConfigInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QoEReference.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QoEReference marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ServiceType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServiceType marshal failed")
	}

	// optional field
	if x.AreaScopeOfQMC != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaScopeOfQMC.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaScopeOfQMC marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MeasCollEntityIPAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MeasCollEntityIPAddress marshal failed")
	}

	// optional field
	if x.QoEMeasurementStatus != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.QoEMeasurementStatus), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.ContainerForAppLayerMeasConfig != nil {
		// Write OctetString (Pointer)
		*sLb, *sUb = 1, 8000
		err = pd.WriteOctetString(*(x.ContainerForAppLayerMeasConfig), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "octetString marshal failed")
		}
	}

	// optional field
	if x.MeasConfigAppLayerID != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 15
		err = pd.WriteInteger(*(x.MeasConfigAppLayerID), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.SliceSupportListQMC != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SliceSupportListQMC.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SliceSupportListQMC marshal failed")
		}
	}

	// optional field
	if x.MDTAlignmentInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MDTAlignmentInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MDTAlignmentInfo marshal failed")
		}
	}

	// optional field
	if x.AvailableRANVisibleQoEMetrics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AvailableRANVisibleQoEMetrics.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AvailableRANVisibleQoEMetrics marshal failed")
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

func (x *UEAppLayerMeasConfigInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEAppLayerMeasConfigInfoOptPresentFlag := make([]bool, 8)
	err = pd.ReadSequencePreambleBitMap(&UEAppLayerMeasConfigInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QoEReference = new(QoEReference)
	err = x.QoEReference.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QoEReference error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServiceType = new(ServiceType)
	err = x.ServiceType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServiceType error")
	}

	// optional field (optPresentFlag index: 0)
	if UEAppLayerMeasConfigInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AreaScopeOfQMC = new(AreaScopeOfQMC)
		err = x.AreaScopeOfQMC.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaScopeOfQMC error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MeasCollEntityIPAddress = new(TransportLayerAddress)
	err = x.MeasCollEntityIPAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MeasCollEntityIPAddress error")
	}

	// optional field (optPresentFlag index: 1)
	if UEAppLayerMeasConfigInfoOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.QoEMeasurementStatus = new(aper.Enumerated)
		*(x.QoEMeasurementStatus), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if UEAppLayerMeasConfigInfoOptPresentFlag[2] {
		// Read OctetString (Pointer)
		*sLb, *sUb = 1, 8000
		x.ContainerForAppLayerMeasConfig = new(aper.OctetString)
		*(x.ContainerForAppLayerMeasConfig), err = pd.ReadOctetString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if UEAppLayerMeasConfigInfoOptPresentFlag[3] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 15
		x.MeasConfigAppLayerID = new(int64)
		*(x.MeasConfigAppLayerID), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 4)
	if UEAppLayerMeasConfigInfoOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.SliceSupportListQMC = new(SliceSupportListQMC)
		err = x.SliceSupportListQMC.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SliceSupportListQMC error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if UEAppLayerMeasConfigInfoOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.MDTAlignmentInfo = new(MDTAlignmentInfo)
		err = x.MDTAlignmentInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MDTAlignmentInfo error")
		}
	}

	// optional field (optPresentFlag index: 6)
	if UEAppLayerMeasConfigInfoOptPresentFlag[6] {
		// Read struct defined elsewhere (Pointer)
		x.AvailableRANVisibleQoEMetrics = new(AvailableRANVisibleQoEMetrics)
		err = x.AvailableRANVisibleQoEMetrics.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AvailableRANVisibleQoEMetrics error")
		}
	}

	// optional field (optPresentFlag index: 7)
	if UEAppLayerMeasConfigInfoOptPresentFlag[7] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEAppLayerMeasConfigInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
