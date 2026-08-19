package ie

import (
	"sort"

	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ProtocolIEContainerECIDMeasurementInitiationRequestIEs struct {
	List []ECIDMeasurementInitiationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementInitiationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID:            0,
		ProtocolIEIDReportCharacteristics:         1,
		ProtocolIEIDMeasurementPeriodicity:        2,
		ProtocolIEIDMeasurementQuantities:         3,
		ProtocolIEIDOtherRATMeasurementQuantities: 4,
		ProtocolIEIDWLANMeasurementQuantities:     5,
		ProtocolIEIDMeasurementPeriodicityNRAoA:   6,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementInitiationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementInitiationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementInitiationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerECIDMeasurementInitiationResponseIEs struct {
	List []ECIDMeasurementInitiationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementInitiationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID:        0,
		ProtocolIEIDRANUEMeasurementID:        1,
		ProtocolIEIDECIDMeasurementResult:     2,
		ProtocolIEIDCriticalityDiagnostics:    3,
		ProtocolIEIDCellPortionID:             4,
		ProtocolIEIDOtherRATMeasurementResult: 5,
		ProtocolIEIDWLANMeasurementResult:     6,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementInitiationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementInitiationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementInitiationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerECIDMeasurementInitiationFailureIEs struct {
	List []ECIDMeasurementInitiationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementInitiationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID:     0,
		ProtocolIEIDCause:                  1,
		ProtocolIEIDCriticalityDiagnostics: 2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementInitiationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementInitiationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementInitiationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerECIDMeasurementFailureIndicationIEs struct {
	List []ECIDMeasurementFailureIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID: 0,
		ProtocolIEIDRANUEMeasurementID: 1,
		ProtocolIEIDCause:              2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementFailureIndicationIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementFailureIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerECIDMeasurementReportIEs struct {
	List []ECIDMeasurementReportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID:    0,
		ProtocolIEIDRANUEMeasurementID:    1,
		ProtocolIEIDECIDMeasurementResult: 2,
		ProtocolIEIDCellPortionID:         3,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementReportIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementReportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerECIDMeasurementTerminationCommandIEs struct {
	List []ECIDMeasurementTerminationCommandIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerECIDMeasurementTerminationCommandIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFUEMeasurementID: 0,
		ProtocolIEIDRANUEMeasurementID: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerECIDMeasurementTerminationCommandIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ECIDMeasurementTerminationCommandIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ECIDMeasurementTerminationCommandIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerOTDOAInformationRequestIEs struct {
	List []OTDOAInformationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerOTDOAInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDOTDOAInformationTypeGroup: 0,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerOTDOAInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []OTDOAInformationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val OTDOAInformationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerOTDOAInformationResponseIEs struct {
	List []OTDOAInformationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerOTDOAInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDOTDOACells:             0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerOTDOAInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []OTDOAInformationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val OTDOAInformationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerOTDOAInformationFailureIEs struct {
	List []OTDOAInformationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerOTDOAInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerOTDOAInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []OTDOAInformationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val OTDOAInformationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAssistanceInformationControlIEs struct {
	List []AssistanceInformationControlIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAssistanceInformationControlIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAssistanceInformation:     0,
		ProtocolIEIDBroadcast:                 1,
		ProtocolIEIDPositioningBroadcastCells: 2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerAssistanceInformationControlIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AssistanceInformationControlIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val AssistanceInformationControlIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAssistanceInformationFeedbackIEs struct {
	List []AssistanceInformationFeedbackIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAssistanceInformationFeedbackIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAssistanceInformationFailureList: 0,
		ProtocolIEIDPositioningBroadcastCells:        1,
		ProtocolIEIDCriticalityDiagnostics:           2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerAssistanceInformationFeedbackIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AssistanceInformationFeedbackIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val AssistanceInformationFeedbackIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerErrorIndicationIEs struct {
	List []ErrorIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerErrorIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerErrorIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ErrorIndicationIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val ErrorIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningInformationRequestIEs struct {
	List []PositioningInformationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRequestedSRSTransmissionCharacteristics: 0,
		ProtocolIEIDUEReportingInformation:                  1,
		ProtocolIEIDUETEGInfoRequest:                        2,
		ProtocolIEIDUETEGReportingPeriodicity:               3,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningInformationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningInformationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningInformationResponseIEs struct {
	List []PositioningInformationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDSRSConfiguration:       0,
		ProtocolIEIDSFNInitialisationTime:  1,
		ProtocolIEIDCriticalityDiagnostics: 2,
		ProtocolIEIDUETxTEGAssociationList: 3,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningInformationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningInformationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningInformationFailureIEs struct {
	List []PositioningInformationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningInformationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningInformationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningInformationUpdateIEs struct {
	List []PositioningInformationUpdateIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningInformationUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDSRSConfiguration:       0,
		ProtocolIEIDSFNInitialisationTime:  1,
		ProtocolIEIDUETxTEGAssociationList: 2,
		ProtocolIEIDSRSTransmissionStatus:  3,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningInformationUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningInformationUpdateIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningInformationUpdateIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementRequestIEs struct {
	List []MeasurementRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID:                           0,
		ProtocolIEIDTRPMeasurementRequestList:                  1,
		ProtocolIEIDReportCharacteristics:                      2,
		ProtocolIEIDMeasurementPeriodicity:                     3,
		ProtocolIEIDTRPMeasurementQuantities:                   4,
		ProtocolIEIDSFNInitialisationTime:                      5,
		ProtocolIEIDSRSConfiguration:                           6,
		ProtocolIEIDMeasurementBeamInfoRequest:                 7,
		ProtocolIEIDSystemFrameNumber:                          8,
		ProtocolIEIDSlotNumber:                                 9,
		ProtocolIEIDMeasurementPeriodicityExtended:             10,
		ProtocolIEIDResponseTime:                               11,
		ProtocolIEIDMeasurementCharacteristicsRequestIndicator: 12,
		ProtocolIEIDMeasurementTimeOccasion:                    13,
		ProtocolIEIDMeasurementAmount:                          14,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementResponseIEs struct {
	List []MeasurementResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID:           0,
		ProtocolIEIDRANMeasurementID:           1,
		ProtocolIEIDTRPMeasurementResponseList: 2,
		ProtocolIEIDCriticalityDiagnostics:     3,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementFailureIEs struct {
	List []MeasurementFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID:       0,
		ProtocolIEIDCause:                  1,
		ProtocolIEIDCriticalityDiagnostics: 2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementReportIEs struct {
	List []MeasurementReportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID:         0,
		ProtocolIEIDRANMeasurementID:         1,
		ProtocolIEIDTRPMeasurementReportList: 2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementReportIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementReportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementUpdateIEs struct {
	List []MeasurementUpdateIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID:                           0,
		ProtocolIEIDRANMeasurementID:                           1,
		ProtocolIEIDSRSConfiguration:                           2,
		ProtocolIEIDTRPMeasurementUpdateList:                   3,
		ProtocolIEIDMeasurementCharacteristicsRequestIndicator: 4,
		ProtocolIEIDMeasurementTimeOccasion:                    5,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementUpdateIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementUpdateIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementAbortIEs struct {
	List []MeasurementAbortIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementAbortIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID: 0,
		ProtocolIEIDRANMeasurementID: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementAbortIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementAbortIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementAbortIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementFailureIndicationIEs struct {
	List []MeasurementFailureIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDLMFMeasurementID: 0,
		ProtocolIEIDRANMeasurementID: 1,
		ProtocolIEIDCause:            2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementFailureIndicationIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementFailureIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerTRPInformationRequestIEs struct {
	List []TRPInformationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerTRPInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDTRPList:                      0,
		ProtocolIEIDTRPInformationTypeListTRPReq: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerTRPInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []TRPInformationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val TRPInformationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerTRPInformationResponseIEs struct {
	List []TRPInformationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerTRPInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDTRPInformationListTRPResp: 0,
		ProtocolIEIDCriticalityDiagnostics:    1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerTRPInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []TRPInformationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val TRPInformationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerTRPInformationFailureIEs struct {
	List []TRPInformationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerTRPInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerTRPInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []TRPInformationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val TRPInformationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningActivationRequestIEs struct {
	List []PositioningActivationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningActivationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDSRSType:        0,
		ProtocolIEIDActivationTime: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningActivationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningActivationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningActivationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningActivationResponseIEs struct {
	List []PositioningActivationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningActivationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCriticalityDiagnostics: 0,
		ProtocolIEIDSystemFrameNumber:      1,
		ProtocolIEIDSlotNumber:             2,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningActivationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningActivationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningActivationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningActivationFailureIEs struct {
	List []PositioningActivationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningActivationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningActivationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningActivationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningActivationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPositioningDeactivationIEs struct {
	List []PositioningDeactivationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPositioningDeactivationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAbortTransmission: 0,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPositioningDeactivationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PositioningDeactivationIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PositioningDeactivationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPRSConfigurationRequestIEs struct {
	List []PRSConfigurationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPRSConfigurationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPRSConfigRequestType: 0,
		ProtocolIEIDPRSTRPList:           1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPRSConfigurationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PRSConfigurationRequestIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PRSConfigurationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPRSConfigurationResponseIEs struct {
	List []PRSConfigurationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPRSConfigurationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPRSTransmissionTRPList: 0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPRSConfigurationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PRSConfigurationResponseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PRSConfigurationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPRSConfigurationFailureIEs struct {
	List []PRSConfigurationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPRSConfigurationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerPRSConfigurationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PRSConfigurationFailureIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val PRSConfigurationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementPreconfigurationRequiredIEs struct {
	List []MeasurementPreconfigurationRequiredIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementPreconfigurationRequiredIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDTRPPRSInformationList: 0,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementPreconfigurationRequiredIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementPreconfigurationRequiredIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementPreconfigurationRequiredIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementPreconfigurationConfirmIEs struct {
	List []MeasurementPreconfigurationConfirmIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementPreconfigurationConfirmIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPreconfigurationResult: 0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementPreconfigurationConfirmIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementPreconfigurationConfirmIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementPreconfigurationConfirmIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementPreconfigurationRefuseIEs struct {
	List []MeasurementPreconfigurationRefuseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementPreconfigurationRefuseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDCriticalityDiagnostics: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementPreconfigurationRefuseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementPreconfigurationRefuseIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementPreconfigurationRefuseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMeasurementActivationIEs struct {
	List []MeasurementActivationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMeasurementActivationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRequestType:             0,
		ProtocolIEIDPRSMeasurementsInfoList: 1,
	}
	sort.Slice(x.List, func(i, j int) bool {
		if x.List[i].Id() == nil || x.List[j].Id() == nil {
			return false
		}
		// Notes: the order is not guaruanteed if id cant be find in order map
		return order[x.List[i].Id().Value] < order[x.List[j].Id().Value]
	})

	// Write Sequence Of
	*sLb, *sUb = 0, 65535
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *ProtocolIEContainerMeasurementActivationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MeasurementActivationIEs{}
	for i := 0; i < int(numElementsList); i++ {
		var val MeasurementActivationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
