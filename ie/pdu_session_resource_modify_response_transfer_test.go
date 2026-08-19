package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceModifyResponseTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceModifyResponseTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: &PDUSessionResourceModifyResponseTransfer{
				QosFlowAddOrModifyResponseList: &QosFlowAddOrModifyResponseList{
					List: []QosFlowAddOrModifyResponseItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 2,
							},
						},
					},
				},
			},
			expected: []byte{
				0x10, 0x00, 0x08,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := MarshalBinary(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestPDUSessionResourceModifyResponseTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceModifyResponseTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: []byte{
				0x10, 0x00, 0x08,
			},
			expected: &PDUSessionResourceModifyResponseTransfer{
				QosFlowAddOrModifyResponseList: &QosFlowAddOrModifyResponseList{
					List: []QosFlowAddOrModifyResponseItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 2,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PDUSessionResourceModifyResponseTransfer)
			err := UnmarshalBinary(tc.input, ie)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, ie)
			} else {
				require.Error(t, err)
			}
		})
	}
}
