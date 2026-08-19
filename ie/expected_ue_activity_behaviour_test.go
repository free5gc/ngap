package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpectedUEActivityBehaviourMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *ExpectedUEActivityBehaviour
		expected []byte
	}{
		{
			name: "Case 1",
			input: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 30,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 1,
				},
			},
			expected: []byte{0x60, 0x74, 0x00},
		},
		{
			name: "Case 2",
			input: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 41,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 180,
				},
			},
			expected: []byte{0x64, 0x01, 0x29, 0x59, 0x80},
		},
		{
			name: "Case 3",
			input: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 0,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 1000,
				},
			},
			expected: []byte{0x64, 0x01, 0x00, 0x80, 0x02, 0x03, 0xE8},
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

func TestExpectedUEActivityBehaviourUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *ExpectedUEActivityBehaviour
	}{
		{
			name:  "Case 1",
			input: []byte{0x60, 0x74, 0x00},
			expected: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 30,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 1,
				},
			},
		},
		{
			name:  "Case 2",
			input: []byte{0x64, 0x01, 0x29, 0x59, 0x80},
			expected: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 41,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 180,
				},
			},
		},
		{
			name:  "Case 3",
			input: []byte{0x64, 0x01, 0x00, 0x80, 0x02, 0x03, 0xE8},
			expected: &ExpectedUEActivityBehaviour{
				ExpectedActivityPeriod: &ExpectedActivityPeriod{
					Value: 0,
				},
				ExpectedIdlePeriod: &ExpectedIdlePeriod{
					Value: 1000,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(ExpectedUEActivityBehaviour)
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
