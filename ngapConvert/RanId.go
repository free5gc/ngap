package ngapConvert

import (
	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

func RanIdToModels(ranNodeId ngapType.GlobalRANNodeID) (ranId models.GlobalRanNodeId) {
	present := ranNodeId.Present
	switch present {
	case ngapType.GlobalRANNodeIDPresentGlobalGNBID:
		ranId.GNbId = new(models.GNbId)
		gnbId := ranId.GNbId
		ngapGnbId := ranNodeId.GlobalGNBID
		plmnid := PlmnIdToModels(ngapGnbId.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapGnbId.GNBID.Present == ngapType.GNBIDPresentGNBID {
			choiceGnbId := ngapGnbId.GNBID.GNBID
			gnbId.BitLength = int32(choiceGnbId.BitLength)
			gnbId.GNBValue = BitStringToHex(choiceGnbId)
		}
	case ngapType.GlobalRANNodeIDPresentGlobalNgENBID:
		ngapNgENBID := ranNodeId.GlobalNgENBID
		plmnid := PlmnIdToModels(ngapNgENBID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		switch ngapNgENBID.NgENBID.Present {
		case ngapType.NgENBIDPresentMacroNgENBID:
			macroNgENBID := ngapNgENBID.NgENBID.MacroNgENBID
			ranId.NgeNbId = "MacroNGeNB-" + BitStringToHex(macroNgENBID)
		case ngapType.NgENBIDPresentShortMacroNgENBID:
			shortMacroNgENBID := ngapNgENBID.NgENBID.ShortMacroNgENBID
			ranId.NgeNbId = "SMacroNGeNB-" + BitStringToHex(shortMacroNgENBID)
		case ngapType.NgENBIDPresentLongMacroNgENBID:
			longMacroNgENBID := ngapNgENBID.NgENBID.LongMacroNgENBID
			ranId.NgeNbId = "LMacroNGeNB-" + BitStringToHex(longMacroNgENBID)
		}
	case ngapType.GlobalRANNodeIDPresentGlobalN3IWFID:
		ngapN3IWFID := ranNodeId.GlobalN3IWFID
		plmnid := PlmnIdToModels(ngapN3IWFID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapN3IWFID.N3IWFID.Present == ngapType.N3IWFIDPresentN3IWFID {
			choiceN3IWFID := ngapN3IWFID.N3IWFID.N3IWFID
			ranId.N3IwfId = BitStringToHex(choiceN3IWFID)
		}
	case ngapType.GlobalRANNodeIDPresentChoiceExtensions:
		switch ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.Present {
		case ngapType.GlobalRANNodeIDExtIEsPresentGlobalTNGFID:
			ngapTNGFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalTNGFID
			plmnid := PlmnIdToModels(ngapTNGFID.PLMNIdentity)
			ranId.PlmnId = &plmnid
			if ngapTNGFID.TNGFID.Present == ngapType.TNGFIDPresentTNGFID {
				choiceTNGFID := ngapTNGFID.TNGFID.TNGFID
				ranId.TngfId = BitStringToHex(choiceTNGFID)
			}
		case ngapType.GlobalRANNodeIDExtIEsPresentGlobalTWIFID:
			ngapTWIFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalTWIFID
			plmnid := PlmnIdToModels(ngapTWIFID.PLMNIdentity)
			ranId.PlmnId = &plmnid
			if ngapTWIFID.TWIFID.Present == ngapType.TWIFIDPresentTWIFID {
				choiceTWIFID := ngapTWIFID.TWIFID.TWIFID
				ranId.TwifId = BitStringToHex(choiceTWIFID)
			}
		case ngapType.GlobalRANNodeIDExtIEsPresentGlobalWAGFID:
			ngapWAGFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalWAGFID
			plmnid := PlmnIdToModels(ngapWAGFID.PLMNIdentity)
			ranId.PlmnId = &plmnid
			if ngapWAGFID.WAGFID.Present == ngapType.WAGFIDPresentWAGFID {
				choiceWAGFID := ngapWAGFID.WAGFID.WAGFID
				ranId.WagfId = BitStringToHex(choiceWAGFID)
			}
		}
	}

	return ranId
}

func RanIDToNgap(modelsRanNodeId models.GlobalRanNodeId) ngapType.GlobalRANNodeID {
	var ngapRanNodeId ngapType.GlobalRANNodeID

	if modelsRanNodeId.GNbId.BitLength != 0 {
		ngapRanNodeId.Present = ngapType.GlobalRANNodeIDPresentGlobalGNBID
		ngapRanNodeId.GlobalGNBID = new(ngapType.GlobalGNBID)
		globalGNBID := ngapRanNodeId.GlobalGNBID

		globalGNBID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalGNBID.GNBID.Present = ngapType.GNBIDPresentGNBID
		globalGNBID.GNBID.GNBID = new(aper.BitString)
		*globalGNBID.GNBID.GNBID = HexToBitString(modelsRanNodeId.GNbId.GNBValue, int(modelsRanNodeId.GNbId.BitLength))
	} else if modelsRanNodeId.NgeNbId != "" {
		ngapRanNodeId.Present = ngapType.GlobalRANNodeIDPresentGlobalNgENBID
		ngapRanNodeId.GlobalNgENBID = new(ngapType.GlobalNgENBID)
		globalNgENBID := ngapRanNodeId.GlobalNgENBID

		globalNgENBID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		ngENBID := &globalNgENBID.NgENBID
		if modelsRanNodeId.NgeNbId[:11] == "MacroNGeNB-" {
			ngENBID.Present = ngapType.NgENBIDPresentMacroNgENBID
			ngENBID.MacroNgENBID = new(aper.BitString)
			*ngENBID.MacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[11:], 18)
		} else if modelsRanNodeId.NgeNbId[:12] == "SMacroNGeNB-" {
			ngENBID.Present = ngapType.NgENBIDPresentShortMacroNgENBID
			ngENBID.ShortMacroNgENBID = new(aper.BitString)
			*ngENBID.ShortMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 20)
		} else if modelsRanNodeId.NgeNbId[:12] == "LMacroNGeNB-" {
			ngENBID.Present = ngapType.NgENBIDPresentLongMacroNgENBID
			ngENBID.LongMacroNgENBID = new(aper.BitString)
			*ngENBID.LongMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 21)
		}
	} else if modelsRanNodeId.N3IwfId != "" {
		ngapRanNodeId.Present = ngapType.GlobalRANNodeIDPresentGlobalN3IWFID
		ngapRanNodeId.GlobalN3IWFID = new(ngapType.GlobalN3IWFID)
		globalN3IWFID := ngapRanNodeId.GlobalN3IWFID

		globalN3IWFID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalN3IWFID.N3IWFID.Present = ngapType.N3IWFIDPresentN3IWFID
		globalN3IWFID.N3IWFID.N3IWFID = new(aper.BitString)
		*globalN3IWFID.N3IWFID.N3IWFID = HexToBitString(modelsRanNodeId.N3IwfId, len(modelsRanNodeId.N3IwfId)*4)
	}

	return ngapRanNodeId
}
