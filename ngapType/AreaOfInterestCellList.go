package ngapType

// Need to import "github.com/free5gc/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct AreaOfInterestCellList */
/* AreaOfInterestCellItem */
type AreaOfInterestCellList struct {
	List []AreaOfInterestCellItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
