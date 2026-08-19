package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type WLANMeasurementResultItem struct {
	WLANRSSI        *WLANRSSI
	SSID            *SSID                                                      // optional
	BSSID           *BSSID                                                     // optional
	HESSID          *HESSID                                                    // optional
	OperatingClass  *WLANOperatingClass                                        // optional
	CountryCode     *WLANCountryCode                                           // valueExt,valueLB:0,valueUB:3,optional
	WLANChannelList *WLANChannelList                                           // optional
	WLANBand        *WLANBand                                                  // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions    *ProtocolExtensionContainerWLANMeasurementResultItemExtIEs // optional
}

func (x *WLANMeasurementResultItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasurementResultItemOptPresentFlag := []bool{}
	// mandatory field
	if x.WLANRSSI == nil {
		return errors.Errorf("WLANRSSI is missing")
	}
	// optional field
	if x.SSID != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.BSSID != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.HESSID != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.OperatingClass != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.CountryCode != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.WLANChannelList != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.WLANBand != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, true)
	} else {
		WLANMeasurementResultItemOptPresentFlag = append(WLANMeasurementResultItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasurementResultItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.WLANRSSI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "WLANRSSI marshal failed")
	}

	// optional field
	if x.SSID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSID marshal failed")
		}
	}

	// optional field
	if x.BSSID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BSSID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BSSID marshal failed")
		}
	}

	// optional field
	if x.HESSID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.HESSID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "HESSID marshal failed")
		}
	}

	// optional field
	if x.OperatingClass != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OperatingClass.Write(pd)
		if err != nil {
			return errors.Wrap(err, "OperatingClass marshal failed")
		}
	}

	// optional field
	if x.CountryCode != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CountryCode.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CountryCode marshal failed")
		}
	}

	// optional field
	if x.WLANChannelList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANChannelList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "WLANChannelList marshal failed")
		}
	}

	// optional field
	if x.WLANBand != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANBand.Write(pd)
		if err != nil {
			return errors.Wrap(err, "WLANBand marshal failed")
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

func (x *WLANMeasurementResultItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasurementResultItemOptPresentFlag := make([]bool, 8)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasurementResultItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.WLANRSSI = new(WLANRSSI)
	err = x.WLANRSSI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode WLANRSSI error")
	}

	// optional field (optPresentFlag index: 0)
	if WLANMeasurementResultItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SSID = new(SSID)
		err = x.SSID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if WLANMeasurementResultItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.BSSID = new(BSSID)
		err = x.BSSID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BSSID error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if WLANMeasurementResultItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.HESSID = new(HESSID)
		err = x.HESSID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode HESSID error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if WLANMeasurementResultItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.OperatingClass = new(WLANOperatingClass)
		err = x.OperatingClass.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode OperatingClass error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if WLANMeasurementResultItemOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.CountryCode = new(WLANCountryCode)
		err = x.CountryCode.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CountryCode error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if WLANMeasurementResultItemOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.WLANChannelList = new(WLANChannelList)
		err = x.WLANChannelList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode WLANChannelList error")
		}
	}

	// optional field (optPresentFlag index: 6)
	if WLANMeasurementResultItemOptPresentFlag[6] {
		// Read struct defined elsewhere (Pointer)
		x.WLANBand = new(WLANBand)
		err = x.WLANBand.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode WLANBand error")
		}
	}

	// optional field (optPresentFlag index: 7)
	if WLANMeasurementResultItemOptPresentFlag[7] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerWLANMeasurementResultItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
