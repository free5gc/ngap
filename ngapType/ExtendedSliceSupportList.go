package ngapType

/* Sequence of = 35, FULL Name = struct ExtendedSliceSupportList */
/* SliceSupportItem */
type ExtendedSliceSupportList struct {
	List []SliceSupportItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
