package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SpatialRelationforResourceID struct {
	List []SpatialRelationforResourceIDItem // valueExt,sizeLB:1,sizeUB:64
}

func (x *SpatialRelationforResourceID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Sequence Of
	*sLb, *sUb = 1, 64
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

func (x *SpatialRelationforResourceID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []SpatialRelationforResourceIDItem{}
	for i := 0; i < int(numElementsList); i++ {
		var val SpatialRelationforResourceIDItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
