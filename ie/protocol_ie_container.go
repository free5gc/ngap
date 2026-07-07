package ie

import (
	"sort"

	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ProtocolIEContainerPDUSessionResourceSetupRequestIEs struct {
	List []PDUSessionResourceSetupRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceSetupRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                      0,
		ProtocolIEIDRANUENGAPID:                      1,
		ProtocolIEIDRANPagingPriority:                2,
		ProtocolIEIDNASPDU:                           3,
		ProtocolIEIDPDUSessionResourceSetupListSUReq: 4,
		ProtocolIEIDUEAggregateMaximumBitRate:        5,
		ProtocolIEIDUESliceMaximumBitRateList:        6,
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

func (x *ProtocolIEContainerPDUSessionResourceSetupRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceSetupRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceSetupRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceSetupResponseIEs struct {
	List []PDUSessionResourceSetupResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceSetupResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                              0,
		ProtocolIEIDRANUENGAPID:                              1,
		ProtocolIEIDPDUSessionResourceSetupListSURes:         2,
		ProtocolIEIDPDUSessionResourceFailedToSetupListSURes: 3,
		ProtocolIEIDCriticalityDiagnostics:                   4,
		ProtocolIEIDUserLocationInformation:                  5,
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

func (x *ProtocolIEContainerPDUSessionResourceSetupResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceSetupResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceSetupResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceReleaseCommandIEs struct {
	List []PDUSessionResourceReleaseCommandIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceReleaseCommandIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                           0,
		ProtocolIEIDRANUENGAPID:                           1,
		ProtocolIEIDRANPagingPriority:                     2,
		ProtocolIEIDNASPDU:                                3,
		ProtocolIEIDPDUSessionResourceToReleaseListRelCmd: 4,
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

func (x *ProtocolIEContainerPDUSessionResourceReleaseCommandIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceReleaseCommandIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceReleaseCommandIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceReleaseResponseIEs struct {
	List []PDUSessionResourceReleaseResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceReleaseResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                          0,
		ProtocolIEIDRANUENGAPID:                          1,
		ProtocolIEIDPDUSessionResourceReleasedListRelRes: 2,
		ProtocolIEIDUserLocationInformation:              3,
		ProtocolIEIDCriticalityDiagnostics:               4,
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

func (x *ProtocolIEContainerPDUSessionResourceReleaseResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceReleaseResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceReleaseResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceModifyRequestIEs struct {
	List []PDUSessionResourceModifyRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceModifyRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                        0,
		ProtocolIEIDRANUENGAPID:                        1,
		ProtocolIEIDRANPagingPriority:                  2,
		ProtocolIEIDPDUSessionResourceModifyListModReq: 3,
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

func (x *ProtocolIEContainerPDUSessionResourceModifyRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceModifyRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceModifyRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceModifyResponseIEs struct {
	List []PDUSessionResourceModifyResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceModifyResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDPDUSessionResourceModifyListModRes:         2,
		ProtocolIEIDPDUSessionResourceFailedToModifyListModRes: 3,
		ProtocolIEIDUserLocationInformation:                    4,
		ProtocolIEIDCriticalityDiagnostics:                     5,
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

func (x *ProtocolIEContainerPDUSessionResourceModifyResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceModifyResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceModifyResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceNotifyIEs struct {
	List []PDUSessionResourceNotifyIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceNotifyIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                       0,
		ProtocolIEIDRANUENGAPID:                       1,
		ProtocolIEIDPDUSessionResourceNotifyList:      2,
		ProtocolIEIDPDUSessionResourceReleasedListNot: 3,
		ProtocolIEIDUserLocationInformation:           4,
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

func (x *ProtocolIEContainerPDUSessionResourceNotifyIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceNotifyIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceNotifyIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceModifyIndicationIEs struct {
	List []PDUSessionResourceModifyIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceModifyIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                        0,
		ProtocolIEIDRANUENGAPID:                        1,
		ProtocolIEIDPDUSessionResourceModifyListModInd: 2,
		ProtocolIEIDUserLocationInformation:            3,
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

func (x *ProtocolIEContainerPDUSessionResourceModifyIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceModifyIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceModifyIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceModifyConfirmIEs struct {
	List []PDUSessionResourceModifyConfirmIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceModifyConfirmIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDPDUSessionResourceModifyListModCfm:         2,
		ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm: 3,
		ProtocolIEIDCriticalityDiagnostics:                     4,
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

func (x *ProtocolIEContainerPDUSessionResourceModifyConfirmIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceModifyConfirmIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceModifyConfirmIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerInitialContextSetupRequestIEs struct {
	List []InitialContextSetupRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerInitialContextSetupRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                 0,
		ProtocolIEIDRANUENGAPID:                                 1,
		ProtocolIEIDOldAMF:                                      2,
		ProtocolIEIDUEAggregateMaximumBitRate:                   3,
		ProtocolIEIDCoreNetworkAssistanceInformationForInactive: 4,
		ProtocolIEIDGUAMI:                                       5,
		ProtocolIEIDPDUSessionResourceSetupListCxtReq:           6,
		ProtocolIEIDAllowedNSSAI:                                7,
		ProtocolIEIDUESecurityCapabilities:                      8,
		ProtocolIEIDSecurityKey:                                 9,
		ProtocolIEIDTraceActivation:                             10,
		ProtocolIEIDMobilityRestrictionList:                     11,
		ProtocolIEIDUERadioCapability:                           12,
		ProtocolIEIDIndexToRFSP:                                 13,
		ProtocolIEIDMaskedIMEISV:                                14,
		ProtocolIEIDNASPDU:                                      15,
		ProtocolIEIDEmergencyFallbackIndicator:                  16,
		ProtocolIEIDRRCInactiveTransitionReportRequest:          17,
		ProtocolIEIDUERadioCapabilityForPaging:                  18,
		ProtocolIEIDRedirectionVoiceFallback:                    19,
		ProtocolIEIDLocationReportingRequestType:                20,
		ProtocolIEIDCNAssistedRANTuning:                         21,
		ProtocolIEIDSRVCCOperationPossible:                      22,
		ProtocolIEIDIABAuthorized:                               23,
		ProtocolIEIDEnhancedCoverageRestriction:                 24,
		ProtocolIEIDExtendedConnectedTime:                       25,
		ProtocolIEIDUEDifferentiationInfo:                       26,
		ProtocolIEIDNRV2XServicesAuthorized:                     27,
		ProtocolIEIDLTEV2XServicesAuthorized:                    28,
		ProtocolIEIDNRUESidelinkAggregateMaximumBitrate:         29,
		ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate:        30,
		ProtocolIEIDPC5QoSParameters:                            31,
		ProtocolIEIDCEmodeBrestricted:                           32,
		ProtocolIEIDUEUPCIoTSupport:                             33,
		ProtocolIEIDRGLevelWirelineAccessCharacteristics:        34,
		ProtocolIEIDManagementBasedMDTPLMNList:                  35,
		ProtocolIEIDUERadioCapabilityID:                         36,
		ProtocolIEIDTimeSyncAssistanceInfo:                      37,
		ProtocolIEIDQMCConfigInfo:                               38,
		ProtocolIEIDTargetNSSAIInformation:                      39,
		ProtocolIEIDUESliceMaximumBitRateList:                   40,
		ProtocolIEIDFiveGProSeAuthorized:                        41,
		ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate:      42,
		ProtocolIEIDFiveGProSePC5QoSParameters:                  43,
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

func (x *ProtocolIEContainerInitialContextSetupRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []InitialContextSetupRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val InitialContextSetupRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerInitialContextSetupResponseIEs struct {
	List []InitialContextSetupResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerInitialContextSetupResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                               0,
		ProtocolIEIDRANUENGAPID:                               1,
		ProtocolIEIDPDUSessionResourceSetupListCxtRes:         2,
		ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes: 3,
		ProtocolIEIDCriticalityDiagnostics:                    4,
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

func (x *ProtocolIEContainerInitialContextSetupResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []InitialContextSetupResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val InitialContextSetupResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerInitialContextSetupFailureIEs struct {
	List []InitialContextSetupFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerInitialContextSetupFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail: 2,
		ProtocolIEIDCause:                                      3,
		ProtocolIEIDCriticalityDiagnostics:                     4,
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

func (x *ProtocolIEContainerInitialContextSetupFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []InitialContextSetupFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val InitialContextSetupFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextReleaseRequestIEs struct {
	List []UEContextReleaseRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextReleaseRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                     0,
		ProtocolIEIDRANUENGAPID:                     1,
		ProtocolIEIDPDUSessionResourceListCxtRelReq: 2,
		ProtocolIEIDCause:                           3,
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

func (x *ProtocolIEContainerUEContextReleaseRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextReleaseRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextReleaseRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextReleaseCommandIEs struct {
	List []UEContextReleaseCommandIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextReleaseCommandIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUENGAPIDs: 0,
		ProtocolIEIDCause:     1,
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

func (x *ProtocolIEContainerUEContextReleaseCommandIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextReleaseCommandIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextReleaseCommandIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextReleaseCompleteIEs struct {
	List []UEContextReleaseCompleteIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextReleaseCompleteIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDUserLocationInformation:                    2,
		ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging: 3,
		ProtocolIEIDPDUSessionResourceListCxtRelCpl:            4,
		ProtocolIEIDCriticalityDiagnostics:                     5,
		ProtocolIEIDPagingAssisDataforCEcapabUE:                6,
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

func (x *ProtocolIEContainerUEContextReleaseCompleteIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextReleaseCompleteIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextReleaseCompleteIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextResumeRequestIEs struct {
	List []UEContextResumeRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextResumeRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDRRCResumeCause:                             2,
		ProtocolIEIDPDUSessionResourceResumeListRESReq:         3,
		ProtocolIEIDPDUSessionResourceFailedToResumeListRESReq: 4,
		ProtocolIEIDSuspendRequestIndication:                   5,
		ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging: 6,
		ProtocolIEIDPagingAssisDataforCEcapabUE:                7,
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

func (x *ProtocolIEContainerUEContextResumeRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextResumeRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextResumeRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextResumeResponseIEs struct {
	List []UEContextResumeResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextResumeResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDPDUSessionResourceResumeListRESRes:         2,
		ProtocolIEIDPDUSessionResourceFailedToResumeListRESRes: 3,
		ProtocolIEIDSecurityContext:                            4,
		ProtocolIEIDSuspendResponseIndication:                  5,
		ProtocolIEIDExtendedConnectedTime:                      6,
		ProtocolIEIDCriticalityDiagnostics:                     7,
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

func (x *ProtocolIEContainerUEContextResumeResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextResumeResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextResumeResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextResumeFailureIEs struct {
	List []UEContextResumeFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextResumeFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
		ProtocolIEIDCause:                  2,
		ProtocolIEIDCriticalityDiagnostics: 3,
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

func (x *ProtocolIEContainerUEContextResumeFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextResumeFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextResumeFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextSuspendRequestIEs struct {
	List []UEContextSuspendRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextSuspendRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                0,
		ProtocolIEIDRANUENGAPID:                                1,
		ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging: 2,
		ProtocolIEIDPagingAssisDataforCEcapabUE:                3,
		ProtocolIEIDPDUSessionResourceSuspendListSUSReq:        4,
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

func (x *ProtocolIEContainerUEContextSuspendRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextSuspendRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextSuspendRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextSuspendResponseIEs struct {
	List []UEContextSuspendResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextSuspendResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
		ProtocolIEIDSecurityContext:        2,
		ProtocolIEIDCriticalityDiagnostics: 3,
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

func (x *ProtocolIEContainerUEContextSuspendResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextSuspendResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextSuspendResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextSuspendFailureIEs struct {
	List []UEContextSuspendFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextSuspendFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
		ProtocolIEIDCause:                  2,
		ProtocolIEIDCriticalityDiagnostics: 3,
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

func (x *ProtocolIEContainerUEContextSuspendFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextSuspendFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextSuspendFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextModificationRequestIEs struct {
	List []UEContextModificationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextModificationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                 0,
		ProtocolIEIDRANUENGAPID:                                 1,
		ProtocolIEIDRANPagingPriority:                           2,
		ProtocolIEIDSecurityKey:                                 3,
		ProtocolIEIDIndexToRFSP:                                 4,
		ProtocolIEIDUEAggregateMaximumBitRate:                   5,
		ProtocolIEIDUESecurityCapabilities:                      6,
		ProtocolIEIDCoreNetworkAssistanceInformationForInactive: 7,
		ProtocolIEIDEmergencyFallbackIndicator:                  8,
		ProtocolIEIDNewAMFUENGAPID:                              9,
		ProtocolIEIDRRCInactiveTransitionReportRequest:          10,
		ProtocolIEIDNewGUAMI:                                    11,
		ProtocolIEIDCNAssistedRANTuning:                         12,
		ProtocolIEIDSRVCCOperationPossible:                      13,
		ProtocolIEIDIABAuthorized:                               14,
		ProtocolIEIDNRV2XServicesAuthorized:                     15,
		ProtocolIEIDLTEV2XServicesAuthorized:                    16,
		ProtocolIEIDNRUESidelinkAggregateMaximumBitrate:         17,
		ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate:        18,
		ProtocolIEIDPC5QoSParameters:                            19,
		ProtocolIEIDUERadioCapabilityID:                         20,
		ProtocolIEIDRGLevelWirelineAccessCharacteristics:        21,
		ProtocolIEIDTimeSyncAssistanceInfo:                      22,
		ProtocolIEIDQMCConfigInfo:                               23,
		ProtocolIEIDQMCDeactivation:                             24,
		ProtocolIEIDUESliceMaximumBitRateList:                   25,
		ProtocolIEIDManagementBasedMDTPLMNModificationList:      26,
		ProtocolIEIDFiveGProSeAuthorized:                        27,
		ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate:      28,
		ProtocolIEIDFiveGProSePC5QoSParameters:                  29,
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

func (x *ProtocolIEContainerUEContextModificationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextModificationRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextModificationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextModificationResponseIEs struct {
	List []UEContextModificationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextModificationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:             0,
		ProtocolIEIDRANUENGAPID:             1,
		ProtocolIEIDRRCState:                2,
		ProtocolIEIDUserLocationInformation: 3,
		ProtocolIEIDCriticalityDiagnostics:  4,
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

func (x *ProtocolIEContainerUEContextModificationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextModificationResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextModificationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEContextModificationFailureIEs struct {
	List []UEContextModificationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEContextModificationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
		ProtocolIEIDCause:                  2,
		ProtocolIEIDCriticalityDiagnostics: 3,
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

func (x *ProtocolIEContainerUEContextModificationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEContextModificationFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEContextModificationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRRCInactiveTransitionReportIEs struct {
	List []RRCInactiveTransitionReportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRRCInactiveTransitionReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:             0,
		ProtocolIEIDRANUENGAPID:             1,
		ProtocolIEIDRRCState:                2,
		ProtocolIEIDUserLocationInformation: 3,
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

func (x *ProtocolIEContainerRRCInactiveTransitionReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RRCInactiveTransitionReportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RRCInactiveTransitionReportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRetrieveUEInformationIEs struct {
	List []RetrieveUEInformationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRetrieveUEInformationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDFiveGSTMSI: 0,
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

func (x *ProtocolIEContainerRetrieveUEInformationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RetrieveUEInformationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RetrieveUEInformationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUEInformationTransferIEs struct {
	List []UEInformationTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUEInformationTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDFiveGSTMSI:            0,
		ProtocolIEIDNBIoTUEPriority:       1,
		ProtocolIEIDUERadioCapability:     2,
		ProtocolIEIDSNSSAI:                3,
		ProtocolIEIDAllowedNSSAI:          4,
		ProtocolIEIDUEDifferentiationInfo: 5,
		ProtocolIEIDMaskedIMEISV:          6,
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

func (x *ProtocolIEContainerUEInformationTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UEInformationTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UEInformationTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRANCPRelocationIndicationIEs struct {
	List []RANCPRelocationIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRANCPRelocationIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRANUENGAPID:             0,
		ProtocolIEIDFiveGSTMSI:              1,
		ProtocolIEIDEUTRACGI:                2,
		ProtocolIEIDTAI:                     3,
		ProtocolIEIDULCPSecurityInformation: 4,
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

func (x *ProtocolIEContainerRANCPRelocationIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RANCPRelocationIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RANCPRelocationIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverRequiredIEs struct {
	List []HandoverRequiredIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverRequiredIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                        0,
		ProtocolIEIDRANUENGAPID:                        1,
		ProtocolIEIDHandoverType:                       2,
		ProtocolIEIDCause:                              3,
		ProtocolIEIDTargetID:                           4,
		ProtocolIEIDDirectForwardingPathAvailability:   5,
		ProtocolIEIDPDUSessionResourceListHORqd:        6,
		ProtocolIEIDSourceToTargetTransparentContainer: 7,
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

func (x *ProtocolIEContainerHandoverRequiredIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverRequiredIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverRequiredIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverCommandIEs struct {
	List []HandoverCommandIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverCommandIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                          0,
		ProtocolIEIDRANUENGAPID:                          1,
		ProtocolIEIDHandoverType:                         2,
		ProtocolIEIDNASSecurityParametersFromNGRAN:       3,
		ProtocolIEIDPDUSessionResourceHandoverList:       4,
		ProtocolIEIDPDUSessionResourceToReleaseListHOCmd: 5,
		ProtocolIEIDTargetToSourceTransparentContainer:   6,
		ProtocolIEIDCriticalityDiagnostics:               7,
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

func (x *ProtocolIEContainerHandoverCommandIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverCommandIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverCommandIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverPreparationFailureIEs struct {
	List []HandoverPreparationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverPreparationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                               0,
		ProtocolIEIDRANUENGAPID:                               1,
		ProtocolIEIDCause:                                     2,
		ProtocolIEIDCriticalityDiagnostics:                    3,
		ProtocolIEIDTargettoSourceFailureTransparentContainer: 4,
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

func (x *ProtocolIEContainerHandoverPreparationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverPreparationFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverPreparationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverRequestIEs struct {
	List []HandoverRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                 0,
		ProtocolIEIDHandoverType:                                1,
		ProtocolIEIDCause:                                       2,
		ProtocolIEIDUEAggregateMaximumBitRate:                   3,
		ProtocolIEIDCoreNetworkAssistanceInformationForInactive: 4,
		ProtocolIEIDUESecurityCapabilities:                      5,
		ProtocolIEIDSecurityContext:                             6,
		ProtocolIEIDNewSecurityContextInd:                       7,
		ProtocolIEIDNASC:                                        8,
		ProtocolIEIDPDUSessionResourceSetupListHOReq:            9,
		ProtocolIEIDAllowedNSSAI:                                10,
		ProtocolIEIDTraceActivation:                             11,
		ProtocolIEIDMaskedIMEISV:                                12,
		ProtocolIEIDSourceToTargetTransparentContainer:          13,
		ProtocolIEIDMobilityRestrictionList:                     14,
		ProtocolIEIDLocationReportingRequestType:                15,
		ProtocolIEIDRRCInactiveTransitionReportRequest:          16,
		ProtocolIEIDGUAMI:                                       17,
		ProtocolIEIDRedirectionVoiceFallback:                    18,
		ProtocolIEIDCNAssistedRANTuning:                         19,
		ProtocolIEIDSRVCCOperationPossible:                      20,
		ProtocolIEIDIABAuthorized:                               21,
		ProtocolIEIDEnhancedCoverageRestriction:                 22,
		ProtocolIEIDUEDifferentiationInfo:                       23,
		ProtocolIEIDNRV2XServicesAuthorized:                     24,
		ProtocolIEIDLTEV2XServicesAuthorized:                    25,
		ProtocolIEIDNRUESidelinkAggregateMaximumBitrate:         26,
		ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate:        27,
		ProtocolIEIDPC5QoSParameters:                            28,
		ProtocolIEIDCEmodeBrestricted:                           29,
		ProtocolIEIDUEUPCIoTSupport:                             30,
		ProtocolIEIDManagementBasedMDTPLMNList:                  31,
		ProtocolIEIDUERadioCapabilityID:                         32,
		ProtocolIEIDExtendedConnectedTime:                       33,
		ProtocolIEIDTimeSyncAssistanceInfo:                      34,
		ProtocolIEIDUESliceMaximumBitRateList:                   35,
		ProtocolIEIDFiveGProSeAuthorized:                        36,
		ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate:      37,
		ProtocolIEIDFiveGProSePC5QoSParameters:                  38,
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

func (x *ProtocolIEContainerHandoverRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverRequestAcknowledgeIEs struct {
	List []HandoverRequestAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverRequestAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                              0,
		ProtocolIEIDRANUENGAPID:                              1,
		ProtocolIEIDPDUSessionResourceAdmittedList:           2,
		ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck: 3,
		ProtocolIEIDTargetToSourceTransparentContainer:       4,
		ProtocolIEIDCriticalityDiagnostics:                   5,
		ProtocolIEIDNPNAccessInformation:                     6,
		ProtocolIEIDRedCapIndication:                         7,
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

func (x *ProtocolIEContainerHandoverRequestAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverRequestAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverRequestAcknowledgeIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverFailureIEs struct {
	List []HandoverFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                               0,
		ProtocolIEIDCause:                                     1,
		ProtocolIEIDCriticalityDiagnostics:                    2,
		ProtocolIEIDTargettoSourceFailureTransparentContainer: 3,
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

func (x *ProtocolIEContainerHandoverFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverNotifyIEs struct {
	List []HandoverNotifyIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverNotifyIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:             0,
		ProtocolIEIDRANUENGAPID:             1,
		ProtocolIEIDUserLocationInformation: 2,
		ProtocolIEIDNotifySourceNGRANNode:   3,
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

func (x *ProtocolIEContainerHandoverNotifyIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverNotifyIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverNotifyIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPathSwitchRequestIEs struct {
	List []PathSwitchRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPathSwitchRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRANUENGAPID:                              0,
		ProtocolIEIDSourceAMFUENGAPID:                        1,
		ProtocolIEIDUserLocationInformation:                  2,
		ProtocolIEIDUESecurityCapabilities:                   3,
		ProtocolIEIDPDUSessionResourceToBeSwitchedDLList:     4,
		ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq: 5,
		ProtocolIEIDRRCResumeCause:                           6,
		ProtocolIEIDRedCapIndication:                         7,
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

func (x *ProtocolIEContainerPathSwitchRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PathSwitchRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PathSwitchRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPathSwitchRequestAcknowledgeIEs struct {
	List []PathSwitchRequestAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPathSwitchRequestAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                                 0,
		ProtocolIEIDRANUENGAPID:                                 1,
		ProtocolIEIDUESecurityCapabilities:                      2,
		ProtocolIEIDSecurityContext:                             3,
		ProtocolIEIDNewSecurityContextInd:                       4,
		ProtocolIEIDPDUSessionResourceSwitchedList:              5,
		ProtocolIEIDPDUSessionResourceReleasedListPSAck:         6,
		ProtocolIEIDAllowedNSSAI:                                7,
		ProtocolIEIDCoreNetworkAssistanceInformationForInactive: 8,
		ProtocolIEIDRRCInactiveTransitionReportRequest:          9,
		ProtocolIEIDCriticalityDiagnostics:                      10,
		ProtocolIEIDRedirectionVoiceFallback:                    11,
		ProtocolIEIDCNAssistedRANTuning:                         12,
		ProtocolIEIDSRVCCOperationPossible:                      13,
		ProtocolIEIDEnhancedCoverageRestriction:                 14,
		ProtocolIEIDExtendedConnectedTime:                       15,
		ProtocolIEIDUEDifferentiationInfo:                       16,
		ProtocolIEIDNRV2XServicesAuthorized:                     17,
		ProtocolIEIDLTEV2XServicesAuthorized:                    18,
		ProtocolIEIDNRUESidelinkAggregateMaximumBitrate:         19,
		ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate:        20,
		ProtocolIEIDPC5QoSParameters:                            21,
		ProtocolIEIDCEmodeBrestricted:                           22,
		ProtocolIEIDUEUPCIoTSupport:                             23,
		ProtocolIEIDUERadioCapabilityID:                         24,
		ProtocolIEIDManagementBasedMDTPLMNList:                  25,
		ProtocolIEIDTimeSyncAssistanceInfo:                      26,
		ProtocolIEIDFiveGProSeAuthorized:                        27,
		ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate:      28,
		ProtocolIEIDFiveGProSePC5QoSParameters:                  29,
		ProtocolIEIDManagementBasedMDTPLMNModificationList:      30,
		ProtocolIEIDIABAuthorized:                               31,
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

func (x *ProtocolIEContainerPathSwitchRequestAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PathSwitchRequestAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PathSwitchRequestAcknowledgeIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPathSwitchRequestFailureIEs struct {
	List []PathSwitchRequestFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPathSwitchRequestFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                          0,
		ProtocolIEIDRANUENGAPID:                          1,
		ProtocolIEIDPDUSessionResourceReleasedListPSFail: 2,
		ProtocolIEIDCriticalityDiagnostics:               3,
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

func (x *ProtocolIEContainerPathSwitchRequestFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PathSwitchRequestFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PathSwitchRequestFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverCancelIEs struct {
	List []HandoverCancelIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverCancelIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
		ProtocolIEIDCause:       2,
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

func (x *ProtocolIEContainerHandoverCancelIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverCancelIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverCancelIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverCancelAcknowledgeIEs struct {
	List []HandoverCancelAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverCancelAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
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

func (x *ProtocolIEContainerHandoverCancelAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverCancelAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverCancelAcknowledgeIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerHandoverSuccessIEs struct {
	List []HandoverSuccessIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerHandoverSuccessIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
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

func (x *ProtocolIEContainerHandoverSuccessIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []HandoverSuccessIEs{}
	for i := 0; i < int(numElements); i++ {
		var val HandoverSuccessIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkRANEarlyStatusTransferIEs struct {
	List []UplinkRANEarlyStatusTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkRANEarlyStatusTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                             0,
		ProtocolIEIDRANUENGAPID:                             1,
		ProtocolIEIDEarlyStatusTransferTransparentContainer: 2,
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

func (x *ProtocolIEContainerUplinkRANEarlyStatusTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkRANEarlyStatusTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkRANEarlyStatusTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkRANEarlyStatusTransferIEs struct {
	List []DownlinkRANEarlyStatusTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkRANEarlyStatusTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                             0,
		ProtocolIEIDRANUENGAPID:                             1,
		ProtocolIEIDEarlyStatusTransferTransparentContainer: 2,
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

func (x *ProtocolIEContainerDownlinkRANEarlyStatusTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkRANEarlyStatusTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkRANEarlyStatusTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkRANStatusTransferIEs struct {
	List []UplinkRANStatusTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkRANStatusTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                           0,
		ProtocolIEIDRANUENGAPID:                           1,
		ProtocolIEIDRANStatusTransferTransparentContainer: 2,
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

func (x *ProtocolIEContainerUplinkRANStatusTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkRANStatusTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkRANStatusTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkRANStatusTransferIEs struct {
	List []DownlinkRANStatusTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkRANStatusTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                           0,
		ProtocolIEIDRANUENGAPID:                           1,
		ProtocolIEIDRANStatusTransferTransparentContainer: 2,
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

func (x *ProtocolIEContainerDownlinkRANStatusTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkRANStatusTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkRANStatusTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPagingIEs struct {
	List []PagingIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPagingIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUEPagingIdentity:            0,
		ProtocolIEIDPagingDRX:                   1,
		ProtocolIEIDTAIListForPaging:            2,
		ProtocolIEIDPagingPriority:              3,
		ProtocolIEIDUERadioCapabilityForPaging:  4,
		ProtocolIEIDPagingOrigin:                5,
		ProtocolIEIDAssistanceDataForPaging:     6,
		ProtocolIEIDNBIoTPagingEDRXInfo:         7,
		ProtocolIEIDNBIoTPagingDRX:              8,
		ProtocolIEIDEnhancedCoverageRestriction: 9,
		ProtocolIEIDWUSAssistanceInformation:    10,
		ProtocolIEIDEUTRAPagingeDRXInformation:  11,
		ProtocolIEIDCEmodeBrestricted:           12,
		ProtocolIEIDNRPagingeDRXInformation:     13,
		ProtocolIEIDPagingCause:                 14,
		ProtocolIEIDPEIPSassistanceInformation:  15,
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

func (x *ProtocolIEContainerPagingIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PagingIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PagingIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerInitialUEMessageIEs struct {
	List []InitialUEMessageIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerInitialUEMessageIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRANUENGAPID:                         0,
		ProtocolIEIDNASPDU:                              1,
		ProtocolIEIDUserLocationInformation:             2,
		ProtocolIEIDRRCEstablishmentCause:               3,
		ProtocolIEIDFiveGSTMSI:                          4,
		ProtocolIEIDAMFSetID:                            5,
		ProtocolIEIDUEContextRequest:                    6,
		ProtocolIEIDAllowedNSSAI:                        7,
		ProtocolIEIDSourceToTargetAMFInformationReroute: 8,
		ProtocolIEIDSelectedPLMNIdentity:                9,
		ProtocolIEIDIABNodeIndication:                   10,
		ProtocolIEIDCEmodeBSupportIndicator:             11,
		ProtocolIEIDLTEMIndication:                      12,
		ProtocolIEIDEDTSession:                          13,
		ProtocolIEIDAuthenticatedIndication:             14,
		ProtocolIEIDNPNAccessInformation:                15,
		ProtocolIEIDRedCapIndication:                    16,
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

func (x *ProtocolIEContainerInitialUEMessageIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []InitialUEMessageIEs{}
	for i := 0; i < int(numElements); i++ {
		var val InitialUEMessageIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkNASTransportIEs struct {
	List []DownlinkNASTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkNASTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                 0,
		ProtocolIEIDRANUENGAPID:                 1,
		ProtocolIEIDOldAMF:                      2,
		ProtocolIEIDRANPagingPriority:           3,
		ProtocolIEIDNASPDU:                      4,
		ProtocolIEIDMobilityRestrictionList:     5,
		ProtocolIEIDIndexToRFSP:                 6,
		ProtocolIEIDUEAggregateMaximumBitRate:   7,
		ProtocolIEIDAllowedNSSAI:                8,
		ProtocolIEIDSRVCCOperationPossible:      9,
		ProtocolIEIDEnhancedCoverageRestriction: 10,
		ProtocolIEIDExtendedConnectedTime:       11,
		ProtocolIEIDUEDifferentiationInfo:       12,
		ProtocolIEIDCEmodeBrestricted:           13,
		ProtocolIEIDUERadioCapability:           14,
		ProtocolIEIDUECapabilityInfoRequest:     15,
		ProtocolIEIDEndIndication:               16,
		ProtocolIEIDUERadioCapabilityID:         17,
		ProtocolIEIDTargetNSSAIInformation:      18,
		ProtocolIEIDMaskedIMEISV:                19,
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

func (x *ProtocolIEContainerDownlinkNASTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkNASTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkNASTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkNASTransportIEs struct {
	List []UplinkNASTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkNASTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:             0,
		ProtocolIEIDRANUENGAPID:             1,
		ProtocolIEIDNASPDU:                  2,
		ProtocolIEIDUserLocationInformation: 3,
		ProtocolIEIDWAGFIdentityInformation: 4,
		ProtocolIEIDTNGFIdentityInformation: 5,
		ProtocolIEIDTWIFIdentityInformation: 6,
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

func (x *ProtocolIEContainerUplinkNASTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkNASTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkNASTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNASNonDeliveryIndicationIEs struct {
	List []NASNonDeliveryIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNASNonDeliveryIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
		ProtocolIEIDNASPDU:      2,
		ProtocolIEIDCause:       3,
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

func (x *ProtocolIEContainerNASNonDeliveryIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NASNonDeliveryIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NASNonDeliveryIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRerouteNASRequestIEs struct {
	List []RerouteNASRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRerouteNASRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRANUENGAPID:                         0,
		ProtocolIEIDAMFUENGAPID:                         1,
		ProtocolIEIDNGAPMessage:                         2,
		ProtocolIEIDAMFSetID:                            3,
		ProtocolIEIDAllowedNSSAI:                        4,
		ProtocolIEIDSourceToTargetAMFInformationReroute: 5,
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

func (x *ProtocolIEContainerRerouteNASRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RerouteNASRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RerouteNASRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNGSetupRequestIEs struct {
	List []NGSetupRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNGSetupRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDGlobalRANNodeID:        0,
		ProtocolIEIDRANNodeName:            1,
		ProtocolIEIDSupportedTAList:        2,
		ProtocolIEIDDefaultPagingDRX:       3,
		ProtocolIEIDUERetentionInformation: 4,
		ProtocolIEIDNBIoTDefaultPagingDRX:  5,
		ProtocolIEIDExtendedRANNodeName:    6,
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

func (x *ProtocolIEContainerNGSetupRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NGSetupRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NGSetupRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNGSetupResponseIEs struct {
	List []NGSetupResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNGSetupResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFName:                0,
		ProtocolIEIDServedGUAMIList:        1,
		ProtocolIEIDRelativeAMFCapacity:    2,
		ProtocolIEIDPLMNSupportList:        3,
		ProtocolIEIDCriticalityDiagnostics: 4,
		ProtocolIEIDUERetentionInformation: 5,
		ProtocolIEIDIABSupported:           6,
		ProtocolIEIDExtendedAMFName:        7,
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

func (x *ProtocolIEContainerNGSetupResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NGSetupResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NGSetupResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNGSetupFailureIEs struct {
	List []NGSetupFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNGSetupFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDTimeToWait:             1,
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

func (x *ProtocolIEContainerNGSetupFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NGSetupFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NGSetupFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRANConfigurationUpdateIEs struct {
	List []RANConfigurationUpdateIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRANConfigurationUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRANNodeName:                     0,
		ProtocolIEIDSupportedTAList:                 1,
		ProtocolIEIDDefaultPagingDRX:                2,
		ProtocolIEIDGlobalRANNodeID:                 3,
		ProtocolIEIDNGRANTNLAssociationToRemoveList: 4,
		ProtocolIEIDNBIoTDefaultPagingDRX:           5,
		ProtocolIEIDExtendedRANNodeName:             6,
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

func (x *ProtocolIEContainerRANConfigurationUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RANConfigurationUpdateIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RANConfigurationUpdateIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs struct {
	List []RANConfigurationUpdateAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCriticalityDiagnostics: 0,
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

func (x *ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RANConfigurationUpdateAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RANConfigurationUpdateAcknowledgeIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerRANConfigurationUpdateFailureIEs struct {
	List []RANConfigurationUpdateFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerRANConfigurationUpdateFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDTimeToWait:             1,
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

func (x *ProtocolIEContainerRANConfigurationUpdateFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []RANConfigurationUpdateFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val RANConfigurationUpdateFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAMFConfigurationUpdateIEs struct {
	List []AMFConfigurationUpdateIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAMFConfigurationUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFName:                       0,
		ProtocolIEIDServedGUAMIList:               1,
		ProtocolIEIDRelativeAMFCapacity:           2,
		ProtocolIEIDPLMNSupportList:               3,
		ProtocolIEIDAMFTNLAssociationToAddList:    4,
		ProtocolIEIDAMFTNLAssociationToRemoveList: 5,
		ProtocolIEIDAMFTNLAssociationToUpdateList: 6,
		ProtocolIEIDExtendedAMFName:               7,
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

func (x *ProtocolIEContainerAMFConfigurationUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AMFConfigurationUpdateIEs{}
	for i := 0; i < int(numElements); i++ {
		var val AMFConfigurationUpdateIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs struct {
	List []AMFConfigurationUpdateAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFTNLAssociationSetupList:         0,
		ProtocolIEIDAMFTNLAssociationFailedToSetupList: 1,
		ProtocolIEIDCriticalityDiagnostics:             2,
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

func (x *ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AMFConfigurationUpdateAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val AMFConfigurationUpdateAcknowledgeIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAMFConfigurationUpdateFailureIEs struct {
	List []AMFConfigurationUpdateFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAMFConfigurationUpdateFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:                  0,
		ProtocolIEIDTimeToWait:             1,
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

func (x *ProtocolIEContainerAMFConfigurationUpdateFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AMFConfigurationUpdateFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val AMFConfigurationUpdateFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAMFStatusIndicationIEs struct {
	List []AMFStatusIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAMFStatusIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUnavailableGUAMIList: 0,
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

func (x *ProtocolIEContainerAMFStatusIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AMFStatusIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val AMFStatusIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNGResetIEs struct {
	List []NGResetIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNGResetIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCause:     0,
		ProtocolIEIDResetType: 1,
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

func (x *ProtocolIEContainerNGResetIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NGResetIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NGResetIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerNGResetAcknowledgeIEs struct {
	List []NGResetAcknowledgeIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerNGResetAcknowledgeIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUEAssociatedLogicalNGConnectionList: 0,
		ProtocolIEIDCriticalityDiagnostics:              1,
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

func (x *ProtocolIEContainerNGResetAcknowledgeIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []NGResetAcknowledgeIEs{}
	for i := 0; i < int(numElements); i++ {
		var val NGResetAcknowledgeIEs
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
		ProtocolIEIDAMFUENGAPID:            0,
		ProtocolIEIDRANUENGAPID:            1,
		ProtocolIEIDCause:                  2,
		ProtocolIEIDCriticalityDiagnostics: 3,
		ProtocolIEIDFiveGSTMSI:             4,
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
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ErrorIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val ErrorIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerOverloadStartIEs struct {
	List []OverloadStartIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerOverloadStartIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFOverloadResponse:               0,
		ProtocolIEIDAMFTrafficLoadReductionIndication: 1,
		ProtocolIEIDOverloadStartNSSAIList:            2,
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

func (x *ProtocolIEContainerOverloadStartIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []OverloadStartIEs{}
	for i := 0; i < int(numElements); i++ {
		var val OverloadStartIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerOverloadStopIEs struct {
	List []OverloadStopIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerOverloadStopIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{}
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

func (x *ProtocolIEContainerOverloadStopIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []OverloadStopIEs{}
	for i := 0; i < int(numElements); i++ {
		var val OverloadStopIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkRANConfigurationTransferIEs struct {
	List []UplinkRANConfigurationTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkRANConfigurationTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDSONConfigurationTransferUL:            0,
		ProtocolIEIDENDCSONConfigurationTransferUL:        1,
		ProtocolIEIDIntersystemSONConfigurationTransferUL: 2,
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

func (x *ProtocolIEContainerUplinkRANConfigurationTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkRANConfigurationTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkRANConfigurationTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkRANConfigurationTransferIEs struct {
	List []DownlinkRANConfigurationTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkRANConfigurationTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDSONConfigurationTransferDL:            0,
		ProtocolIEIDENDCSONConfigurationTransferDL:        1,
		ProtocolIEIDIntersystemSONConfigurationTransferDL: 2,
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

func (x *ProtocolIEContainerDownlinkRANConfigurationTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkRANConfigurationTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkRANConfigurationTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerWriteReplaceWarningRequestIEs struct {
	List []WriteReplaceWarningRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerWriteReplaceWarningRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMessageIdentifier:           0,
		ProtocolIEIDSerialNumber:                1,
		ProtocolIEIDWarningAreaList:             2,
		ProtocolIEIDRepetitionPeriod:            3,
		ProtocolIEIDNumberOfBroadcastsRequested: 4,
		ProtocolIEIDWarningType:                 5,
		ProtocolIEIDWarningSecurityInfo:         6,
		ProtocolIEIDDataCodingScheme:            7,
		ProtocolIEIDWarningMessageContents:      8,
		ProtocolIEIDConcurrentWarningMessageInd: 9,
		ProtocolIEIDWarningAreaCoordinates:      10,
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

func (x *ProtocolIEContainerWriteReplaceWarningRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []WriteReplaceWarningRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val WriteReplaceWarningRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerWriteReplaceWarningResponseIEs struct {
	List []WriteReplaceWarningResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerWriteReplaceWarningResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMessageIdentifier:          0,
		ProtocolIEIDSerialNumber:               1,
		ProtocolIEIDBroadcastCompletedAreaList: 2,
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

func (x *ProtocolIEContainerWriteReplaceWarningResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []WriteReplaceWarningResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val WriteReplaceWarningResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPWSCancelRequestIEs struct {
	List []PWSCancelRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPWSCancelRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMessageIdentifier:        0,
		ProtocolIEIDSerialNumber:             1,
		ProtocolIEIDWarningAreaList:          2,
		ProtocolIEIDCancelAllWarningMessages: 3,
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

func (x *ProtocolIEContainerPWSCancelRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PWSCancelRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PWSCancelRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPWSCancelResponseIEs struct {
	List []PWSCancelResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPWSCancelResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMessageIdentifier:          0,
		ProtocolIEIDSerialNumber:               1,
		ProtocolIEIDBroadcastCancelledAreaList: 2,
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

func (x *ProtocolIEContainerPWSCancelResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PWSCancelResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PWSCancelResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPWSRestartIndicationIEs struct {
	List []PWSRestartIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPWSRestartIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDCellIDListForRestart:          0,
		ProtocolIEIDGlobalRANNodeID:               1,
		ProtocolIEIDTAIListForRestart:             2,
		ProtocolIEIDEmergencyAreaIDListForRestart: 3,
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

func (x *ProtocolIEContainerPWSRestartIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PWSRestartIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PWSRestartIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPWSFailureIndicationIEs struct {
	List []PWSFailureIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPWSFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPWSFailedCellIDList: 0,
		ProtocolIEIDGlobalRANNodeID:     1,
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

func (x *ProtocolIEContainerPWSFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PWSFailureIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PWSFailureIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkUEAssociatedNRPPaTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
		ProtocolIEIDRoutingID:   2,
		ProtocolIEIDNRPPaPDU:    3,
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

func (x *ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkUEAssociatedNRPPaTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkUEAssociatedNRPPaTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs struct {
	List []UplinkUEAssociatedNRPPaTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
		ProtocolIEIDRoutingID:   2,
		ProtocolIEIDNRPPaPDU:    3,
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

func (x *ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkUEAssociatedNRPPaTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkUEAssociatedNRPPaTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkNonUEAssociatedNRPPaTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRoutingID: 0,
		ProtocolIEIDNRPPaPDU:  1,
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

func (x *ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkNonUEAssociatedNRPPaTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkNonUEAssociatedNRPPaTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs struct {
	List []UplinkNonUEAssociatedNRPPaTransportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRoutingID: 0,
		ProtocolIEIDNRPPaPDU:  1,
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

func (x *ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkNonUEAssociatedNRPPaTransportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkNonUEAssociatedNRPPaTransportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerTraceStartIEs struct {
	List []TraceStartIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerTraceStartIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:     0,
		ProtocolIEIDRANUENGAPID:     1,
		ProtocolIEIDTraceActivation: 2,
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

func (x *ProtocolIEContainerTraceStartIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []TraceStartIEs{}
	for i := 0; i < int(numElements); i++ {
		var val TraceStartIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerTraceFailureIndicationIEs struct {
	List []TraceFailureIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerTraceFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:  0,
		ProtocolIEIDRANUENGAPID:  1,
		ProtocolIEIDNGRANTraceID: 2,
		ProtocolIEIDCause:        3,
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

func (x *ProtocolIEContainerTraceFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []TraceFailureIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val TraceFailureIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDeactivateTraceIEs struct {
	List []DeactivateTraceIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDeactivateTraceIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:  0,
		ProtocolIEIDRANUENGAPID:  1,
		ProtocolIEIDNGRANTraceID: 2,
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

func (x *ProtocolIEContainerDeactivateTraceIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DeactivateTraceIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DeactivateTraceIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerCellTrafficTraceIEs struct {
	List []CellTrafficTraceIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerCellTrafficTraceIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                    0,
		ProtocolIEIDRANUENGAPID:                    1,
		ProtocolIEIDNGRANTraceID:                   2,
		ProtocolIEIDNGRANCGI:                       3,
		ProtocolIEIDTraceCollectionEntityIPAddress: 4,
		ProtocolIEIDPrivacyIndicator:               5,
		ProtocolIEIDTraceCollectionEntityURI:       6,
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

func (x *ProtocolIEContainerCellTrafficTraceIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []CellTrafficTraceIEs{}
	for i := 0; i < int(numElements); i++ {
		var val CellTrafficTraceIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerLocationReportingControlIEs struct {
	List []LocationReportingControlIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerLocationReportingControlIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                  0,
		ProtocolIEIDRANUENGAPID:                  1,
		ProtocolIEIDLocationReportingRequestType: 2,
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

func (x *ProtocolIEContainerLocationReportingControlIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []LocationReportingControlIEs{}
	for i := 0; i < int(numElements); i++ {
		var val LocationReportingControlIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerLocationReportingFailureIndicationIEs struct {
	List []LocationReportingFailureIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerLocationReportingFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
		ProtocolIEIDCause:       2,
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

func (x *ProtocolIEContainerLocationReportingFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []LocationReportingFailureIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val LocationReportingFailureIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerLocationReportIEs struct {
	List []LocationReportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerLocationReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                    0,
		ProtocolIEIDRANUENGAPID:                    1,
		ProtocolIEIDUserLocationInformation:        2,
		ProtocolIEIDUEPresenceInAreaOfInterestList: 3,
		ProtocolIEIDLocationReportingRequestType:   4,
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

func (x *ProtocolIEContainerLocationReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []LocationReportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val LocationReportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUETNLABindingReleaseRequestIEs struct {
	List []UETNLABindingReleaseRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUETNLABindingReleaseRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID: 0,
		ProtocolIEIDRANUENGAPID: 1,
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

func (x *ProtocolIEContainerUETNLABindingReleaseRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UETNLABindingReleaseRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UETNLABindingReleaseRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUERadioCapabilityInfoIndicationIEs struct {
	List []UERadioCapabilityInfoIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUERadioCapabilityInfoIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                  0,
		ProtocolIEIDRANUENGAPID:                  1,
		ProtocolIEIDUERadioCapability:            2,
		ProtocolIEIDUERadioCapabilityForPaging:   3,
		ProtocolIEIDUERadioCapabilityEUTRAFormat: 4,
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

func (x *ProtocolIEContainerUERadioCapabilityInfoIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UERadioCapabilityInfoIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UERadioCapabilityInfoIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUERadioCapabilityCheckRequestIEs struct {
	List []UERadioCapabilityCheckRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUERadioCapabilityCheckRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:         0,
		ProtocolIEIDRANUENGAPID:         1,
		ProtocolIEIDUERadioCapability:   2,
		ProtocolIEIDUERadioCapabilityID: 3,
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

func (x *ProtocolIEContainerUERadioCapabilityCheckRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UERadioCapabilityCheckRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UERadioCapabilityCheckRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUERadioCapabilityCheckResponseIEs struct {
	List []UERadioCapabilityCheckResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUERadioCapabilityCheckResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:              0,
		ProtocolIEIDRANUENGAPID:              1,
		ProtocolIEIDIMSVoiceSupportIndicator: 2,
		ProtocolIEIDCriticalityDiagnostics:   3,
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

func (x *ProtocolIEContainerUERadioCapabilityCheckResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UERadioCapabilityCheckResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UERadioCapabilityCheckResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerSecondaryRATDataUsageReportIEs struct {
	List []SecondaryRATDataUsageReportIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerSecondaryRATDataUsageReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                             0,
		ProtocolIEIDRANUENGAPID:                             1,
		ProtocolIEIDPDUSessionResourceSecondaryRATUsageList: 2,
		ProtocolIEIDHandoverFlag:                            3,
		ProtocolIEIDUserLocationInformation:                 4,
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

func (x *ProtocolIEContainerSecondaryRATDataUsageReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []SecondaryRATDataUsageReportIEs{}
	for i := 0; i < int(numElements); i++ {
		var val SecondaryRATDataUsageReportIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUplinkRIMInformationTransferIEs struct {
	List []UplinkRIMInformationTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUplinkRIMInformationTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRIMInformationTransfer: 0,
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

func (x *ProtocolIEContainerUplinkRIMInformationTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UplinkRIMInformationTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UplinkRIMInformationTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDownlinkRIMInformationTransferIEs struct {
	List []DownlinkRIMInformationTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDownlinkRIMInformationTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDRIMInformationTransfer: 0,
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

func (x *ProtocolIEContainerDownlinkRIMInformationTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DownlinkRIMInformationTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DownlinkRIMInformationTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerConnectionEstablishmentIndicationIEs struct {
	List []ConnectionEstablishmentIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerConnectionEstablishmentIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:                 0,
		ProtocolIEIDRANUENGAPID:                 1,
		ProtocolIEIDUERadioCapability:           2,
		ProtocolIEIDEndIndication:               3,
		ProtocolIEIDSNSSAI:                      4,
		ProtocolIEIDAllowedNSSAI:                5,
		ProtocolIEIDUEDifferentiationInfo:       6,
		ProtocolIEIDDLCPSecurityInformation:     7,
		ProtocolIEIDNBIoTUEPriority:             8,
		ProtocolIEIDEnhancedCoverageRestriction: 9,
		ProtocolIEIDCEmodeBrestricted:           10,
		ProtocolIEIDUERadioCapabilityID:         11,
		ProtocolIEIDMaskedIMEISV:                12,
		ProtocolIEIDOldAMF:                      13,
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

func (x *ProtocolIEContainerConnectionEstablishmentIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []ConnectionEstablishmentIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val ConnectionEstablishmentIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUERadioCapabilityIDMappingRequestIEs struct {
	List []UERadioCapabilityIDMappingRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUERadioCapabilityIDMappingRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUERadioCapabilityID: 0,
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

func (x *ProtocolIEContainerUERadioCapabilityIDMappingRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UERadioCapabilityIDMappingRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UERadioCapabilityIDMappingRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerUERadioCapabilityIDMappingResponseIEs struct {
	List []UERadioCapabilityIDMappingResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerUERadioCapabilityIDMappingResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDUERadioCapabilityID:    0,
		ProtocolIEIDUERadioCapability:      1,
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

func (x *ProtocolIEContainerUERadioCapabilityIDMappingResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []UERadioCapabilityIDMappingResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val UERadioCapabilityIDMappingResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerAMFCPRelocationIndicationIEs struct {
	List []AMFCPRelocationIndicationIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerAMFCPRelocationIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDAMFUENGAPID:  0,
		ProtocolIEIDRANUENGAPID:  1,
		ProtocolIEIDSNSSAI:       2,
		ProtocolIEIDAllowedNSSAI: 3,
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

func (x *ProtocolIEContainerAMFCPRelocationIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AMFCPRelocationIndicationIEs{}
	for i := 0; i < int(numElements); i++ {
		var val AMFCPRelocationIndicationIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionSetupRequestIEs struct {
	List []BroadcastSessionSetupRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionSetupRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:   0,
		ProtocolIEIDSNSSAI:         1,
		ProtocolIEIDMBSServiceArea: 2,
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

func (x *ProtocolIEContainerBroadcastSessionSetupRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionSetupRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionSetupRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionSetupResponseIEs struct {
	List []BroadcastSessionSetupResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionSetupResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerBroadcastSessionSetupResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionSetupResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionSetupResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionSetupFailureIEs struct {
	List []BroadcastSessionSetupFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionSetupFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerBroadcastSessionSetupFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionSetupFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionSetupFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionModificationRequestIEs struct {
	List []BroadcastSessionModificationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionModificationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:   0,
		ProtocolIEIDMBSServiceArea: 1,
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

func (x *ProtocolIEContainerBroadcastSessionModificationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionModificationRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionModificationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionModificationResponseIEs struct {
	List []BroadcastSessionModificationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionModificationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerBroadcastSessionModificationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionModificationResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionModificationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionModificationFailureIEs struct {
	List []BroadcastSessionModificationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionModificationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerBroadcastSessionModificationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionModificationFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionModificationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionReleaseRequestIEs struct {
	List []BroadcastSessionReleaseRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionReleaseRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
		ProtocolIEIDCause:        1,
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

func (x *ProtocolIEContainerBroadcastSessionReleaseRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionReleaseRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionReleaseRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionReleaseRequiredIEs struct {
	List []BroadcastSessionReleaseRequiredIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionReleaseRequiredIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
		ProtocolIEIDCause:        1,
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

func (x *ProtocolIEContainerBroadcastSessionReleaseRequiredIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionReleaseRequiredIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionReleaseRequiredIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerBroadcastSessionReleaseResponseIEs struct {
	List []BroadcastSessionReleaseResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerBroadcastSessionReleaseResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerBroadcastSessionReleaseResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []BroadcastSessionReleaseResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val BroadcastSessionReleaseResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDistributionSetupRequestIEs struct {
	List []DistributionSetupRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDistributionSetupRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:     0,
		ProtocolIEIDMBSAreaSessionID: 1,
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

func (x *ProtocolIEContainerDistributionSetupRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DistributionSetupRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DistributionSetupRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDistributionSetupResponseIEs struct {
	List []DistributionSetupResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDistributionSetupResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:     0,
		ProtocolIEIDMBSAreaSessionID: 1,
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

func (x *ProtocolIEContainerDistributionSetupResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DistributionSetupResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DistributionSetupResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDistributionSetupFailureIEs struct {
	List []DistributionSetupFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDistributionSetupFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:     0,
		ProtocolIEIDMBSAreaSessionID: 1,
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

func (x *ProtocolIEContainerDistributionSetupFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DistributionSetupFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DistributionSetupFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDistributionReleaseRequestIEs struct {
	List []DistributionReleaseRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDistributionReleaseRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:     0,
		ProtocolIEIDMBSAreaSessionID: 1,
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

func (x *ProtocolIEContainerDistributionReleaseRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DistributionReleaseRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DistributionReleaseRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerDistributionReleaseResponseIEs struct {
	List []DistributionReleaseResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerDistributionReleaseResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
		ProtocolIEIDMBSAreaSessionID:       1,
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

func (x *ProtocolIEContainerDistributionReleaseResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []DistributionReleaseResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val DistributionReleaseResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionActivationRequestIEs struct {
	List []MulticastSessionActivationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionActivationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerMulticastSessionActivationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionActivationRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionActivationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionActivationResponseIEs struct {
	List []MulticastSessionActivationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionActivationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
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

func (x *ProtocolIEContainerMulticastSessionActivationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionActivationResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionActivationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionActivationFailureIEs struct {
	List []MulticastSessionActivationFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionActivationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
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

func (x *ProtocolIEContainerMulticastSessionActivationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionActivationFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionActivationFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionDeactivationRequestIEs struct {
	List []MulticastSessionDeactivationRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionDeactivationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID: 0,
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

func (x *ProtocolIEContainerMulticastSessionDeactivationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionDeactivationRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionDeactivationRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionDeactivationResponseIEs struct {
	List []MulticastSessionDeactivationResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionDeactivationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
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

func (x *ProtocolIEContainerMulticastSessionDeactivationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionDeactivationResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionDeactivationResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionUpdateRequestIEs struct {
	List []MulticastSessionUpdateRequestIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionUpdateRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:     0,
		ProtocolIEIDMBSAreaSessionID: 1,
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

func (x *ProtocolIEContainerMulticastSessionUpdateRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionUpdateRequestIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionUpdateRequestIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionUpdateResponseIEs struct {
	List []MulticastSessionUpdateResponseIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionUpdateResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
		ProtocolIEIDMBSAreaSessionID:       1,
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

func (x *ProtocolIEContainerMulticastSessionUpdateResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionUpdateResponseIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionUpdateResponseIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionUpdateFailureIEs struct {
	List []MulticastSessionUpdateFailureIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionUpdateFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:           0,
		ProtocolIEIDMBSAreaSessionID:       1,
		ProtocolIEIDCause:                  2,
		ProtocolIEIDCriticalityDiagnostics: 3,
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

func (x *ProtocolIEContainerMulticastSessionUpdateFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionUpdateFailureIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionUpdateFailureIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastGroupPagingIEs struct {
	List []MulticastGroupPagingIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastGroupPagingIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:                 0,
		ProtocolIEIDMBSServiceArea:               1,
		ProtocolIEIDMulticastGroupPagingAreaList: 2,
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

func (x *ProtocolIEContainerMulticastGroupPagingIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastGroupPagingIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastGroupPagingIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMBSSessionSetupOrModRequestTransferIEs struct {
	List []MBSSessionSetupOrModRequestTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMBSSessionSetupOrModRequestTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionTNLInfo5GC:        0,
		ProtocolIEIDMBSQoSFlowsToBeSetupModList: 1,
		ProtocolIEIDMBSSessionFSAIDList:         2,
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

func (x *ProtocolIEContainerMBSSessionSetupOrModRequestTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MBSSessionSetupOrModRequestTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MBSSessionSetupOrModRequestTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerMulticastSessionUpdateRequestTransferIEs struct {
	List []MulticastSessionUpdateRequestTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerMulticastSessionUpdateRequestTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDMBSSessionID:                0,
		ProtocolIEIDMBSServiceArea:              1,
		ProtocolIEIDMBSQoSFlowsToBeSetupModList: 2,
		ProtocolIEIDMBSQoSFlowToReleaseList:     3,
		ProtocolIEIDMBSSessionTNLInfo5GC:        4,
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

func (x *ProtocolIEContainerMulticastSessionUpdateRequestTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []MulticastSessionUpdateRequestTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val MulticastSessionUpdateRequestTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs struct {
	List []PDUSessionResourceModifyRequestTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPDUSessionAggregateMaximumBitRate:        0,
		ProtocolIEIDULNGUUPTNLModifyList:                     1,
		ProtocolIEIDNetworkInstance:                          2,
		ProtocolIEIDQosFlowAddOrModifyRequestList:            3,
		ProtocolIEIDQosFlowToReleaseList:                     4,
		ProtocolIEIDAdditionalULNGUUPTNLInformation:          5,
		ProtocolIEIDCommonNetworkInstance:                    6,
		ProtocolIEIDAdditionalRedundantULNGUUPTNLInformation: 7,
		ProtocolIEIDRedundantCommonNetworkInstance:           8,
		ProtocolIEIDRedundantULNGUUPTNLInformation:           9,
		ProtocolIEIDSecurityIndication:                       10,
		ProtocolIEIDMBSSessionSetuporModifyRequestList:       11,
		ProtocolIEIDMBSSessionToReleaseList:                  12,
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

func (x *ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceModifyRequestTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceModifyRequestTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

type ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs struct {
	List []PDUSessionResourceSetupRequestTransferIEs // sizeLB:0,sizeUB:65535
}

func (x *ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Sort items to be in sequence
	order := map[int64]int{
		ProtocolIEIDPDUSessionAggregateMaximumBitRate:        0,
		ProtocolIEIDULNGUUPTNLInformation:                    1,
		ProtocolIEIDAdditionalULNGUUPTNLInformation:          2,
		ProtocolIEIDDataForwardingNotPossible:                3,
		ProtocolIEIDPDUSessionType:                           4,
		ProtocolIEIDSecurityIndication:                       5,
		ProtocolIEIDNetworkInstance:                          6,
		ProtocolIEIDQosFlowSetupRequestList:                  7,
		ProtocolIEIDCommonNetworkInstance:                    8,
		ProtocolIEIDDirectForwardingPathAvailability:         9,
		ProtocolIEIDRedundantULNGUUPTNLInformation:           10,
		ProtocolIEIDAdditionalRedundantULNGUUPTNLInformation: 11,
		ProtocolIEIDRedundantCommonNetworkInstance:           12,
		ProtocolIEIDRedundantPDUSessionInformation:           13,
		ProtocolIEIDMBSSessionSetupRequestList:               14,
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

func (x *ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 0, 65535
	var numElements uint64
	numElements, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PDUSessionResourceSetupRequestTransferIEs{}
	for i := 0; i < int(numElements); i++ {
		var val PDUSessionResourceSetupRequestTransferIEs
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
