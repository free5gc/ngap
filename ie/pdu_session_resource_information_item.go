package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceInformationItem struct {
	PDUSessionID              *PDUSessionID
	QosFlowInformationList    *QosFlowInformationList
	DRBsToQosFlowsMappingList *DRBsToQosFlowsMappingList                                         // optional
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs // optional
}

func (x *PDUSessionResourceInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// mandatory field
	if x.QosFlowInformationList == nil {
		return errors.Errorf("QosFlowInformationList is missing")
	}
	// optional field
	if x.DRBsToQosFlowsMappingList != nil {
		PDUSessionResourceInformationItemOptPresentFlag = append(PDUSessionResourceInformationItemOptPresentFlag, true)
	} else {
		PDUSessionResourceInformationItemOptPresentFlag = append(PDUSessionResourceInformationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceInformationItemOptPresentFlag = append(PDUSessionResourceInformationItemOptPresentFlag, true)
	} else {
		PDUSessionResourceInformationItemOptPresentFlag = append(PDUSessionResourceInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PDUSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PDUSessionID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowInformationList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowInformationList marshal failed")
	}

	// optional field
	if x.DRBsToQosFlowsMappingList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DRBsToQosFlowsMappingList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DRBsToQosFlowsMappingList marshal failed")
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

func (x *PDUSessionResourceInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceInformationItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PDUSessionID = new(PDUSessionID)
	err = x.PDUSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PDUSessionID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowInformationList = new(QosFlowInformationList)
	err = x.QosFlowInformationList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowInformationList error")
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DRBsToQosFlowsMappingList = new(DRBsToQosFlowsMappingList)
		err = x.DRBsToQosFlowsMappingList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DRBsToQosFlowsMappingList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceInformationItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
