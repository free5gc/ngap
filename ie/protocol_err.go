package ie

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// TS 38.413 10.1 defines 3 protocol error cases:
// 1. Transfer Syntax Error (TS 38.413 10.2)
// 2. Abstract Syntax Error (TS 38.413 10.3)
// 3. Logical Error (Ts 38.413 10.4)

// Protocol error Case 1. - TransferSyntaxErr (TS 38.413 10.2)
//
//	occurs when the receiver is not able to decode the received physical message
//	(always detected in the process of ASN.1 decoding)
type TransferSyntaxErr struct {
	errTrace error
}

func (e *TransferSyntaxErr) Error() string {
	return fmt.Sprintf("Transfer Syntax Error: %+v", e.errTrace.Error())
}

func (e *TransferSyntaxErr) ErrorTrace() string {
	return fmt.Sprintf("error trace: %+v", e.errTrace)
}

func (e *TransferSyntaxErr) GetCause() *Cause {
	return &Cause{
		Choice: &CauseProtocol{
			Value: CauseProtocolPresentTransferSyntaxError,
		},
	}
}

func BuildTransferSyntaxErr(errTrace error) *TransferSyntaxErr {
	return &TransferSyntaxErr{
		errTrace: errTrace,
	}
}

// Protocol error Case 2. - AbstractSyntaxErr (TS 38.413 10.3)
// TODO:
// i. If an Abstract Syntax Error occurs, the receiver shall read the remaining message
//
//	Current implementation only detects one Abstract Syntax Error and then terminates the decoding procedure
type AbstractSyntaxErr struct {
	// Info of message that Abstract Syntax Error happened
	procedureCode int64
	msgType       aper.Enumerated
	procedureCrit aper.Enumerated
	// Abstract Syntax Error case (may contain report IE info)
	errCase AbstractSyntaxErrCase
	// Error trace before Abstract Syntax Error is built
	errTrace error
}

// Error returns the error message for AbstractSyntaxErr
func (e *AbstractSyntaxErr) Error() string {
	return fmt.Sprintf("Abstract Syntax Error: %+v (description: %s)", e.errTrace.Error(), e.errCase.GetDesc())
}

// ErrorTrace returns the error trace for AbstractSyntaxErr
func (e *AbstractSyntaxErr) ErrorTrace() string {
	return fmt.Sprintf("error trace: %+v", e.errTrace)
}

// ErrorCase returns the errCase for AbstractSyntaxErr
func (e *AbstractSyntaxErr) ErrorCase() AbstractSyntaxErrCase {
	return e.errCase
}

// GetCause returns the Cause UE,
//
//	which is used to build Error Indication or Response Message by NFs,
//	according to AbstractSyntaxErrCase
func (e *AbstractSyntaxErr) GetCause() (*Cause, error) {
	switch errCase := e.errCase.(type) {
	case *AbstractSyntaxErrNotComprehendedIE:
		if errCase.ReportIe == nil {
			return nil, errors.Errorf("No reporting ie set for AbstractSyntaxErrNotComprehendedIE")
		}
		if errCase.ReportIe.ieid == AbstractSyntaxErrReportIeIdProcedureCode {
			// TS 38.413 10.3.4.1
			// Not Comprehended ProcedureCode IE shall not require Cause IE
			// (It uses the Error Indication procedure to reject a procedure or to report an ignored procedure, which shall include the
			//  Procedure Code IE, the Triggering Message IE, and the Procedure Criticality IE in the Criticality Diagnostics IE)
			return nil, errors.Errorf("Not Comprehended ProcedureCode IE shall not request for Cause IE")
		} else if errCase.ReportIe.ieid == AbstractSyntaxErrReportIeIdMessageType {
			// TS 38.41.3 10.3.4.1A
			// the Error Indication procedure shall be initiated with an appropriate cause value
			if errCase.ReportIe.ieCrit == CriticalityReject {
				return &Cause{
					Choice: &CauseProtocol{
						Value: CauseProtocolPresentAbstractSyntaxErrorReject,
					},
				}, nil
			} else if errCase.ReportIe.ieCrit == CriticalityNotify {
				return &Cause{
					Choice: &CauseProtocol{
						Value: CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify,
					},
				}, nil
			} else {
				return nil, errors.Errorf("Ignore Criticality doesn't report error and shall not request for Cause IE")
			}
		} else {
			// TS 38.413 10.3.4.2
			// IEs other than the Procedure Code and Type of Message
			// Report with Response Message (may requires Cause IE) or Error Indication
			if errCase.ReportIe.ieCrit == CriticalityReject {
				return &Cause{
					Choice: &CauseProtocol{
						Value: CauseProtocolPresentAbstractSyntaxErrorReject,
					},
				}, nil
			} else if errCase.ReportIe.ieCrit == CriticalityNotify {
				return &Cause{
					Choice: &CauseProtocol{
						Value: CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify,
					},
				}, nil
			} else {
				return nil, errors.Errorf("Ignore Criticality doesn't report error and shall not request for Cause IE")
			}
		}
	case *AbstractSyntaxErrMissingIE:
		if errCase.ReportIe == nil {
			return nil, errors.Errorf("No reporting ie set for AbstractSyntaxErrMissingIE")
		}
		// Cause
		if errCase.ReportIe.ieCrit == CriticalityReject {
			return &Cause{
				Choice: &CauseProtocol{
					Value: CauseProtocolPresentAbstractSyntaxErrorReject,
				},
			}, nil
		} else if errCase.ReportIe.ieCrit == CriticalityNotify {
			return &Cause{
				Choice: &CauseProtocol{
					Value: CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify,
				},
			}, nil
		} else {
			return nil, errors.Errorf("Ignore Criticality doesn't report error and shall not request for Cause IE")
		}
	case *AbstractSyntaxErrIEWrongOrderOrTooManyOccur:
	case *AbstractSyntaxErrIEErrPresent:
		// TS 38.413 10.3.6
		// The receiving node shall reject the procedure and report the cause value
		// "Abstract Syntax Error (Falsely Constructed Message)"
		return &Cause{
			Choice: &CauseProtocol{
				Value: CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage,
			},
		}, nil
	}
	return nil, errors.Errorf("Unknown AbstractSyntaxErr case")
}

// GetCritDiag returns the CriticalityDiagnostics IE,
//
//	which is used to build Error Indication or Response Message by NFs,
//	according to AbstractSyntaxErrCase
//
// Param: forErrInd - if is false, Procedure Code, Triggering Message, and Procedure Criticality shall not be included in Criticality Diagnostics IE
func (e *AbstractSyntaxErr) GetCritDiag(forErrInd bool) (*CriticalityDiagnostics, error) {
	switch errCase := e.errCase.(type) {
	case *AbstractSyntaxErrNotComprehendedIE:
		if errCase.ReportIe == nil {
			return nil, errors.Errorf("No reporting ie set for AbstractSyntaxErrNotComprehendedIE")
		}
		if errCase.ReportIe.ieid == AbstractSyntaxErrReportIeIdProcedureCode {
			// TS 38.413 10.3.4.1
			// When using the Error Indication procedure to reject a procedure or to report an ignored procedure
			// it shall include the Procedure Code IE, the Triggering Message IE, and the Procedure Criticality IE in
			// the Criticality Diagnostics
			// (i.e. Build Error Indication using CritDiag and don't include
			// Information Element Criticality Diagnostics IE in CritiDiag)

			// CriticalityDiagnostics
			return &CriticalityDiagnostics{
				ProcedureCode: &ProcedureCode{
					Value: e.procedureCode,
				},
				TriggeringMessage: &TriggeringMessage{
					Value: e.msgType,
				},
				ProcedureCriticality: &Criticality{
					Value: e.procedureCrit,
				},
			}, nil
		} else if errCase.ReportIe.ieid == AbstractSyntaxErrReportIeIdMessageType {
			// TS 38.413 10.3.4.1A
			// the Error Indication procedure shall be initiated with an appropriate cause value
			return nil, errors.Errorf("Not comprehended Message Type IE shall not request for Criticality Diagnostics")
		} else {
			// TS 38.413 10.3.4.2
			// IEs other than the Procedure Code and Type of Message
			// Report with Response Message or Error Indication
			// If Criticality Diagnostics is not used in Error Indication, Procedure Code, Triggering Message, and Procedure Criticality shall not be included
			critDiag := &CriticalityDiagnostics{
				IEsCriticalityDiagnostics: &CriticalityDiagnosticsIEList{
					List: []CriticalityDiagnosticsIEItem{
						{
							IECriticality: &Criticality{
								Value: errCase.ReportIe.ieCrit,
							},
							IEID: &ProtocolIEID{
								Value: errCase.ReportIe.ieid,
							},
							TypeOfError: &TypeOfError{
								Value: TypeOfErrorPresentNotUnderstood,
							},
						},
					},
				},
			}
			if forErrInd {
				critDiag.ProcedureCode = &ProcedureCode{
					Value: e.procedureCode,
				}
				critDiag.TriggeringMessage = &TriggeringMessage{
					Value: e.msgType,
				}
				critDiag.ProcedureCriticality = &Criticality{
					Value: e.procedureCrit,
				}
			}
			return critDiag, nil
		}
	case *AbstractSyntaxErrMissingIE:
		if errCase.ReportIe == nil {
			return nil, errors.Errorf("No reporting ie set for AbstractSyntaxErrMissingIE")
		}

		// Similar to the cases of AbstractSyntaxErrNotComprehendedIE (other than ProcedureCode and MessageType),
		// except for that the TypeOfError is TypeOfErrorPresentMissing
		critDiag := &CriticalityDiagnostics{
			IEsCriticalityDiagnostics: &CriticalityDiagnosticsIEList{
				List: []CriticalityDiagnosticsIEItem{
					{
						IECriticality: &Criticality{
							Value: errCase.ReportIe.ieCrit,
						},
						IEID: &ProtocolIEID{
							Value: errCase.ReportIe.ieid,
						},
						TypeOfError: &TypeOfError{
							Value: TypeOfErrorPresentMissing,
						},
					},
				},
			},
		}
		if forErrInd {
			critDiag.ProcedureCode = &ProcedureCode{
				Value: e.procedureCode,
			}
			critDiag.TriggeringMessage = &TriggeringMessage{
				Value: e.msgType,
			}
			critDiag.ProcedureCriticality = &Criticality{
				Value: e.procedureCrit,
			}
		}
		return critDiag, nil
	case *AbstractSyntaxErrIEWrongOrderOrTooManyOccur:
	case *AbstractSyntaxErrIEErrPresent:
		// TS 38.413 10.3.6
		// The receiving node shall reject the procedure and report the cause value
		// "Abstract Syntax Error (Falsely Constructed Message)"
		return nil, errors.Errorf("AbstractSyntaxErr of Falsely Constructed Message case shall not request for Criticality Diagnostics")
	}
	return nil, errors.Errorf("Unknown AbstractSyntaxErr case")
}

// IsInRspMsg returns whether the report ie is in a response message
// (i.e. an Abstract Syntax Error happens in a response message)
const (
	// messageTypeInitiatingMessage   aper.Enumerated = 0
	messageTypeSuccessfulOutcome   aper.Enumerated = 1
	messageTypeUnsuccessfulOutcome aper.Enumerated = 2
)

func (e *AbstractSyntaxErr) IsInRspMsg() bool {
	return (e.msgType == messageTypeSuccessfulOutcome) || (e.msgType == messageTypeUnsuccessfulOutcome)
}

func (e *AbstractSyntaxErr) GetProcedureCode() int64 {
	return e.procedureCode
}

func (e *AbstractSyntaxErr) GetMsgType() aper.Enumerated {
	return e.msgType
}

func (e *AbstractSyntaxErr) GetProcedureCrit() aper.Enumerated {
	return e.procedureCrit
}

func BuildAbstractSyntaxErr(procedureCode int64, msgType aper.Enumerated, procedureCrit aper.Enumerated,
	errCase AbstractSyntaxErrCase, errTrace error,
) *AbstractSyntaxErr {
	return &AbstractSyntaxErr{
		procedureCode: procedureCode,
		msgType:       msgType,
		procedureCrit: procedureCrit,
		errTrace:      errTrace,
		errCase:       errCase,
	}
}

// AbstractSyntaxErrCase defines the interface used for AbstractSyntaxErr.errCase
// Notes:
// i. Abstract Syntax Error Cases are defined in 38.413 10.3:
//
//	a. Not Comprehended IE/IE group
//		- unknown IE ID
//		- logical range violated
//	b. missing IEs or IE groups
//	c. IEs or IE groups are in wrong order or with too many occurences
//	d. erroneously present of conditional IEs or IE groups
type AbstractSyntaxErrCase interface {
	GetDesc() string // Get description of the error case
}

// AbstractSyntaxErrReportIe defines the required information to build Criticality Diagnostics IE
// Notes:
// i. Only Abstract Syntax Error cases a and b have to build Information Element Criticality Diagnostics IE
//
//	for CriticalityDiagonosis and use this AbstractSyntaxErrReport
//
// ii. When it comes to uncomprehended ProcedureCode, use AbstractSyntaxErrReportIeIdProcedureCode for
//
//	AbstractSyntaxErrReportid
//
// iii. When it comes to uncomprehended ProcedureCode, use AbstractSyntaxErrReportIeIdProcedureCode for
//
//	AbstractSyntaxErrReportid and AbstractSyntaxErrReportIeCritDontCare for AbstractSyntaxErrReportCriti
type AbstractSyntaxErrReportIe struct {
	// Info of the reportIe
	ieid   int64
	ieCrit aper.Enumerated
}

const (
	AbstractSyntaxErrReportIeIdProcedureCode = -1
	AbstractSyntaxErrReportIeIdMessageType   = -2
)

// IsProcedureCodeIe returns whether the report ie is ProcedureCode IE
func (e *AbstractSyntaxErrReportIe) IsProcedureCodeIe() bool {
	return e.ieid == AbstractSyntaxErrReportIeIdProcedureCode
}

// IsMessageTypeIe returns whether the report ie is MessageType IE
func (e *AbstractSyntaxErrReportIe) IsMsgTypeIe() bool {
	return e.ieid == AbstractSyntaxErrReportIeIdMessageType
}

// IsCritRejectOrNotif returns whether the criticality of the report ie is either Reject or Notify
func (e *AbstractSyntaxErrReportIe) IsCritRejectOrNotif() bool {
	return (e.ieCrit == CriticalityReject) || (e.ieCrit == CriticalityNotify)
}

// BuildAbstractSyntaxErrReportIe builds and returns AbstractSyntaxErrReportIe
func BuildAbstractSyntaxErrReportIe(ieId int64, ieCrit aper.Enumerated) *AbstractSyntaxErrReportIe {
	return &AbstractSyntaxErrReportIe{
		ieid:   ieId,
		ieCrit: ieCrit,
	}
}

// Abstract Syntax Error case a. - Not comprehended IEs/IE groups
// AbstractSyntaxErrNotComprehendedIE implements AbstractSyntaxErrCase interface
// TODO:
// i. check logical range of IEs
type AbstractSyntaxErrNotComprehendedIE struct {
	ReportIe *AbstractSyntaxErrReportIe
}

func (e *AbstractSyntaxErrNotComprehendedIE) GetDesc() string {
	return "Not Comprehended IE/IE group (unknown IE ID or logical range violated)"
}

// Abstract Syntax Error case b. - missing IEs/IE groups
// AbstractSyntaxErrMissingIE implements AbstractSyntaxErrCase interface
type AbstractSyntaxErrMissingIE struct {
	ReportIe *AbstractSyntaxErrReportIe
}

func (e *AbstractSyntaxErrMissingIE) GetDesc() string {
	return "Missing IEs or IE groups"
}

// Abstract Syntax Error case c. - IEs or IE groups are in wrong order or with too many occurences
// AbstractSyntaxErrIEWrongOrderOrTooManyOccur implements AbstractSyntaxErrCase interface
// TODO: check order and occurences of IEs
type AbstractSyntaxErrIEWrongOrderOrTooManyOccur struct{}

func (e *AbstractSyntaxErrIEWrongOrderOrTooManyOccur) GetDesc() string {
	return "IEs or IE groups are in wrong order or with too many occurences"
}

// Abstract Syntax Error case d. - erroneously present of conditional IEs or IE groups
// AbstractSyntaxErrIEErrPresent implements AbstractSyntaxErrCase interface
type AbstractSyntaxErrIEErrPresent struct{}

func (e *AbstractSyntaxErrIEErrPresent) GetDesc() string {
	return "conditional IEs or IE groups are erroneously present"
}

// Protocol error Case 3. - Logical Error (Ts 38.413 10.4)
//		information contained within the message is not valid (i.e., semantic error),
//		or describes a procedure which is not compatible with the state of the receiver
// TODO:
// i. Detect and handle logical error in NFs
